package util

import (
	"fmt"
	"testing"
)


func before()  {
	fmt.Println("this is before func")
}
func after() {
	fmt.Println("this after func")
}
func add()  {
	fmt.Println("this add func")
}
func TestInterceptorFunc(t *testing.T) {
	if add,err := Interceptor(add, before, after); err == nil {
		add()
	} else {
		fmt.Println(err)
	}
}
