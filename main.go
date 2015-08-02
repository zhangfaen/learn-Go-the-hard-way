package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	fmt.Println(reflect.TypeOf(slice))
	switch reflect.TypeOf(slice).Kind() {

	case reflect.Ptr:
		v := reflect.ValueOf(slice)
		fmt.Println(v.Elem().Len())
		for i, j := 0, v.Elem().Len()-1; i < j; i, j = i+1, j-1 {
			fmt.Println("here")
			fmt.Println(v.Elem().Index(i).Interface())
			t := v.Elem().Index(i).Interface()
			v.Elem().Index(i).Set(reflect.ValueOf(v.Elem().Index(j).Interface()))
			v.Elem().Index(j).Set(reflect.ValueOf(t))
		}

	case reflect.Slice:
		v := reflect.ValueOf(slice)
		fmt.Println("slice:", v)

	}
}

func main() {
	s := []int{1, 2, 3}
	Reverse(&s)
	fmt.Println(s)
	//println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
