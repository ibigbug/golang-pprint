package pprint

import (
	"testing"
)

type T1 struct {
	Name string
	Age  T2
}

type T2 struct {
	age string
}

func (t T2) String() string {
	return t.age
}

func TestPprint(t *testing.T) {
	type args struct {
		o interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				o: T1{Name: "yyy", Age: T2{age: "uuu"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pprint(tt.args.o)
		})
	}
}
