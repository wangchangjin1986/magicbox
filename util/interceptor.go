package util

import "reflect"

func Interceptor(f interface{}, before func(), after func()) (func(args ...interface{}), error) {
	v := reflect.ValueOf(f)
	argsNum := v.Type().NumIn()
	params := make([]reflect.Value, argsNum)

	return func(args ...interface{}){
		if before != nil {
			before()
		}
		for i:=0;i<argsNum;i++ {
			params[i] = reflect.ValueOf(args[i])
		}
		v.Call(params)
		if after != nil {
			after()
		}
	}, nil
}
