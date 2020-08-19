// sqlx

package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// db是一個全局連接池對象
var db *sqlx.DB

func initDB() (err error) {

	// 使用者名稱:密碼@tcp(IP:端口號)/數據庫名稱
	dsn := "root:123456@tcp(192.168.183.128:3306)/goTest"
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"

	// 這邊不會校驗用戶名稱和密碼是否正確
	// dsn格式不正確的時候會報錯
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		return err
	}

	// 設置數據庫連接池的最大連接數
	db.SetMaxOpenConns(10)

	// 設置最大空閒連接數
	db.SetMaxIdleConns(5)

	return nil
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()

	if err != nil {
		fmt.Printf("init db failed, err%v\n", err)
		return
	}

	fmt.Printf("連接MySQL數據庫成功!!\n")

	sqlStr1 := `select id,name,age from user where id=6`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed , err:%v\n", err)
	}
	fmt.Printf("userList:%#v\n", userList)
}
