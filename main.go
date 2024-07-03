// @Author huzejun 2024/7/3 22:09:00
package main

import (
	"fmt"
	"io/ioutil"
)

// 16-IO
func main() {
	//writeFile("B站学习之路")
	//writeFile("插入数据")
	//wr()
	re()
}

func wr() {
	err := ioutil.WriteFile("./yyy.txt", []byte("学习之路"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

/*func writeFile(context string) {
	file, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}
*/
/*
func writeFile(context string) {
	//file, err := os.Create("./test.txt")	//先创建
	// os.O_RDWR 打开模式
	// 655 权限控制
	//特殊权限位 拥有者位 同组用户位 其余用户位 wr = 6 rx = 5
	//w写 r读 x执行 w 2 r 4 x 1
	file, err := os.OpenFile("./test.txt", os.O_RDWR, 655)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Seek(4, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte(context))
	if err != nil {
		log.Fatal(err)
	}
}
*/
