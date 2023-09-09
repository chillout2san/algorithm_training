package main

import (
	"fmt"
	"strings"
)

// https://atcoder.jp/contests/abc081/tasks/abc081_a
func main() {
	var in1 int

	_, err := fmt.Scan(&in1)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Count(fmt.Sprintf("%d", in1), "1"))
}
