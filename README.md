# Golang_MySQL

# 下載依賴

```shell
go get -u -v github.com/go-sql-driver/mysql
```

# 使用MySQL驅動

```go
func Open(driverName, dataSourceName string) (*DB, error)
```

# 建一個測試用的庫

```sql
CREATE DATABASE goTest;
use goTest;
```

# 建一個測試用的數據表

```sql
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

![image](./images/20200819132136.png)

# 插入一些數據
```sql
INSERT INTO `goTest`.`user` (`name`, `age`) VALUES ('IvesShe', '18'); 
INSERT INTO `goTest`.`user` (`name`, `age`) VALUES ('Jack', '30'); 
INSERT INTO `goTest`.`user` (`name`, `age`) VALUES ('ChiChi', '18'); 
INSERT INTO `goTest`.`user` (`name`, `age`) VALUES ('Alex', '25'); 
```

![image](./images/20200819132918.png)

# 查詢的語句

```sql
select * from user;
```

![image](./images/20200819132921.png)

```sql
select id,name,age from user where id=1
select id,name,age from user where id=3
select name,age from user where id=4;
```

![image](./images/20200819133434.png)


# 執行畫面

## 插入數據

![image](./images/20200819145521.png)

![image](./images/20200819145542.png)

## 更新數據

![image](./images/20200819150116.png)

![image](./images/20200819150138.png)

## 刪除數據

![image](./images/20200819151006.png)

![image](./images/20200819151033.png)

## 預處理

- 可以預先讓數據庫編譯，增加效能
- 防止sql注入

![image](./images/20200819152309.png)

![image](./images/20200819152349.png)

## 事務操作

- 全部成功提交，才會更新數據庫

![image](./images/20200819154358.png)

![image](./images/20200819154529.png)

- 只要其中出現失敗，即會回滾

![image](./images/20200819154740.png)

------

# sqlx使用

程式放置 ./sqlx/ 資料夾下

安裝
```shell
go get -u -v github.com/jmoiron/sqlx
```

![image](./images/20200819174442.png)

[參考文檔](https://www.liwenzhou.com/posts/Go/sqlx/)

------

# sql注入

## 任何時候都不應該自己拼接SQL語句

輸入以下字符串都可以引發SQL注入的問題
```shell
"xxx' or 1=1#"
"xxx' union select * from user #"
"xxx' and (select count(*) from user) <10 #"

```
