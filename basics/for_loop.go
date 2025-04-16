package main

import "fmt"

func main() {

	// Einfache Schleife
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Scheife und collection
	numbers := []int{1, 2, 3, 4, 5, 6}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
