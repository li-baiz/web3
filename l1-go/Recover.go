package main

import "fmt"

func maypanic() {
	panic("a problem")
}

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	maypanic()
	fmt.Println("After mayPanic()")
}
