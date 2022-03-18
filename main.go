package main

import (
	"fmt"
	"reflect"
)

type TestReflections struct {
	id    string `dummy:"ID"`
	name  string `dummy:"NAME"`
	email string `dummy:"EMAIL"`
	foo   int    `dummmy:"FOO"`
}

func main() {
	var test TestReflections

	ty := reflect.TypeOf(test)

	for i := 0; i < ty.NumField(); i++ {
		d, ok := ty.Field(i).Tag.Lookup("dummy")
		if ok {
			fmt.Println(d)
		}
	}
}
