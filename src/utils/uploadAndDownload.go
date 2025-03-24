package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//解决跨域问题
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//获取上传文件
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Fatalln("获取文件失败...", err)
		return
	}
	defer file.Close()
	//创建目标文件
	create, err := os.OpenFile("D:\\userTest\\"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("创建文件失败...", err)
		return
	}
	_, err = io.Copy(create, file)
	if err != nil {
		log.Fatalln("文件上传失败...", err)
		return
	}
	fmt.Println("文件上传成功...")
	w.Write([]byte("success!"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	//打开目标文件
	file, err := os.Open("D:\\userTest\\userTest.png")
	if err != nil {
		log.Fatalln("打开文件失败...", err)
	}
	defer file.Close()
	//设置响应头
	w.Header().Set("Content-Disposition", "attachment; filename=\""+file.Name()+"\"")
	//解决跨域问题
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//将文件写入响应体
	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("文件下载成功...")
	w.Write([]byte("success!"))
}
func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe("localhost:8080", nil)
}
