package mysql

import (
	// "fmt"
	"os"
	. "test/common"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(queryResult [](map[string]string)) error {
	if queryResult == nil || len(queryResult) <= 0 {
		return &Error{"Data Error", nil}
	}
	// tableName, ok := queryResult[""]
	f, err3 := os.Create("./output3.txt") //创建文件
	check(err3)
	defer f.Close()
	for i := 0; i < len(queryResult); i++ {
		for key, value := range queryResult[i] {
			_, err3 := f.WriteString(key + " = " + value) //写入文件(字节数组)
			check(err3)
		}

		_, err := f.WriteString("\r\n")
		check(err)
	}
	f.Sync()
	return nil
}
