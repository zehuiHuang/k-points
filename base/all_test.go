package base

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"go-learn/base/method"
	"testing"
)

func TestName3(t *testing.T) {
	addFunc := addGenerator()
	fmt.Println(addFunc(1)) // 输出 1
	fmt.Println(addFunc(2)) // 输出 3
	fmt.Println(addFunc(3)) // 输出 6
}

func TestName(t *testing.T) {
	msg := "Hello, World!"
	deferFunc := doLater(msg)
	defer deferFunc()
	msg = "Hello, World2222!"
	fmt.Println("Doing something...")
}

// 延迟处理时的指针参数，修改时，会被修改
func TestName2(t *testing.T) {
	p := &person{
		name: "huangzehui",
		age:  34,
	}
	deferFunc := doLater2(p)
	defer deferFunc()
	p.name = "hzh"
	fmt.Println("Doing something...")
}

func TestName4(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	forEach(numbers, func(num int) {
		fmt.Println("Number:", num)
	})
}

func TestName5(t *testing.T) {
	f := newCounter()
	f1 := newCounter()
	fmt.Println(f())  //1
	fmt.Println(f1()) //1
	fmt.Println(f())  //2
}

func TestName22(t *testing.T) {
	method.T.F1(method.T{Name: "eggo"})
}

var context = `{ "ak": "xxx", "sk": "xxxxxxxxxxx","name":"iivv" }`

//var context = `[{"key":"ak","value":"aksss"},{"key":"sk","value":"xxxx"},{"key":"adr","value":["url1","url2"]}]`

type Config struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func TestName1111(t *testing.T) {
	var c []Config
	err := json.Unmarshal([]byte(""), &c)
	fmt.Println(err)

}
func ConvertTypeViaJSON[T any](interfaceJson any) (T, error) {
	tmp, err := sonic.Marshal(interfaceJson)
	var result T
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(tmp, &result)
	return result, nil
}

func A(i int) {
	i++
	fmt.Println(i)
}
func B() {
	f1 := A
	f1(1)
}
func C() {
	f2 := A
	f2(1)
}
func TestNamexx(t *testing.T) {
	B()
	C()
}

func create() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}
func TestName666(t *testing.T) {
	a := make([]int, 0, 10)
	fmt.Println(a)
	fs := create()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
}

func TestName777(t *testing.T) {
	example()
}
