package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

func OpensslEncrypt(filename, dst string, key, iv []byte) {
	// 打开要加密的文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建加密文件
	encryptedFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer encryptedFile.Close()

	// 设置加密密钥和IV
	//key = []byte("1234567890123456") // 16/24/32 bytes (AES-128/192/256)
	//iv = []byte("1234567890123456")  // 16 bytes for AES-CBC

	// 创建AES块加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	// 创建加密流
	stream := cipher.NewCBCEncrypter(block, iv)

	// 使用缓冲区逐块读取和加密文件
	buffer := make([]byte, 1024*64) // 64KB缓冲区
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			continue
		}
		if n == 0 {
			break
		}

		// 加密当前块
		stream.CryptBlocks(buffer[:n], buffer[:n])

		// 写入加密文件
		_, err = encryptedFile.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}

	fmt.Println("File encrypted with AES-CBC!")
}

func OpensslDecrypt(filename, dst string, key, iv []byte) {
	// 打开加密文件
	encryptedFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer encryptedFile.Close()

	// 创建解密文件
	decryptedFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer decryptedFile.Close()

	// 设置解密密钥和IV
	//key := []byte("1234567890123456") // 16/24/32 bytes (AES-128/192/256)
	//iv := []byte("1234567890123456")  // 16 bytes for AES-CBC

	// 创建AES块解密器
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	// 创建解密流
	stream := cipher.NewCBCDecrypter(block, iv)

	// 使用缓冲区逐块读取和解密文件
	buffer := make([]byte, 1024*64) // 64KB缓冲区
	for {
		n, err := encryptedFile.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			continue
		}
		if n == 0 {
			break
		}

		// 解密当前块
		stream.CryptBlocks(buffer[:n], buffer[:n])

		// 写入解密文件
		_, err = decryptedFile.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}

	fmt.Println("File decrypted successfully!")
}
