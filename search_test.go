package linqo

import "testing"

func TestOr(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo OR bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo AND bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNot(t *testing.T) {
	type args struct {
		term SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				term: "foo",
			},
			want: "(NOT foo)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Not(tt.args.term); string(got) != tt.want {
				t.Errorf("Not() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnknown(t *testing.T) {
	type args struct {
		term SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				term: "foo",
			},
			want: "(foo IS UNKNOWN)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnknown(tt.args.term); string(got) != tt.want {
				t.Errorf("IsUnknown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo = bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo <> bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEquals(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("NotEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo < bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessThan(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo > bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreaterThan(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLessOrEqual(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo <= bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessOrEqual(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("LessOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
	type args struct {
		left  SearchTerm
		right SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				left: "foo",
				right: "bar",
			},
			want: "(foo >= bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreaterOrEqual(tt.args.left, tt.args.right); string(got) != tt.want {
				t.Errorf("GreaterOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		value      SearchTerm
		lowerBound SearchTerm
		upperBound SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				lowerBound: "0",
				upperBound: "1",
			},
			want: "(foo BETWEEN 0 AND 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.args.value, tt.args.lowerBound, tt.args.upperBound); string(got) != tt.want {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotBetween(t *testing.T) {
	type args struct {
		value      SearchTerm
		lowerBound SearchTerm
		upperBound SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				lowerBound: "0",
				upperBound: "1",
			},
			want: "(foo NOT BETWEEN 0 AND 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotBetween(tt.args.value, tt.args.lowerBound, tt.args.upperBound); string(got) != tt.want {
				t.Errorf("NotBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		value SearchTerm
		list  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				list: []string{"foo", "bar", "baz"},
			},
			want: "(foo IN (foo,bar,baz))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.value, tt.args.list...); string(got) != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIn(t *testing.T) {
	type args struct {
		value SearchTerm
		list  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				list: []string{"foo", "bar", "baz"},
			},
			want: "(foo NOT IN (foo,bar,baz))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotIn(tt.args.value, tt.args.list...); string(got) != tt.want {
				t.Errorf("NotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLike(t *testing.T) {
	type args struct {
		value   string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				pattern: "[fb][oa][or]",
			},
			want: "(foo LIKE [fb][oa][or])",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Like(tt.args.value, tt.args.pattern); string(got) != tt.want {
				t.Errorf("Like() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotLike(t *testing.T) {
	type args struct {
		value   string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				pattern: "[fb][oa][or]",
			},
			want: "(foo NOT LIKE [fb][oa][or])",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotLike(tt.args.value, tt.args.pattern); string(got) != tt.want {
				t.Errorf("NotLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLikeEscaped(t *testing.T) {
	type args struct {
		value   string
		pattern string
		escape  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				pattern: "[fb][oa][or]",
				escape: "f",
			},
			want: "(foo LIKE [fb][oa][or] ESCAPE f)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LikeEscaped(tt.args.value, tt.args.pattern, tt.args.escape); string(got) != tt.want {
				t.Errorf("LikeEscaped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotLikeEscaped(t *testing.T) {
	type args struct {
		value   string
		pattern string
		escape  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
				pattern: "[fb][oa][or]",
				escape: "f",
			},
			want: "(foo NOT LIKE [fb][oa][or] ESCAPE f)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotLikeEscaped(tt.args.value, tt.args.pattern, tt.args.escape); string(got) != tt.want {
				t.Errorf("NotLikeEscaped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNull(t *testing.T) {
	type args struct {
		value SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
			},
			want: "(foo IS NULL)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNull(tt.args.value); string(got) != tt.want {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotNull(t *testing.T) {
	type args struct {
		value SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value: "foo",
			},
			want: "(foo IS NOT NULL)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotNull(tt.args.value); string(got) != tt.want {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		query SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				query: "foo",
			},
			want: "(EXISTS foo)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.query); string(got) != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		query SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				query: "foo",
			},
			want: "(UNIQUE foo)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.query); string(got) != tt.want {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverlaps(t *testing.T) {
	type args struct {
		value1 SearchTerm
		value2 SearchTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				value1: "foo",
				value2: "bar",
			},
			want: "(foo OVERLAPS bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Overlaps(tt.args.value1, tt.args.value2); string(got) != tt.want {
				t.Errorf("Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
