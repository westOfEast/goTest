package main

import (
	"fmt"
	. "test/mysql"
)

func main() {
	// fmt.Println("Hello word")
	// fmt.Println("hello mytest")
	// Test()
	// test()
	var strTest = string("汉字，测试")
	fmt.Println(strTest)
	// Hello()
	ConnectTest()
}

func test() {
	fmt.Println("hello mytest")
}
