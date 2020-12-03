package day1

import (
	"bufio"
	"os"
	"strconv"
)

//Solve day1
func Solve() (int, error) {
	inputFile, err := os.Open("day1/input")
	if err != nil {
		return 0, err
	}
	defer inputFile.Close()

	var nums []int
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			return 0, err
		}
		nums = append(nums, int(number))
	}

	// solve using 3 sum with hashtable
	hashTable := map[int]bool{}
	sum := 2020

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			c := sum - a - b
			if hashTable[c] == true {
				return a * b * c, nil
			}
		}
		hashTable[a] = true
	}

	return 0, nil
}
