package main

import "fmt"

func main() {
	messages := make(chan string)
	var f = func() {
		messages <- "李意嘉"
	}
	go f()
	msg := <-messages
	fmt.Println(msg)
}
