package main

import (
	"fmt"
	"reflect"
)

// 用反射的方法创建一个类型 T 的实例（零值）

func NewInstance[T any]() T {
	typ := TypeOf[T]()

	switch typ.Kind() {
	case reflect.Map:
		return reflect.MakeMap(typ).Interface().(T)
	case reflect.Slice, reflect.Array:
		return reflect.MakeSlice(typ, 0, 0).Interface().(T)
	case reflect.Ptr:
		typ = typ.Elem()
		origin := reflect.New(typ)
		inst := origin

		for typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			inst = inst.Elem()
			inst.Set(reflect.New(typ))
		}

		return origin.Interface().(T)
	default:
		var t T
		return t
	}
}

func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
func main() {
	type A struct {
		x, y int
	}
	a := NewInstance[A]()
	fmt.Println(a)
}
