package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//resp, _ := http.Get("http://localhost:8888/go")
	resp, _ := http.Get("http://localhost:8888/login?username=zs&password=123456")
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		} else {
			log.Printf("读取到的内容===>\n%s\n", string(buf[:n]))
			break
		}

	}

}
