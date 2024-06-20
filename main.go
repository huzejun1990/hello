// package 定义包名 main 包名
package main

import (
	"fmt"
	"hello/model"
)

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
	fmt.Println(model.Pi)
}

/*func getValue() {
	var i = 10
	fmt.Println(i)
}*/
