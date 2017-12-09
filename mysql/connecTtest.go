package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectTest() {
	db, err := sql.Open("mysql", "root:mysqlsybercare@tcp(192.168.1.204:3306)/cj.sybercare.com_2020?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM studio_person")
	checkErr(err)

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err")
	}
}
