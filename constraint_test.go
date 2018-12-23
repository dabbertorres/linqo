package linqo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestConstraintReferences(t *testing.T) {
	type args struct {
		table   string
		columns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				table: "foo",
				columns: []string{ "bar", "baz", "qux" },
			},
			want: "REFERENCES foo (bar,baz,qux)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstraintReferences(tt.args.table, tt.args.columns...); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("ConstraintReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constraintReferences_Match(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		m ConstraintMatch
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "full",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ConstraintMatchFull},
			want: " MATCH FULL",
		},
		{
			name: "partial",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ConstraintMatchPartial},
			want: " MATCH PARTIAL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := constraintReferences{
				Buffer: tt.fields.Buffer,
			}
			if got := c.Match(tt.args.m); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("constraintReferences.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constraintReferences_OnUpdate(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		a ReferentialAction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "cascade",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionCascade},
			want: " ON UPDATE CASCADE",
		},
		{
			name: "set null",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionSetNull},
			want: " ON UPDATE SET NULL",
		},
		{
			name: "set default",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionSetDefault},
			want: " ON UPDATE SET DEFAULT",
		},
		{
			name: "no action",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{NoAction},
			want: " ON UPDATE NO ACTION",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := constraintReferences{
				Buffer: tt.fields.Buffer,
			}
			if got := c.OnUpdate(tt.args.a); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("constraintReferences.OnUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constraintReferences_OnDelete(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		a ReferentialAction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "cascade",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionCascade},
			want: " ON DELETE CASCADE",
		},
		{
			name: "set null",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionSetNull},
			want: " ON DELETE SET NULL",
		},
		{
			name: "set default",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{ActionSetDefault},
			want: " ON DELETE SET DEFAULT",
		},
		{
			name: "no action",
			fields: fields{bytes.NewBuffer(nil)},
			args: args{NoAction},
			want: " ON DELETE NO ACTION",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := constraintReferences{
				Buffer: tt.fields.Buffer,
			}
			if got := c.OnDelete(tt.args.a); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("constraintReferences.OnDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
