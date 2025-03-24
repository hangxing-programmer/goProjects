package basic

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main09() {
	hash := md5.New()
	hash.Write([]byte("Hello World"))
	sum := hash.Sum([]byte(""))
	fmt.Printf("%x\n", sum)

	hash = sha1.New()
	hash.Write([]byte("Hello World"))
	sum = hash.Sum([]byte(""))
	fmt.Printf("%x", sum)
}
