package main

import (
	"reflect"
	"testing"
)

func TestReturnFibonacci_Single(t *testing.T) {
	in := 3
	want := []int{0, 1, 1}

	got := ReturnFibonacci(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReturnFibonacci(%d) != %v, got = %v", in, want, got)
	}
}

func TestReturnFibonacci_Table(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3 parameters",
			args: args{
				n: 3,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "4 parameters",
			args: args{
				n: 4,
			},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "5 parameters",
			args: args{
				n: 5,
			},
			want: []int{0, 1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReturnFibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
