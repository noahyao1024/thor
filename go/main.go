package main

import (
	"fmt"
	"golib/server"
)

func main() {
	fmt.Println("hello world")
	server.Listen("9000")
}
