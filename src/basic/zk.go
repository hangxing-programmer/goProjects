package basic

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

func Init(servers []string) (*zk.Conn, error) {
	connect, _, err := zk.Connect(servers, time.Second*30)
	if err != nil {
		fmt.Println("Connect err:", err)
		return connect, err
	}
	return connect, nil
}

// 先创建，才能进行set 修改
func Create(path string, data []byte, conn *zk.Conn) {
	//path := "/test-node"
	//data := []byte("initial data")
	flags := 0                     // 0 表示普通节点，其他标志位可以参考 ZooKeeper 文档
	acl := zk.WorldACL(zk.PermAll) // 使用默认权限

	_, err := conn.Create(path, data, int32(flags), acl)
	if err != nil {
		fmt.Println("Failed to create node:", err)
		return
	}
}

//	func write(path, value string) {
//		if _, err := os.Stat(path); os.IsNotExist(err) {
//			// 文件不存在，创建文件
//			fmt.Println("File does not exist, creating...")
//			file, err := os.Create(path)
//			if err != nil {
//				fmt.Printf("Failed to create file: %v\n", err)
//				return
//			}
//			// 关闭文件
//			defer file.Close()
//		} else if err != nil {
//			// 其他错误
//			fmt.Printf("Error checking file: %v\n", err)
//			return
//		} else {
//			// 文件存在
//			fmt.Println("File already exists.")
//		}
//		err := os.WriteFile(path, []byte(value), os.ModePerm)
//		if err != nil {
//			fmt.Printf("write err:%v\n", err)
//		}
//	}
//
//	func read(path string) []byte {
//		bytes, err := os.ReadFile(path)
//		if err != nil {
//			fmt.Printf("read err:%v\n", err)
//		}
//		return bytes
//	}
func main() {
	//z, _ := Init([]string{"10.0.11.33:2181"})
	//res, _, err := z.Get("/oeos/Sys/Resource/Storage/StoragePool/LOCAL")
	//if err != nil {
	//	fmt.Println("Get err:", err)
	//	return
	//}
	//path := "D:\\temp\\zkTest.txt"
	//write(path, string(res))
	//bytes := read(path)
	//Create("/oeos/Sys/Resource/Storage/StorageComponent/LOSTO", bytes, z)
	//z.Set("/oeos/Sys/Resource/Storage/StorageComponent/LOSTO", bytes, -1)

}
