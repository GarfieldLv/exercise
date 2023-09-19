package main

import (
	"fmt"
	"reflect"
)

type tag struct {
	name string `info:"name" doc:"myDoc"`
	age  int    `info:"age"`
}

func FindTag(arg interface{}) {
	typ := reflect.TypeOf(arg).Elem()
	for i := 0; i < typ.NumField(); i++ {
		tagInfo := typ.Field(i).Tag.Get("info")
		tagDoc := typ.Field(i).Tag.Get("doc")
		fmt.Printf("info: %s, doc: %s\n", tagInfo, tagDoc)
	}
}
func main() {
	t := tag{"zhangsan", 18}
	FindTag(&t)
}
