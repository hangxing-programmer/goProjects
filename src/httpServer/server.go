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
