package main

import (
	"fmt"
	name_test "net/http"
)

func main() {
	fmt.Println("Hello, Go Standard LIbraty")

	resp, err := name_test.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return

	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status: ", resp.Status)
}
