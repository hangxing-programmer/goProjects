package basic

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main12() {
	r := mux.NewRouter()

	// 定义一个处理函数，用于处理 /hello 路径的 GET 请求
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, MUX!"))
	}).Methods("GET")

	// 定义一个处理函数，用于处理 /users/{id} 路径的 GET 请求，并获取路径参数 id
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		w.Write([]byte("User ID: " + id))
	}).Methods("GET")

	// 添加一个中间件，用于打印请求日志
	r.Use(loggingMiddleware)

	//启动一个 HTTP 服务器并在发生错误时记录错误并终止程序
	log.Fatal(http.ListenAndServe(":8081", r))
}

// 中间件函数，打印请求日志
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		//确保了请求在经过这个中间件记录后，继续沿着处理链传递给后续的处理程序进行实际的业务逻辑处理
		next.ServeHTTP(w, r)
	})
}
