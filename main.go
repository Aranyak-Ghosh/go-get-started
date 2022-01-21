package main

import (
	"fmt"
	"time"
)

func hello()test chan bool{
	fmt.Println("Hello world goroutine")
	
}
func main() {
	x := go hello()
	y <- x
	fmt.Println("main function")
}
