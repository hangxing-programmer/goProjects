package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")
	writer.Write([]byte(username + "\n" + password))
}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("url:%s,method:%s,header:%v", request.URL.Path, request.Method, request.Header)
	writer.Write([]byte("hello world"))
}

//		//	r := mux.NewRouter()
//		//
//		//	// 定义一个处理函数，用于处理 /hello 路径的 GET 请求
//		//	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//		//		w.Write([]byte("Hello, MUX!"))
//		//	}).Methods("GET")
//		//
//		//	// 定义一个处理函数，用于处理 /users/{id} 路径的 GET 请求，并获取路径参数 id
//		//	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
//		//		vars := mux.Vars(r)
//		//		id := vars["id"]
//		//		w.Write([]byte("User ID: " + id))
//		//	}).Methods("GET")
//		//
//		//	// 添加一个中间件，用于打印请求日志
//		//	r.Use(loggingMiddleware)
//		//
//		//	//启动一个 HTTP 服务器并在发生错误时记录错误并终止程序
//		//	log.Fatal(http.ListenAndServe(":8081", r))
//		//}
//
//		// 中间件函数，打印请求日志
//		//func loggingMiddleware(next http.Handler) http.Handler {
//		//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//		log.Println(r.Method, r.URL.Path)
//		//		next.ServeHTTP(w, r)
//		//	})
