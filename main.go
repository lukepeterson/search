package main

import (
	"fmt"

	"github.com/lukepeterson/search/binary"
)

func main() {
	data := []int{1, 2, 4, 5, 7, 8, 12}
	target := 5
	result := binary.BinarySearch(data, target)
	if result != -1 {
		fmt.Printf("Target %d found at index %d\n", target, result)
	} else {
		fmt.Println("Target not found")
	}
}
