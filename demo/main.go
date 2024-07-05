//@Author huzejun 2024/7/3 22:09:00
/*package main

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
}*/

/*func re() {
	content, err := ioutil.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}*/

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

package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	//1.监听端口 IP地址 你的电脑会有很多程序 需要一份上端口来标识某一个程序
	// 127.0.0.1 本地环回地址	本机IP
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//接收请求 阻塞等待 建立连接
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		v := string(buf[:n])
		log.Printf("接收到客户端的信息：%s \n", v)
		conn.Write(buf[:n])
	}
}
