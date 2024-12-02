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
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		numsRaw := strings.Split(line, " ")
		if IsValidLine(numsRaw) {
			count += 1
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
