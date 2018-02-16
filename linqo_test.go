package linqo

import (
	"database/sql"
	"fmt"
	"math"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type customer struct {
	lastName      string
	firstName     string
	totalSpending float64
}

func (c customer) String() string {
	return fmt.Sprintf("%10s, %10s: $%7.2f", c.lastName, c.firstName, c.totalSpending)
}

func (c customer) StringNoSpending() string {
	return fmt.Sprintf("%10s, %10s", c.lastName, c.firstName)
}

// welfords implements Welford's running (online) variance algorithm (and calculates means as a bonus)
func welfords(vals []float64) (dev, avg float64) {
	// note: in the loop, 'dev' and 'newDev' are actually variance
	for i, v := range vals {
		newAvg := avg + (v-avg)/float64(i+1)
		newDev := dev + (v-avg)*(v-newAvg)
		avg, dev = newAvg, newDev
	}
	// for our purposes, we're assuming vals is a whole population.
	// in most cases, you're dealing with a sample, so you'd divide by len(vals) - 1 instead
	dev = math.Sqrt(dev / float64(len(vals)))
	return
}

func Example() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	create :=
		Create("customers").
			Columns(
				Column("firstName", "varchar(64)"),
				Column("lastName", "varchar(64)"),
				Column("totalSpending", "decimal"),
			)

	insert :=
		Insert("customers").
			Columns("firstName", "lastName", "totalSpending").
			Values(
				ValuesList{"Annyong", "Annyong", 5000.00},
				ValuesList{"Bob", "Loblaw", 2000.00},
				ValuesList{"Steve", "Holt", 10.00},
				ValuesList{"Barry", "Zuckerkorn", 9999.99},
				ValuesList{"Doctor", "Fishman", 2499.99},
				ValuesList{"J. Walter", "Weatherman", 2999.95},
				ValuesList{"Rita", "Leeds", 0.00},
			)

	_, err = db.Exec(create.String())
	if err != nil {
		fmt.Println("Error executing create:", err)
		fmt.Println(create)
		return
	}

	_, err = db.Exec(insert.String())
	if err != nil {
		fmt.Println("Error executing insert:", err)
		fmt.Println(insert)
		return
	}

	selectAll :=
		Select().
			From("customers").
			OrderBy(SortSpec{
				Key:   "lastName",
				Order: Ascending,
			})

	rows, err := db.Query(selectAll.String())
	if err != nil {
		fmt.Println("Error querying selectAll:", err)
		fmt.Println(selectAll)
		return
	}

	fmt.Println("Customers:")
	totalCustomers := 0
	for rows.Next() {
		totalCustomers++
		var c customer
		err = rows.Scan(&c.firstName, &c.lastName, &c.totalSpending)
		if err != nil {
			fmt.Println("Error scanning customer:", err)
			continue
		}
		fmt.Println(c)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error iterating selectAll results:", err)
		return
	}

	selectSpending :=
		Select("totalSpending").
			From("customers")

	rows, err = db.Query(selectSpending.String())
	if err != nil {
		fmt.Println("error querying selectSpending:", err)
		fmt.Println(selectSpending)
		return
	}

	spendings := make([]float64, totalCustomers)
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&spendings[i])
		if err != nil {
			fmt.Println("Error scanning totalSpending:", err)
			continue
		}
	}

	spendingStdDev, averageSpending := welfords(spendings)
	lowerBound := averageSpending - spendingStdDev
	upperBound := averageSpending + spendingStdDev

	fmt.Println()
	fmt.Printf("  Average Spending: $%.2f\n", averageSpending)
	fmt.Printf("Standard Deviation: $%.2f\n", spendingStdDev)
	fmt.Println()
	fmt.Printf("Removing extremes outside of ($%.2f - $%.2f)\n", lowerBound, upperBound)
	fmt.Println()

	lbStr := SearchTerm(strconv.FormatFloat(lowerBound, 'f', 2, 64))
	ubStr := SearchTerm(strconv.FormatFloat(upperBound, 'f', 2, 64))

	filterHighLow :=
		Delete("customers").
			Where(
				Or(
					LessThan("totalSpending", lbStr),
					GreaterThan("totalSpending", ubStr),
				))

	res, err := db.Exec(filterHighLow.String())
	if err != nil {
		fmt.Println("Error executing filterHighLow:", err)
		fmt.Println(filterHighLow)
		return
	}

	removedRows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error getting removed rows count:", err)
		return
	}

	fmt.Println("Filtered out", removedRows, "rows")

	selectLeftOver :=
		Select("lastName", "firstName").
			From("customers").
			OrderBy(SortSpec{
				Key:   "lastName",
				Order: Ascending,
			})

	rows, err = db.Query(selectLeftOver.String())
	if err != nil {
		fmt.Println("Error querying selectLeftOver:", err)
		fmt.Println(selectLeftOver)
		return
	}

	fmt.Println("Left over customers:")
	for rows.Next() {
		var c customer
		err = rows.Scan(&c.lastName, &c.firstName)
		if err != nil {
			fmt.Println("Error scanning customer:", err)
			continue
		}
		fmt.Println(c.StringNoSpending())
	}
	fmt.Println()

	err = rows.Err()
	if err != nil {
		fmt.Println("Error iterating selectLeftOver results:", err)
		return
	}

	rows, err = db.Query(selectSpending.String())
	if err != nil {
		fmt.Println("Error querying selectSpending:", err)
		fmt.Println(selectSpending)
		return
	}

	totalCustomers -= int(removedRows)
	spendings = spendings[:totalCustomers]

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&spendings[i])
		if err != nil {
			fmt.Println("Error scanning totalSpending:", err)
			continue
		}
	}

	spendingStdDev, averageSpending = welfords(spendings)

	fmt.Println("Filtered stats:")
	fmt.Printf("  Average Spending: $%.2f\n", averageSpending)
	fmt.Printf("Standard Deviation: $%.2f\n", spendingStdDev)

	// Output:
	// Customers:
	//    Annyong,    Annyong: $5000.00
	//    Fishman,     Doctor: $2499.99
	//       Holt,      Steve: $  10.00
	//      Leeds,       Rita: $   0.00
	//     Loblaw,        Bob: $2000.00
	// Weatherman,  J. Walter: $2999.95
	// Zuckerkorn,      Barry: $9999.99
	//
	//   Average Spending: $3215.70
	// Standard Deviation: $3204.11
	//
	// Removing extremes outside of ($11.59 - $6419.81)
	//
	// Filtered out 3 rows
	// Left over customers:
	//    Annyong,    Annyong
	//    Fishman,     Doctor
	//     Loblaw,        Bob
	// Weatherman,  J. Walter
	//
	// Filtered stats:
	//   Average Spending: $3124.99
	// Standard Deviation: $1138.81
}
