// package 定义包名 main 包名
package main

import "fmt"

// func 定义函数	main 函数名

// hello := "hello"
//var hello = 10

/*
var health = 100
var level = 1
*/
/*var (
	health = 100
	level  = 1
)

// PI PI的值
const PI = 3.14
*/
/*type Month int

const (
	//Month = 1 + iota
	January   Month = 1 + iota
	February        = 2
	March           = 3
	April           = 4
	May             = 5
	June            = 6
	August          = 8
	September       = 9
	October         = 10
	November        = 11
	December        = 12
)*/

//var i = 10

func main() {
	/*	// fmt 包名 . 调用 Print 函数，并且输出定义的字符串
		fmt.Print("Hello Golang2024")*/

	//游戏 玩家 血量  等级
	/*	var health int
		var level int
		fmt.Printf("玩家血量：%d, 等级:%d", health, level)
	*/

	/*	var ch byte = 'A'
		var ch1 byte = 65
		//fmt.Printf("ch=%c,ch1=%c", ch, ch1)
		//16进制
		var ch16 byte = '\x41'
		// 8进制
		var ch8 byte = '\101'
		fmt.Printf("ch=%c,ch1=%c,ch16=%c,ch8=%c", ch, ch1, ch16, ch8)
	*/

	/*	var ni rune = 20320
		var hao rune = '\u597D'
		var hao1 = rune('好')
		fmt.Printf("%c,%c,%U", ni, hao, hao1)*/

	/*	var str string = `努力自学\r\nGo教程`
		str = "mszlu"
		fmt.Printf("%s", str)*/

	//8080 8081
	/*	s1 := "localhost:8080"
		strByte := []byte(s1)
		strByte[len(strByte)-1] = '1'
		s2 := string(strByte)
		fmt.Println(s2)
	*/

	//fmt.Println(health)
	// 2 * PI
	//fmt.Println(2 * PI)

	//fmt.Println(December)
	//getValue()

	//{
	//	var i = 10
	//	fmt.Println(i)
	//}
	//fmt.Println(i)
	//fmt.Println(model.Pi)

	/*	var a int = 10
		var p *int
		p = &a
		fmt.Println(*p)
		*p = 20
		fmt.Println(a)*/

	//var p *int
	//var a int = 10
	//p = &a
	//new函数 初始化一个内存空间 申请了一个内存空间，自然就有地址
	//放的什么值? 对应类型的零值
	/*	p = new(int)
	*p = 20
	fmt.Println(p)*/
	//fmt.Println("p是指针，它是一个地址：", p)
	//fmt.Println("取p指向内存空间的值：", *p)

	//声明一个函数
	/*	var swap func(x, y int) (int, int)
		swap = func(x, y int) (int, int) {
			return y, x
		}
		fmt.Println(swap(10, 20))*/
	//swap(1, 2.3, 4.5, 6.7)
	/*	var a = 10
		var b = 20
		changeValue(a, b)
		fmt.Println(a, b)

		changeValue1(&a, &b)
		fmt.Println(a, b)*/

	/*	//这是一个匿名函数
		funA := func() int {
			return 20
		}
		//其实在这里funa就是函数的名字
		funA()
		//这是一个匿名函数调用 可以不用将函数声明为一个变量在使用
		func() {
			fmt.Println("这是一个匿名函数")
		}()
	*/

	//闭包

	//创建一个玩家生成器
	generator := playerGen("战神")
	//返回玩家的名字和血量
	name, hp := generator()
	//打印值
	fmt.Println(name, hp)

}

func playerGen(name string) func() (string, int) {
	//血量一直为150
	hp := 150
	//返回创建的闭包
	return func() (string, int) {
		//将变量引用到闭包中
		return name, hp
	}
}

/*func yes() (int, int, string) {
	var a int
	var b int
	var c string
	return a, b, c
}*/

/*func yes() (a int, b int, c string) {
	a = 10
	b = 20
	c = "hello"
	return
}*/

// 虽然是值传递，但是传递的是地址 修改的时候 还是原来a的内存空间，所以工受影响了
// 引用传递 指针占用的空间 4字节 或者 8字节
/*func changeValue1(x *int, y *int) {
	*x = 30
	*y = 50
}

func changeValue(x int, y int) {
	x = 30
	y = 50
}

// ...不固定参数 数组
func swap(x int, arr ...float64) {
	fmt.Println(x, arr)
}
*/
/*func swap(x, y int) {

}*/

/*func getValue() {
	var i = 10
	fmt.Println(i)
}*/