package main

import (
	"fmt"
)

func named() (n, _ int) {
	return 1, 2
}
func question1() {
	fmt.Print(named())
}

func question2() {
	a := make([]int, 20)
	a = []int{7, 8, 9, 10}
	b := a[15:16]
	fmt.Println(b)
}

func question3() {
	s := []int{9, 8, 7}
	p := &s
	r := *p
	r[0] = 11
	fmt.Println(s[0])
}
func main() {
	question3()
}
