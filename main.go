package main

import "fmt"

type model struct {
	msg string

}

 

func main() {
	m := model{msg: "Hello, World!"}
	fmt.Println(m.msg)
	// Your code here
}