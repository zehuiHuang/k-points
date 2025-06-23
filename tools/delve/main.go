package delve

import "fmt"

func main() {
	a := 3
	b := 10
	c := Foo(a, b)
	fmt.Println(c)
}

func Foo(step, count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += step
	}
	return sum
}
