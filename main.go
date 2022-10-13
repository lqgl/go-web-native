package main

import "net/http"

func main() {
	// 对任意路径的请求进行响应
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	// 设置 web 服务器
	http.ListenAndServe("localhost:8080", nil) // DefaultServeMux
}
