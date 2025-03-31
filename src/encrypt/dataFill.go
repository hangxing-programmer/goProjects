package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func main() {
	// 原始数据(不一定是16字节的倍数)
	originalData := []byte("test")

	// 32字节的密钥(AES-256)
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	// 加密
	encrypted, err := encrypt(originalData, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后数据: %x\n", encrypted)

	// 解密
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后数据: %s\n", decrypted)

	// 验证原始数据和解密后数据是否一致
	if !bytes.Equal(originalData, decrypted) {
		panic("解密后的数据与原始数据不匹配!")
	}
	fmt.Println("验证成功: 解密后的数据与原始数据完全一致")
}

// PKCS#7填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS#7去填充
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7: 输入数据为空")
	}

	padding := int(data[length-1])
	// 检查所有填充字节是否正确
	for i := 0; i < padding; i++ {
		if data[length-padding+i] != byte(padding) {
			return nil, errors.New("pkcs7: 填充内容无效")
		}
	}

	return data[:length-padding], nil
}

// 加密函数
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充原始数据
	paddedPlaintext := pkcs7Pad(plaintext, aes.BlockSize)

	// 创建随机IV
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 使用CBC模式加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)

	return ciphertext, nil
}

// 解密函数
func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("密文长度小于块大小")
	}

	// 提取IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 检查密文长度是否是块大小的倍数
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("密文长度不是块大小的倍数")
	}

	// 使用CBC模式解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去除填充
	plaintext, err := pkcs7Unpad(ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
