package main

import (
	"fmt"
)

// https://atcoder.jp/contests/abc086/tasks/abc086_a
func main() {
	var in1 int
	var in2 int

	_, err := fmt.Scan(&in1, &in2)
	if err != nil {
		panic(err)
	}

	if (in1*in2)%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
