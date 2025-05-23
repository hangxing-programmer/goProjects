package https

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

/*
*
生成证书链：

	根CA → 中间CA → 服务器证书（CommonName: localhost）。
	根CA → 客户端证书（可选，如果服务器要求客户端证书）。

启动服务器：

	加载 服务器证书 + 中间CA证书，并信任 根CA。

客户端请求：

	信任 根CA，访问 https://localhost。
*/
func main() {
	// 生成根证书
	//rootCer := "D:\\test\\public"
	//generateRootCA(rootCer)

	go func() {
		httpsServer()
	}()
	time.Sleep(3 * time.Second)

	sendReq()
}

func generateRootCA(filename string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Example Inc."},
			CommonName:   "Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 有效期10年
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	// 保存根证书
	certOut, err := os.Create(filename + ".crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certOut.Close()

	// 保存根证书私钥
	keyOut, err := os.Create(filename + ".key")
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
	keyOut.Close()
	midCer := "D:\\test\\mid"
	generateIntermediateCA(midCer, certBytes, privateKeyBytes)
	cliCer := "D:\\test\\cli"
	generateClientCert(cliCer, certBytes, privateKeyBytes)
}

func generateIntermediateCA(filename string, certBytes, rootKeyBytes []byte) {

	rootCert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		panic(err)
	}

	rootKey, err := x509.ParseECPrivateKey(rootKeyBytes)
	if err != nil {
		panic(err)
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"Example Inc."},
			CommonName:   "Intermediate CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 0, 0), // 有效期5年
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certBytes, err = x509.CreateCertificate(rand.Reader, &template, rootCert, &privateKey.PublicKey, rootKey)
	if err != nil {
		panic(err)
	}

	// 保存中间证书
	certOut, err := os.Create(filename + ".crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certOut.Close()

	// 保存中间证书私钥
	keyOut, err := os.Create(filename + ".key")
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
	keyOut.Close()
	serverCer := "D:\\test\\private"
	generateServerCert(serverCer, certBytes, privateKeyBytes)
}

func generateServerCert(filename string, intermediateCertBytes, intermediateKeyBytes []byte) {

	intermediateCert, err := x509.ParseCertificate(intermediateCertBytes)
	if err != nil {
		panic(err)
	}

	intermediateKey, err := x509.ParseECPrivateKey(intermediateKeyBytes)
	if err != nil {
		panic(err)
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{"Example Inc."},
			CommonName:   "localhost",
		},
		DNSNames:              []string{"localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(2, 0, 0), // 有效期2年
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, intermediateCert, &privateKey.PublicKey, intermediateKey)
	if err != nil {
		panic(err)
	}

	// 保存服务器证书
	certOut, err := os.Create(filename + ".crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certOut.Close()

	// 保存服务器证书私钥
	keyOut, err := os.Create(filename + ".key")
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
	keyOut.Close()
}

func generateClientCert(filename string, rootCertBytes, rootKeyBytes []byte) {
	rootCert, err := x509.ParseCertificate(rootCertBytes)
	if err != nil {
		panic(err)
	}

	rootKey, err := x509.ParseECPrivateKey(rootKeyBytes)
	if err != nil {
		panic(err)
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(4),
		Subject: pkix.Name{
			Organization: []string{"Example Inc."},
			CommonName:   "client",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCert, &privateKey.PublicKey, rootKey)
	if err != nil {
		panic(err)
	}

	// 保存客户端证书
	certOut, err := os.Create(filename + ".crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certOut.Close()

	// 保存客户端私钥
	keyOut, err := os.Create(filename + ".key")
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
	keyOut.Close()
}

func httpsServer() {
	// 加载服务器证书和私钥
	serverCertFile := "D:\\test\\private.crt"
	serverKeyFile := "D:\\test\\private.key"
	intermediateCAFile := "D:\\test\\mid.crt"

	// 读取服务器证书
	certPEM, err := ioutil.ReadFile(serverCertFile)
	if err != nil {
		log.Fatalf("读取服务器证书失败: %v", err)
	}

	// 读取中间CA证书
	intermediatePEM, err := ioutil.ReadFile(intermediateCAFile)
	if err != nil {
		log.Fatalf("读取中间CA证书失败: %v", err)
	}

	// 合并服务器证书和中间CA证书
	certPEM = append(certPEM, intermediatePEM...)

	// 读取私钥
	keyPEM, err := ioutil.ReadFile(serverKeyFile)
	if err != nil {
		log.Fatalf("读取私钥失败: %v", err)
	}

	// 加载证书和私钥
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatalf("加载证书失败: %v", err)
	}

	// 加载根CA证书（用于验证客户端证书）
	rootCAFile := "D:\\test\\public.crt"
	rootCA, err := ioutil.ReadFile(rootCAFile)
	if err != nil {
		log.Fatalf("加载根CA证书失败: %v", err)
	}
	clientCAs := x509.NewCertPool()
	clientCAs.AppendCertsFromPEM(rootCA)

	// 配置TLS
	config := &tls.Config{
		Certificates: []tls.Certificate{cert}, // 包含服务器证书 + 中间CA证书
		ClientCAs:    clientCAs,
		//ClientAuth:   tls.RequireAndVerifyClientCert, // 需要客户端证书
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   http.HandlerFunc(handleRequest),
		TLSConfig: config,
	}

	log.Println("HTTPS服务器启动，监听443端口")
	log.Fatal(server.ListenAndServeTLS("", "")) // 证书已通过 TLSConfig 加载
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, HTTPS with certificate chain!"))
}

func sendReq() {
	// 加载服务器根证书（如果服务器使用自签名证书）
	rootCer := "D:\\test\\public.crt"
	rootCert, err := ioutil.ReadFile(rootCer)
	if err != nil {
		log.Fatalf("加载根证书失败: %v", err)
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(rootCert)

	// 加载客户端证书和私钥（如果服务器需要客户端证书）
	//serverCer := "D:\\test\\cli"
	//clientCert, err := tls.LoadX509KeyPair(serverCer+".crt", serverCer+".key")
	//if err != nil {
	//	log.Fatalf("加载客户端证书失败: %v", err)
	//}

	// 配置TLS
	config := &tls.Config{
		RootCAs: rootCAs,
		//Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: false, // 不跳过证书验证
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	resp, err := client.Get("https://localhost")
	if err != nil {
		log.Fatalf("请求服务器失败: %v", err)
	}
	defer resp.Body.Close()

	log.Println("请求成功")
}
