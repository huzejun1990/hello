// @Author huzejun 2024/7/6 19:24:00
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	for {
		//n, err := resp.Body.Read(buf[:])
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		} else {
			log.Printf("读取到的内容:%s \n", string(buf[:n]))
			break
		}
	}
}
