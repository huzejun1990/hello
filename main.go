// package 定义包名 main 包名
package main

import (
	"fmt"
	"sync"
	"time"
)

/*func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
*/
/*type Any interface {
}*/

/*type Animal interface {
	//Say 动物可以说话
	Say()
	//Move() 动物可以移动
	Move()
	//Jump 动物可以跳起来
	Jump()
}

type Fighter interface {
	Hurt() int
}

type Dog struct {
}

func (d *Dog) Move() {
	fmt.Println("狗在跑")
}

func (d *Dog) Jump() {
	fmt.Println("狗在跳")
}

func (d *Dog) Say() {
	fmt.Println("汪汪汪")
}

func (d *Dog) Hurt() int {
	return 10
}

type Cat struct {
}

func (c *Cat) Say() {
	fmt.Println("喵喵喵")
}*/

/*
// User 定义结构体 字段可能很多
type User struct {
	Age  int    //8字节
	Name string //16字节
	//Address
}

func (u *User) modifyName(name string) {
	u.Name = name
}

// 64位系统长度为8
type test1 struct {
	a bool   //1
	b int32  //4
	c string //16
}

type test2 struct {
	a int32  // 4
	b string // 16
	c bool   // 1
}*/

/*
type Address struct {
	City   string
	Street string
	Seq    int
}

type myInt int

func (m myInt) getValue() int {
	return int(m)
}*/

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

//func main() {
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

/*	//创建一个玩家生成器
	generator := playerGen("战神")
	//返回玩家的名字和血量
	name, hp := generator()
	//打印值
	fmt.Println(name, hp)
*/

// 0,0,0
//var arr [3]int = [3]int{1, 2, 3}
//var arr = [...]int{1, 2, 3, 4, 5}
//var arr = [5]int{2: 3}
/*	var arr [3]int
	arr[2] = 20
	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Println(arr, arr[2])*/

// 0,0,0
// 1,2,3
/*	var arr [3][3][3][3]int
	fmt.Println(arr)*/

/*	var array [4]int
	slice := array[1:3]
	slice1 := slice[:]
	fmt.Println(cap(slice))
	fmt.Println(len(slice1))
	change(slice)
	fmt.Println(array)*/

/*	var slice []int = []int{0, 0, 0}
	change(slice)
	fmt.Println(slice)*/
// 长度 容量
/*	var s = make([]int, 0, 0)
	s1 := *new([]int)
	//s2 := &s
	fmt.Println(s, s1)*/
/*	s1 := s[0:8]
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d", len(s), cap(s))*/
/*	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6]
	//打印出来长度为2
	fmt.Printf("myslice为 %d, 其长度为：%d\n", myslice, len(myslice))
	fmt.Println(cap(myslice))
	myslice = myslice[:cap(myslice)]
	//为什么 myslice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为：%d", myslice[3])*/

/*	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) //只会复制slice1的前3个元素到slice2中
	//copy(slice1, slice2) //只会复制slice2的3个元素到slice1的前3个位置
	slice2[0] = 10
	fmt.Println(slice1, slice2)*/

//slice1 := []int{1, 2, 3, 4, 5}
/*	slice1 := make([]int, 257, 257)
	slice2 := append(slice1, 10)
	fmt.Printf("cap:%d \n", cap(slice2))
	fmt.Println(slice1, slice2)*/

//go1.17 版本新特性
/*	slice1 := make([]int, 2, 8)
	arr := *(*[2]int)(slice1)
	fmt.Println(arr)*/

//go1.20新特性
/*	slice1 := make([]int, 2, 8)
	arr := [2]int(slice1)
	fmt.Println(arr)*/

/*	var v int = 700
	// v < 1000 这是条件表达式
	if v > 1000 {
		//为true 执行这条语句 false则不执行
		fmt.Printf("v is less than 1000\n")
	} else {
		fmt.Printf("v is greater than 1000\n")
	}
	fmt.Printf("Value of v is : %d \n", v)*/
/*	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
		//报错了，直接结束
		panic("---")
		//return
		//break
	}
	//不会执行
	fmt.Println("结束之后的语句.....")*/
/*	var a = 10
		if a < 7 {
			goto gotoHere
		}
		return
		//goto gotoHere
	gotoHere:
		fmt.Println("tag")
		fmt.Println("tag1")
		fmt.Println("tag2")*/

/*	for x := 0; x < 10; x++ {
			if x == 2 {
				//跳转到标签
				goto breakHere
			}
		}
		//手动返回避免执行进入标签
		return
		//标签
	breakHere:
		fmt.Println("done")*/

/*	var slice []int = []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		go func() {
			//1.22版本这里会输出 1,2,3,4,5 顺序不一定
			//1.22版本以前 大概率会输出5
			//这里原因是1.22版本每次循环 index,value会创建新的变量，而不是共享变量了
			fmt.Printf("index:%d,value:%d \n", index, value)
		}()
	}
	for {

	}*/

/*	for i := range 11 {
	fmt.Println(i)
}*/

/* 定义局部变量	*/
/*	var grade string = "B"
	var score int = 90

	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	//switch后面如果没有条件表达式，则会对true进行匹配
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Println("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)*/

/*	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "world":
		fmt.Println("world")

	}*/

/*OuterLoop:
for i := 0; i < 2; i++ {
	for j := 0; j < 5; j++ {
		switch j {
		case 2:
			fmt.Println(i, j)
			continue OuterLoop

		}
	}
}*/

