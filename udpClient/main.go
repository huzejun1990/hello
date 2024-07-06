// @Author huzejun 2024/7/6 18:48:00
package main

import (
	"log"
	"net"
)

func main() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		log.Println(err)
	}
	defer udp.Close()
	sendData := []byte("hello upd")
	_, err = udp.Write(sendData)
	if err != nil {
		log.Println(err)
	}
	data := make([]byte, 4096)
	n, addr, err := udp.ReadFromUDP(data)
	if err != nil {
		log.Println(err)
	}
	log.Printf("客户端的接收到的数据：%s, addr=%s \n", string(data[:n]), addr)
}
