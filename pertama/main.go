package main

import (
	"fmt"
	"pertama/lib/calculation"
)

func main() {
	fmt.Println("Halo Dunia")

	result := calculation.Add(8, 9)

	fmt.Println(result)
}
