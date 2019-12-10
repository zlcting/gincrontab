package database

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

func DBConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/bishe?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	return db
}

//操作数据库
func ModifyDB(sqlStr string, args ...interface{}) (int64, error) {
	db := DBConn()
	result, err := db.Exec(sqlStr, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sqlStr string) *sql.Rows {

	db := DBConn()
	rows, _ := db.Query(sqlStr)
	return rows
}
