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
	if queryResult == nil || len(queryResult) <= 0 || queryResult[0]["TABLE_NAME"] == "" {
		return &Error{"Data Error", nil}
	}

	// tableName, ok := queryResult[""]
	f, err3 := os.Create("./" + queryResult[0]["TABLE_NAME"] + ".java") //创建文件
	check(err3)
	defer f.Close()

	writeLn("import javax.persistence.Column;", f)
	writeLn("import javax.persistence.Entity;", f)
	writeLn("import javax.persistence.Table;", f)
	//@Entity
	// @Table(name = "studio_users")
	writeLn("", f)
	writeLn("@Entity", f)
	writeLn(`@Table(name = "`+queryResult[0]["TABLE_NAME"]+`")`, f)
	f.WriteString("public class " + queryResult[0]["TABLE_NAME"] + " extends BaseInteger {\r\n")
	for i := 0; i < len(queryResult); i++ {

		// _, err3 := f.WriteString("    private String " + queryResult[i]["COLUMN_NAME"] + ";") //写入文件(字节数组)
		writeLn("", f)
		//  @Column(name = "studio_id", columnDefinition = ("varchar(255)  default null comment '工作室表Id'"))
		writeLn(`@Column(name = "`+queryResult[i]["COLUMN_NAME"]+`", columnDefinition = ("varchar(255)  default null comment '`+queryResult[i]["COLUMN_COMMENT"]+"'", f)
		writeLn("    private String "+queryResult[i]["COLUMN_NAME"]+";", f)
		// for key, value := range queryResult[i] {
		// 	// _, err3 := f.WriteString("    " + key + " = " + value) //写入文件(字节数组)
		// 	check(err3)
		// }
	}

	f.WriteString("}")
	f.Sync()
	return nil
}

func writeLn(code string, f *os.File) {
	_, err := f.WriteString(code + "\r\n")
	check(err)
}
