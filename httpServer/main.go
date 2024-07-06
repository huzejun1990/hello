// @Author huzejun 2024/7/6 19:05:00
package main

import (
	"log"
	"net/http"
)

func main() {
	//http http://127.0.0.1:8000/go 域名其实就是ip+端口 域名默认的端口是80 80可以不写
	// go不同的url,代表不同的请求 可以写不周的处理
	http.HandleFunc("/go", myHandler)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	log.Println(username)
	password := r.URL.Query().Get("password")
	log.Println(password)
	w.Write([]byte("登陆成功！"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url:%s,method:%d,header:%v \n", r.URL.Path, r.Method, r.Header)
	w.Write([]byte("学习之路 go教程"))
}
