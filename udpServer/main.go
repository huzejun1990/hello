// @Author huzejun 2024/7/5 22:42:00
package main

import (
	"log"
	"net"
)

func main() {

	listenUDP, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer listenUDP.Close()
	for {
		var data [1024]byte
		n, addr, err := listenUDP.ReadFromUDP(data[:])
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("服务端接收到的数据：%s,addr=%s \n", string(data[:n]), addr)
		_, err = listenUDP.WriteToUDP(data[:n], addr)
		if err != nil {
			log.Println(err)
		}
	}
}
