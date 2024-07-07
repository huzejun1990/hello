// @Author huzejun 2024/7/7 20:48:00
package main

import "testing"

func Test_reduce(t *testing.T) {

}
func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test with positive numbers",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "Test with negative numbers",
			args: args{a: -2, b: -3},
			want: -5,
		},
		{
			name: "Test with zero",
			args: args{a: 0, b: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
