package main

import (
	"reflect"
)

type MyType struct {
	a string
}

func main() {
	t := MyType{a: "test"}
	state := New(reflect.TypeOf("s"), reflect.TypeOf(t))
	state.Set(reflect.ValueOf("test"), reflect.ValueOf(t))
}
