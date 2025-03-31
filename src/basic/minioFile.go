package basic

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"os"
)

func cli() {
	// 配置 MinIO 客户端
	endpoint := "10.0.11.201:9000"
	accessKey := "admin"
	secretKey := "adminadmin"
	useSSL := false // 如果使用 HTTP 协议，设置为 false

	// 初始化 MinIO 客户端
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("Failed to create MinIO client:", err)
		return
	}

	// 桶名称和对象名称
	bucketName := "zhhhh"
	objectName := "1-3GB.zip"
	localFilePath := "D:\\temp\\1-3GB.zip" // 本地保存路径

	// 获取对象
	obj, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("Failed to get object:", err)
		return
	}
	defer obj.Close()

	// 创建本地文件
	file, err := os.Create(localFilePath)
	if err != nil {
		fmt.Println("Failed to create local file:", err)
		return
	}
	defer file.Close()

	// 将对象内容写入到本地文件
	if _, err := io.Copy(file, obj); err != nil {
		fmt.Println("Failed to save object to file:", err)
		return
	}

	fmt.Println("Object downloaded and saved to:", localFilePath)
}
