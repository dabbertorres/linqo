package linqo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestColumn(t *testing.T) {
	type args struct {
		name     string
		dataType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				name: "foo",
				dataType: "VARCHAR(20)",
			},
			want: "foo VARCHAR(20)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Column(tt.args.name, tt.args.dataType); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_column_Default(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		v Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "base",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{"foo"},
			want: " DEFAULT 'foo'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := column{
				Buffer: tt.fields.Buffer,
			}
			if got := c.Default(tt.args.v); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("column.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_column_Constraints(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		cts []Constraint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "zero constraints",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{},
			want: "",
		},
		{
			name: "one constraint",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{[]Constraint{ConstraintNotNull()}},
			want: " NOT NULL",
		},
		{
			name: "two constraint",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{[]Constraint{ConstraintNotNull(), ConstraintPrimaryKey()}},
			want: " NOT NULL PRIMARY KEY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := column{
				Buffer: tt.fields.Buffer,
			}
			if got := c.Constraints(tt.args.cts...); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("column.Constraints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_column_Collate(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "base",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{"foo"},
			want: " COLLATE foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := column{
				Buffer: tt.fields.Buffer,
			}
			if got := c.Collate(tt.args.name); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("column.Collate() = %v, want %v", got, tt.want)
			}
		})
	}
}
