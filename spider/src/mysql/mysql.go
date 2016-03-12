package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(sql string, param ...interface{}) int64 {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	res, err := db.Exec(sql, param...)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return id
}

func Update(sql string, param ...interface{}) bool {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		return false
	}
	_, err = db.Exec(sql, param...)
	if err != nil {
		return false
	}
	return true
}

func Query(sql string, param ...interface{}) map[string]interface{} {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rows, err := db.Query(sql, param)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	result := make(map[string]interface{})
	var array = make([]interface{}, 0)
	err = rows.Scan(array...)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i, val := range columns {
		result[val] = array[i]
	}

	return result
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:huanxiang@tcp(47.89.41.165:3306)/galaxygo?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
