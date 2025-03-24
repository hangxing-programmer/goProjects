package basic

import (
	"fmt"
	miniogo "github.com/minio/minio-go"
	"io"
	"log"
	"os"
	"time"
)

const (
	endpoint        = "10.0.11.101:19011"
	accessKeyID     = "datacenteradmin"
	secretAccessKey = "datacenteradmin"
	bucketName      = "cyrd-vol2"
	objectName      = "10/10/1010121732155787593/data/20240901-6688330-d59901a8-12dd-43aa-865f-2f6ac5072e65.log"
)

func main01() {

	fmt.Printf("现在是%v", time.Now().Format("2006-01-02 15:04:05"))
	// objectName有 10/10 就会报错 The specified key does not exist
	opts := miniogo.GetObjectOptions{}
	core, _ := miniogo.NewCore(endpoint, accessKeyID, secretAccessKey, false)
	object, _, err := core.GetObject(bucketName, objectName, opts)
	if err != nil {
		log.Printf("GetObject err: %v\n", err)
	}
	//输出到控制台
	io.Copy(os.Stdout, object)
}
