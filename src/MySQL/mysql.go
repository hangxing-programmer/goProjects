package MySQL

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //驱动
	"log"
)

type Contact struct {
	UserID      int
	ContactID   string
	ContactType string
	Status      string
	UserName    string
	Time1       string
	Time2       string
}

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/dishing")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer db.Close()

	//开始连接
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Successfully connected!")
	//查数据
	str := "select * from user_contact where contact_id = 1"
	//执行sql语句
	res, err := db.Exec(str)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("%T\n", res)
	//插入数据
	//str = "insert into user_contact(user_id,contact_id,contact_type,status,user_name) values (4,'1','0','1','zs')"
	//_, err = db.Exec(str)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//fmt.Println("Successfully inserted!")

	//批量处理
	//contacts := []Contact{
	//	{UserID: 5, ContactID: "1", ContactType: "0", Status: "1", UserName: "ls"},
	//	{UserID: 6, ContactID: "1", ContactType: "0", Status: "1", UserName: "ww"},
	//}
	////提高效率，预处理
	//stmt, err := db.Prepare("insert into user_contact (user_id,contact_id,contact_type,status,user_name) values (?,?,?,?,?)")
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//for _, v := range contacts {
	//	_, err = stmt.Exec(v.UserID, v.ContactID, v.ContactType, v.Status, v.UserName)
	//	if err != nil {
	//		log.Fatalln(err)
	//		return
	//	}
	//}
	//fmt.Println("Successfully batch inserted!")

	//单行查询
	per := Contact{}
	row := db.QueryRow("select * from user_contact where user_id = 4")
	err = row.Scan(&per.UserID, &per.ContactID, &per.ContactType, &per.Status, &per.Time1, &per.Time2, &per.UserName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("%+v\n", per)

	//多行查询
	rows, err := db.Query("select * from user_contact")
	if err != nil {
		log.Fatalln(err)
		return
	}
	m := make(map[int]Contact)
	for rows.Next() {
		err = rows.Scan(&per.UserID, &per.ContactID, &per.ContactType, &per.Status, &per.Time1, &per.Time2, &per.UserName)
		if err != nil {
			log.Fatalln(err)
			return
		}
		m[per.UserID] = per
		//fmt.Printf("%+v\n", per)
	}
	fmt.Printf("%+v\n", m)
}
