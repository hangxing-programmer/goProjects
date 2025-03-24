package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	//对Hash操作
	//加数据(单个用HSet,多个用HMSet)
	_, err = conn.Do("HSet", "user01", "name", "Tom")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Do("HSet", "user01", "age", 24)
	if err != nil {
		log.Fatalln(err)
	}
	//设置有效时间
	_, err = conn.Do("expire", "name", 10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("操作成功...")
	//读取(单个用HGet,多个用HMGet,且使用redis.Strings()进行转换类型,注意加's'!!!)
	res, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		log.Fatalln(err)
	}
	res01, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("读取结果==>", res, res01)
}
