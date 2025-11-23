package base

import "fmt"

//泛型和多态

func printValue[T any](value T) {
	fmt.Printf("Value: %v\n", value)
}
