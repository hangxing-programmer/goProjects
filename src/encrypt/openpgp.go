package encrypt

import (
	"bytes"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func OpenpgpEncrypt() {
	// 1. 读取公钥
	pubKey, err := os.ReadFile("public.key")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 解析公钥
	keyring, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(pubKey))
	if err != nil {
		log.Fatal(err)
	}

	// 3. 创建加密文件
	encryptedFile, err := os.Create("encrypted.pgp")
	if err != nil {
		log.Fatal(err)
	}
	defer encryptedFile.Close()

	// 4. 使用 PGP 加密
	plaintext, err := os.ReadFile("plaintext.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 5. 写入加密数据
	armorWriter, err := armor.Encode(encryptedFile, "PGP MESSAGE", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer armorWriter.Close()

	pgpWriter, err := openpgp.Encrypt(armorWriter, keyring, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer pgpWriter.Close()

	_, err = pgpWriter.Write(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File encrypted with PGP!")
}

func OpenpgpDecrypt() {
	// 1. 加载私钥（假设私钥文件是 `private.key`，且受密码保护）
	privateKeyFile, err := os.ReadFile("private.key")
	if err != nil {
		log.Fatal("读取私钥失败:", err)
	}

	// 2. 解析私钥
	keyring, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(privateKeyFile))
	if err != nil {
		log.Fatal("解析私钥失败:", err)
	}

	// 3. 解密时需要私钥的密码（如果有密码保护）
	password := "your-private-key-password" // 替换为你的私钥密码
	key := keyring[0].PrivateKey
	if key.Encrypted {
		err = key.Decrypt([]byte(password))
		if err != nil {
			log.Fatal("私钥解密失败:", err)
		}
	}

	// 4. 打开加密的 PGP 文件
	encryptedFile, err := os.Open("encrypted.pgp")
	if err != nil {
		log.Fatal("打开加密文件失败:", err)
	}
	defer encryptedFile.Close()

	// 5. 解码 PGP Armor 格式（自动识别加密数据块）
	block, err := armor.Decode(encryptedFile)
	if err != nil {
		log.Fatal("解码 PGP Armor 失败:", err)
	}

	// 6. 用私钥解密
	md, err := openpgp.ReadMessage(block.Body, keyring, nil, nil)
	if err != nil {
		log.Fatal("读取加密消息失败:", err)
	}

	// 7. 读取解密后的明文
	plaintext, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		log.Fatal("读取解密内容失败:", err)
	}

	// 8. 输出或保存明文
	err = os.WriteFile("decrypted.txt", plaintext, 0644)
	if err != nil {
		log.Fatal("写入解密文件失败:", err)
	}

	log.Println("PGP 解密成功！解密内容已保存到 decrypted.txt")
}
