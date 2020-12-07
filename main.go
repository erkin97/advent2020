package main

import (
	"advent2020/day1"
	"advent2020/day2"
	"fmt"
)

func main() {
	ans, err := day1.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	ans, err = day2.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
