// @Author huzejun 2024/7/5 22:18:00
package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//1.v建立连接	// 192.168.31.108
	//conn, err := net.Dial("tcp", "127.0.0.1:20000")
	conn, err := net.Dial("tcp", "192.168.31.108:20000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		readString, _ := input.ReadString('\n')
		inputInfo := strings.Trim(readString, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			break
		}
		_, err := conn.Write([]byte(readString))
		if err != nil {
			log.Fatal(err)
		}
		var buf [512]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("客户端接收到的内容：%s \n", string(buf[:n]))
	}
}
