// @Author huzejun 2024/7/7 20:48:00
/*package main

import "testing"

// 单元测试（*testing.T）
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

// 这是一个测试辅助函数，它设置了测试需要的一些初始条件
func setupTest(t *testing.T) (int, int) {
	t.Helper() //标记这个函数为测试辅助函数
	//这里可以是任何设置测试环境的代码
	t.Errorf("setupTest error")
	return 10, 20 //例如，返回两个用于测试的整数
}

// 这是一个测试函数
func TestAdd(t *testing.T) {
	a, b := setupTest(t) //调用测试辅助函数来设置测试数据
	got := add(a, b)
	want := 30
	if got != want {
		t.Errorf("Add(%d,%d) = %d; want %d", a, b, got, want)
	}
}
*/

package main

import "testing"

//基准测试（*testing.B）
//func BenchmarkSprintf(b *testing.B) {
//	num := 10
//	//重置计时器，确保从一个一致的状态开始
//	//基准测试函数 必须
//	//在真正执行时重置，可以忽略准备期的开销
//	b.ResetTimer()
//	//b.N就是执行次数
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("%d", num)
//	}
//}
//
//func BenchmarkFormat(b *testing.B) {
//	num := int64(10)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		strconv.FormatInt(num, 10)
//	}
//}

//func BenchmarkNewSlice(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		newSlice(1000000)
//	}
//}
//
//func BenchmarkNewSliceWithCap(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		newSliceWithCap(1000000)
//	}
//}

// 模糊测试 (*testing.F)
func FuzzDiv(f *testing.F) {
	//f.Add() //添加测试数据
	f.Fuzz(func(t *testing.T, a, b int) {
		Div(a, b)
	})
}
