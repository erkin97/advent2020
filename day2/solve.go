package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

//Solve day2
func Solve() (int, error) {
	inputFile, err := os.Open("day2/input")
	if err != nil {
		return 0, err
	}
	defer inputFile.Close()

	var tests []string
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		tests = append(tests, text)
	}

	// parse with regex
	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): (.+)")
	ans := 0
	for i := 0; i < len(tests); i++ {
		match := re.FindStringSubmatch(tests[i])
		left, _ := strconv.ParseInt(match[1], 10, 32)
		right, _ := strconv.ParseInt(match[2], 10, 32)
		testRune := rune(match[3][0])
		password := match[4]

		for pos, char := range password {
			if char == testRune {
				if pos == int(left-1) {
					left = -1
				} else if pos == int(right-1) {
					right = -1
				}
			}
		}

		if left*right < 0 {
			ans++
		}
	}

	return ans, nil
}
