package main

import (
	"fmt"
	"test/subject"
)

func Hello() {
	fmt.Println("hello go!")
	subject.Test()
	mytest()
}

func mytest() {
	fmt.Println("this is my test!")
}
