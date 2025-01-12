package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, _ := range nums {
		fmt.Println(i)
	}

	fmt.Println("-----------------------")

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, ",", v)
	}

	for k, _ := range kvs {
		fmt.Println(k)
	}

	for _, v := range kvs {
		fmt.Println(v)
	}

	for i, c := range "go" {
		println(i, ",", c)
	}
}
