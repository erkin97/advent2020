package main

import (
	"advent2020/day1"
	"advent2020/day2"
	"advent2020/day3"
	"fmt"
)

func main() {
	//day1
	ans, err := day1.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("day1: ", ans)
	}

	//day2
	ans, err = day2.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("day2: ", ans)
	}

	//day3
	ans, err = day3.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("day3: ", ans)
	}
}
