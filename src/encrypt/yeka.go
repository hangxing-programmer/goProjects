package encrypt

import (
	"fmt"
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
)

func YekaEncrypt(filename, password string) {
	zipFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating zip file", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 使用 ZipCrypto 加密
	w, err := zipWriter.Encrypt("secret.txt", password, zip.AES256Encryption)
	if err != nil {
		fmt.Println("error encrypt", err)
		return
	}
	_, err = w.Write([]byte("This is encrypted with ZipCrypto."))
	if err != nil {
		fmt.Println("error writing ZipCrypto", err)
		return
	}

	log.Println("ZIP file encrypted with ZipCrypto!")
}

func YekaDecrypt(filename, password string) {
	// 打开加密的 ZIP 文件
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Println("error opening zip:", err)
		return
	}
	defer zipReader.Close()

	// 设置密码
	//password = "secret" // 替换为 ZIP 文件的密码

	// 遍历 ZIP 中的文件
	for _, file := range zipReader.File {
		// 设置密码
		file.SetPassword(password)

		// 打开文件
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("error opening file:", err)
			return
		}
		defer fileReader.Close()

		// 创建目标文件
		targetFile, err := os.Create(file.Name)
		if err != nil {
			fmt.Println("error creating file:", err)
			return
		}
		defer targetFile.Close()

		// 复制内容
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			fmt.Println("error copying file:", err)
			return
		}

		fmt.Printf("Extracted: %s\n", file.Name)
	}

	log.Println("Extraction completed!")
}
