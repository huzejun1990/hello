// @Author huzejun 2024/7/5 21:20:00
package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	//1.监听端口 IP地址 你的电脑会有很多程序 需要一份上端口来标识某一个程序
	// 127.0.0.1 本地环回地址	本机IP  0.0.0.0 广播地址 任意地址
	//listen, err := net.Listen("tcp", "127.0.0.1:20000")
	//listen, err := net.Listen("tcp", "192.168.31.108:20000")
	//listen, err := net.Listen("tcp", "0.0.0.0:20000")
	listen, err := net.Listen("tcp", ":20000")
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
