package reflect

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("kind ", k)

	if t.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("FieldName:", t.Field(i).Name, "FiledType:", t.Field(i).Type,
			"FiledValue:", v.Field(i))
	}
}
