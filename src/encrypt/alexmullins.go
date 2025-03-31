package encrypt

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

func AlexmullinsEncrypt(filename, password string) {
	// 1. 创建 ZIP 文件
	zipFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer zipFile.Close()

	// 2. 创建 ZIP Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 3. 定义要加密的文件
	files := []struct {
		Name, Content string
	}{
		{"secret.txt", "This is a secret file."},
		{"config.json", `{"key": "value"}`},
	}

	// 4. 设置 ZIP 密码（AES-256）
	//password = "password"

	// 5. 写入文件（AES 加密）
	for _, file := range files {
		w, err := zipWriter.Encrypt(file.Name, password)
		if err != nil {
			fmt.Println("error encrypting file:", err)
			return
		}
		_, err = w.Write([]byte(file.Content))
		if err != nil {
			fmt.Println("error writing file:", err)
			return
		}
	}

	log.Println("ZIP file encrypted with AES-256!")
}

func AlexmullinsDecrypt(filename, password string) {
	// 1. 打开 ZIP 文件
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Println("error opening zip:", err)
		return
	}
	defer zipReader.Close()

	// 2. 设置密码
	//password = "mypassword"

	// 3. 遍历并解密文件
	for _, file := range zipReader.File {
		file.SetPassword(password) // 必须设置密码
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("error opening file:", err)
			continue
		}
		defer fileReader.Close()

		// 4. 创建目标文件
		targetFile, err := os.Create(file.Name)
		if err != nil {
			fmt.Println("error creating file:", err)
			continue
		}
		defer targetFile.Close()

		// 5. 解密并写入
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			fmt.Println("error copying file:", err)
			continue
		}
		fmt.Printf("Extracted: %s\n", file.Name)
	}
}
