// package main
//
// import (
//
//	"fmt"
//	"os"
//	"time"
//
// )
//
//	func startGCTask(stop chan bool) {
//		t := time.NewTicker(5 * time.Second)
//		defer t.Stop()
//		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
//		for {
//			select {
//			case <-t.C:
//				fmt.Println("循环业务...", time.Now().Format("2006-01-02 15:04:05"))
//			case <-stop:
//				break
//			}
//		}
//	}
//
//	func main() {
//		file, err := os.OpenFile("D:\\vol2\\a.txt", os.O_RDWR|os.O_CREATE, 0644)
//		if err != nil {
//			fmt.Println("打开文件失败:", err)
//			return
//		}
//		defer file.Close()
//
//		_, err = file.Write([]byte("test"))
//		if err != nil {
//			fmt.Println("写入文件失败:", err)
//		} else {
//			fmt.Println("写入成功")
//		}
//		//metaStore, _ := basic.NewMetaStore("10.0.11.101:2379")
//		//err := metaStore.ExecuteTxn(func(txn *transaction.KVTxn) error {
//		//	pathKey := basic.GetPathKey("T02", "test1")
//		//	get, err := txn.Get(context.Background(), pathKey)
//		//	if err != nil {
//		//		return err
//		//	}
//		//	fmt.Println("get = ", string(get))
//		//	objKey := basic.GetObjKey("T02", string(get), "test")
//		//	fmt.Println("objKey = ", string(objKey))
//		//	err = txn.Delete(objKey)
//		//	if err != nil {
//		//		return err
//		//	}
//		//	fmt.Println("succeed...")
//		//	return nil
//		//})
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
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
//	}
package main

import (
	"abc/src/basic"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tikv/client-go/v2/txnkv/transaction"
	"time"
)

type Obj struct {
	Uuid string `json:"uuid"`
}
type LockMeta struct {
	LockTime    int64 `json:"lockTime"`
	MaxDuration int64 `json:"maxDuration"`
}

func main() {

	metaStore, _ := basic.NewMetaStore("10.0.11.33:2379")
	_ = metaStore.ExecuteTxn(func(txn *transaction.KVTxn) error {
		//txn.Delete([]byte("OS/T03/1010121741588318174/Data/Lock/b3b5a095-3c52-4b15-8b4e-f4486e0a1a58"))
		ctx := context.Background()
		get, _ := txn.Get(ctx, []byte("OS/T03/1010121741588318174/Data/LS/1-3GB.zip"))
		var obj Obj
		err := json.Unmarshal(get, &obj)
		if err != nil {
			fmt.Println(err)
			return err
		}
		lookMeta := &LockMeta{
			LockTime:    time.Now().UnixMilli(),
			MaxDuration: 1 * time.Minute.Milliseconds(),
		}
		bytes, err := json.Marshal(lookMeta)
		if err != nil {
			fmt.Println("json marshal err:", err)
			return err
		}
		err = txn.Set([]byte("OS/T03/1010121741588318174/Data/Lock/"+obj.Uuid), bytes)
		if err != nil {
			fmt.Println("set failed, err:", err)
			return err
		}
		return nil
	})
	//// 配置 MinIO 客户端
	//endpoint := "10.0.11.201:9000"
	//accessKey := "admin"
	//secretKey := "adminadmin"
	//useSSL := false // 如果使用 HTTP 协议，设置为 false
	//
	//// 初始化 MinIO 客户端
	//client, err := minio.New(endpoint, &minio.Options{
	//	Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
	//	Secure: useSSL,
	//})
	//if err != nil {
	//	fmt.Println("Failed to create MinIO client:", err)
	//	return
	//}
	//
	//// 桶名称和对象名称
	//bucketName := "zhhhh"
	//objectName := "1-3GB.zip"
	//localFilePath := "D:\\temp\\1-3GB.zip" // 本地保存路径
	//
	//// 获取对象
	//obj, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	//if err != nil {
	//	fmt.Println("Failed to get object:", err)
	//	return
	//}
	//defer obj.Close()
	//
	//// 创建本地文件
	//file, err := os.Create(localFilePath)
	//if err != nil {
	//	fmt.Println("Failed to create local file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//// 将对象内容写入到本地文件
	//if _, err := io.Copy(file, obj); err != nil {
	//	fmt.Println("Failed to save object to file:", err)
	//	return
	//}
	//
	//fmt.Println("Object downloaded and saved to:", localFilePath)
}
