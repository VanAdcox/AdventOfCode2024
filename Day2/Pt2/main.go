/*
# What I learned from Day 2
The importance of a few testcases to understand how your code is working
How slices work under the hood, particularly in assignments like slice1 := slice2
*/


package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	/*
	Test data, all cases should pass
	file := `48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20
`
	scanner := bufio.NewScanner(strings.NewReader(file))
*/
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		numsRaw := strings.Split(line, " ")

		for i, _ := range numsRaw {
			rawCopy := append([]string(nil), numsRaw...)
			if IsValidLine(append(rawCopy[:i], rawCopy[i+1:]...)) {
				count += 1
				break
			}
		}
	}
	fmt.Println(count)
}



func IsValidLine(line []string) bool {
	lastNum, _ := strconv.Atoi(line[0])
	nextNum, _ := strconv.Atoi(line[1])
	isIncreasing := lastNum > nextNum

	for _, n := range line[1:] {
		nextNum, _ = strconv.Atoi(n)

		if (isIncreasing && lastNum < nextNum) || (!isIncreasing && lastNum > nextNum) {
			return false
		}

		if Abs(nextNum-lastNum) <= 3 && Abs(nextNum-lastNum) > 0 {
			lastNum = nextNum
		} else {
			return false
		}
	}
	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
