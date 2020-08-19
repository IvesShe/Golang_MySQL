package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// db是一個全局連接池對象
var db *sql.DB

func initDB() (err error) {

	// 使用者名稱:密碼@tcp(IP:端口號)/數據庫名稱
	dsn := "root:123456@tcp(192.168.183.128:3306)/goTest"
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"

	// 這邊不會校驗用戶名稱和密碼是否正確
	// dsn格式不正確的時候會報錯
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	// 這邊就會嘗試連接數據庫，並檢查用戶名稱及密碼是否正確
	err = db.Ping()
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
	id   int
	name string
	age  int
}

// 查詢單條記錄
func queryOne(id int) {
	var u1 user

	// 寫查詢單條記錄的sql語句
	sqlStr := `select id,name,age from user where id=?;`
	// 執行
	// 從連接池裡拿一個連接出來去數據庫查詢單條記錄
	//rowObj := db.QueryRow(sqlStr, 2)

	// 拿到結果
	// 必須對rowObj對象調用Scan方法，因為該方法會釋放數據庫鏈接
	//rowObj.Scan(&u1.id, &u1.name, &u1.age)

	// 也可以簡化成一行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)

	fmt.Printf("u1: %#v\n", u1)
}

// 查詢多條記錄
func queryMore(n int) {
	// 寫查詢單條記錄的sql語句
	sqlStr := `select id,name,age from user where id > ?;`
	// 執行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed , err:%v\n", sqlStr, err)
		return
	}

	// 一定要關閉rows
	defer rows.Close()

	// 循環取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed , err:%v\n", err)
		}
		fmt.Printf("u1: %#v\n", u1)
	}
}

// 插入數據
func insert() {
	// 寫SQL語句
	sqlStr := `insert into user(name,age) values("Doris",23)`

	// exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed , err:%v\n", err)
		return
	}

	// 如果是插入數據的操作，能夠拿到插入數據的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed , err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新數據
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed , err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed , err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行數據\n", n)
}

// 刪除數據
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed , err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed , err:%v\n", err)
		return
	}
	fmt.Printf("刪除了%d行數據\n", n)
}

// 預處理方式插入多條數據
func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed , err:%v\n", err)
		return
	}
	defer stmt.Close()

	//後續只需要拿到stmt去執行一些操作
	var m = map[string]int{
		"陳希": 37,
		"小芳": 23,
		"德雯": 21,
		"陳婷": 25,
	}
	for k, v := range m {
		// 後續只需要傳值
		stmt.Exec(k, v)
	}
}

// 事務操作
func transactionDemo() {
	// 開啟事務
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed , err:%v\n", err)
		return
	}

	// 執行多個SQL操作
	sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update xxx set age=age+2 where id=3`

	// 執行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 失敗要回滾
		tx.Rollback()
		fmt.Println("執行SQL1出錯了，要回滾!!")
		return
	}

	// 執行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 失敗要回滾
		tx.Rollback()
		fmt.Println("執行SQL2出錯了，要回滾!!")
		return
	}

	// 上面兩步SQL都執行成功，就提交本次事務
	err = tx.Commit()
	if err != nil {
		// 失敗要回滾
		tx.Rollback()
		fmt.Println("提交出錯了，要回滾!!")
		return
	}

	fmt.Println("事務執行成功!!")
}

func main() {
	err := initDB()

	if err != nil {
		fmt.Printf("init db failed, err%v\n", err)
		return
	}

	fmt.Printf("連接MySQL數據庫成功!!\n")

	//queryOne(2)
	//queryMore(0)
	//insert()
	//updateRow(33, 1)
	//deleteRow(5)
	//prepareInsert()
	transactionDemo()
	//defer db.Close()
}