//10-map和结构体
// key int类型 value string类型
//nil 意味着没有内存地址 没有内存空间
//取值，不涉及内存的操作
/*	var m map[int]string = map[int]string{}
	// 2 代表容量，代表提前在内存中分配多大的空间
	// map 动态的结构 内存空间会随着数据变化
	// 扩容操作 双信扩容
	//var m = make(map[int]string)
	change(m)
	m[1] = "go教程"
	fmt.Println(m)*/

/*	m := make(map[string]string)
	m["key"] = "value"
	m["key1"] = "value1"
	m["key2"] = "value1"
	m["key3"] = "value1"
	//通过key取值
	//go中的变量，必须使用 不使用编译 _忽略当前的值
	v := m["key"]
	//delete(m, "key1")
	m = make(map[string]string)
	fmt.Println(v)
	//通过for range取值，注意无序的
	//想要有序的结果，可能遍历后对key排序
	for k, v := range m {
		fmt.Println(k, v)
	}*/

//结构体
/*	var user User
	//var users []*User
	//进行初始化的时候，推荐使用new
	userPoint := new(User)
	change(userPoint)
	fmt.Printf("age=%d\n", user.Age)*/

//方法和函数
/*	var mi myInt = 100
	fmt.Println(mi.getValue())*/

/*	u := &User{
		Name: "initName",
	}
	u.modifyName("dream")
	fmt.Println(u)*/

//内存对齐
/*	var t1 test1
	var t2 test2
	// t1 size if 24
	// t2 size if 32
	fmt.Println("t1 size is ", unsafe.Sizeof(t1))
	fmt.Println("t2 size is ", unsafe.Sizeof(t2))*/

//11-接口
/*	var animal Animal = &Dog{}
	//var animal Animal = &Cat{}
	animal.Say()
	//fmt.Println(animal.Say)
	var fi Fighter = &Dog{}
	fmt.Println(fi.Hurt())*/

//空接口
/*	var a any
	a = 10
		fmt.Println(a)
		a = "努力之路学golang"
	s, ok := a.(string)
	if ok {
		fmt.Printf("我是一个string :%s", s)
	} else {
		fmt.Println("不是一个string")
	}
	fmt.Println(a)*/

/*	var a any
	a = 10
	justifyType(a)*/

//}

// 类型断言
/*func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string,value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}*/

/*func change(user *User) {
	user.Age = 20
	user.City = "北京"
}*/

/*func change(m map[int]string) {
	m[1] = "dream"
}*/

/*func change(s []int) {
	s[0] = 10
}*/

/*func change(array *[4]int) {
	array[0] = 10
}*/

/*
func playerGen(name string) func() (string, int) {
	//血量一直为150
	hp := 150
	//返回创建的闭包
	return func() (string, int) {
		//将变量引用到闭包中
		return name, hp
	}
}*/

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

/*type User[T int | string] interface {
	Name() T
}*/

/*type myInt int

type AllBaseType interface {
	~int | float64 | float32 | int64
}

func Min[E AllBaseType, T []E](x E, y T) E {
	for _, v := range y {
		if x < v {
			return x
		}
	}
	return y[0]
}
func main() {

	fmt.Println(Min(1, []int{1, 2, 3}))
	//var x float64 = 1.2
	//var y float64 = 2.3
	//fmt.Println(Min(x, y))
}*/

// 12-并发

/*var x int = 1

// 任务编排
var w sync.WaitGroup
*/
/*
func main() {
	start := time.Now().UnixMilli()
	w.Add(2)
	go task1()
	go task2()
	w.Wait()
	end := time.Now().UnixMilli() - start
	fmt.Printf("总共用时：%d \n", end)

}

// func task2(w *sync.WaitGroup) {
func task2() {
	time.Sleep(time.Second * 5)
	fmt.Println("task2执行")
	w.Done()
}

// func task1(w *sync.WaitGroup) {
func task1() {
	time.Sleep(time.Second * 3)
	fmt.Println("task1执行")
	w.Done()
}*/

/*func hello(i int) {
	fmt.Println("Hello Goroutine!", i)
	fmt.Println(x + 1)
}

func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 2)
}
*/

/*func main() {

	//无缓冲  1有缓冲了 队列有长度 1
	//	var ch = make(chan int, 1)
	//	// 发送
	//	ch <- 10
	//	// 接收
	//	x := <-ch
	//	fmt.Println(x)
	//	close(ch)
	//	x = <-ch

	var ch = make(chan int)
	go recv(ch)
	//发送
	ch <- 10
}

func recv(ch chan int) {
	//接收
	x := <-ch
	fmt.Println(x)
}*/

/*func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	select {
	case data, ok := <-ch:
		if ok {
			fmt.Println("接收到数据: ", data)
		} else {
			fmt.Println("能并已被关闭")
		}
	case <-time.After(2 * time.Second):
		fmt.Println("超时了！")
	}
}*/
/*
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
*/

// map 100
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	lock.Lock() //加互斥锁
	//rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) //假设读操作耗时10毫秒
	//rwlock.Unlock()                   //解写锁
	lock.Unlock() // 解互斥锁
	wg.Done()
}

func read() {
	lock.Lock() //加互斥锁
	//rwlock.RLock()               //加读锁
	time.Sleep(time.Millisecond) //假设读操作耗时1毫秒
	//rwlock.RUnlock()             //解读锁
	lock.Unlock() //解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
