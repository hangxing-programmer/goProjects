package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var pool *redis.Pool

// 初始化链接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //和数据库的最大链接数,0标识没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
func main() {
	conn := pool.Get() //取出一个链接
	defer conn.Close()

	s, err := conn.Do("Set", "name", "Jerry")
	if err != nil {
		log.Fatalln(err)
		return
	}
	//取出
	s, err = redis.String(conn.Do("Get", "name"))
	fmt.Println("===>", s)

}
