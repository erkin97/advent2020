package main

import (
	"advent2020/day1"
	"fmt"
)

func main() {
	ans, err := day1.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
