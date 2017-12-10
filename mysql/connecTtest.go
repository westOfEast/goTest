package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectTest() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/hospital_kaifa?charset=utf8")
	checkErr(err)

	// select * from information_schema.columns
	// where table_schema = 'db'  #表所在数据库
	// and table_name = 'tablename' ; #你要查的表
	// rows, err := db.Query("SELECT * FROM data_blood_glucose where id < 10")
	// rows, err := db.Query("desc data_blood_glucose")
	rows, err := db.Query("select * from information_schema.columns where table_schema = 'hospital_kaifa' and table_name = 'data_blood_glucose' ")
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
				for _, value := range record[columns[i]] {
					fmt.Printf("%c", value)
				}
				fmt.Println("/n")
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
