package main

import "fmt"

func main() {

	//s := "hello"
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%x ", s[i]) // 输出每个字节的十六进制表示
	//}

	s := "你好" // 在UTF-8编码中，"你"和"好"各占3个字节
	for _, r := range s {
		fmt.Println(string(r))
		fmt.Println(rune(r))
		fmt.Println(r)
		//fmt.Printf("%x ", rune(r)) // 输出每个rune的十六进制表示
	}

	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(rune(s[i]))
	}

	//a := "abcdd"
	//for i := range a {
	//	fmt.Println(string(a[i]))
	//}

	//ints := []int{1, 2}
	//ints = append(ints, 3, 4, 5)
	//fmt.Println(ints)

}
