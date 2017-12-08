package main

import (
	"fmt"
	"test/subject"
)

func Hello() {
	// fmt.Println("hello go!")
	// subject.Test()
	// mytest()
	result := subject.NextMinNum(642531)
	fmt.Println("result is ", result)
}

func mytest() {
	fmt.Println("this is my test!")
}
