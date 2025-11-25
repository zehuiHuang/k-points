package mode

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	options := LogOptions{
		C: "c",
	}

	var optionList []option
	optionList = append(optionList, WithA("a"), WithB("b"))

	for _, v := range optionList {
		v.apply(&options)
	}

	fmt.Println(options.A, options.B)
}

func TestOptions2(t *testing.T) {
	bigclass := NewBigClass(WithAA("a"), WithBB("b"))
	print(bigclass)
}

func WithA(a string) option {
	return newFuncOptions(func(options *LogOptions) {
		options.A = a
	})
}

func WithB(b string) option {
	return newFuncOptions(func(options *LogOptions) {
		options.B = b
	})
}
