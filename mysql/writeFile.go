package mysql

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(d1 string) {
	f, err3 := os.Create("./output3.txt") //创建文件
	check(err3)
	defer f.Close()
	n2, err3 := f.WriteString(d1) //写入文件(字节数组)
	check(err3)
	fmt.Printf("写入 %d 个字节n", n2)
	// n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	// fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()
}
