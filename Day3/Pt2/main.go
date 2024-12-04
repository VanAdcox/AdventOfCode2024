package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"os"
	"bufio"
)


func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	doState := true
	for scanner.Scan() {
		num, afterLineDoState := LineMulValue(scanner.Text(), doState)
		total += num
		doState = afterLineDoState
	}
	fmt.Println(total)

}


func LineMulValue(line string, initDoState bool) (int, bool) {
	pattern := regexp.MustCompile(`(mul\()(\d{1,3},\d{1,3})\)|(do\(\))|(don't\(\))`)
	matches := pattern.FindAllString(line, -1)
	sum := 0
	isDo := initDoState
	for _, rawMatch := range matches {
		if strings.Contains(rawMatch, "do()") {
			isDo = true
			continue
		}
		if strings.Contains(rawMatch, "don't()") {
			isDo = false
			continue
		}
		if isDo {
			sum += mulFromString(rawMatch)
		}
	}
	return sum, isDo
}

func mulFromString(rawMatch string) int {
		match := strings.TrimPrefix(strings.TrimSuffix(rawMatch, ")"), "mul(")

		numbers := strings.Split(match, ",")
		firstNum, _ := strconv.Atoi(numbers[0])
		secondNum, _ := strconv.Atoi(numbers[1])

		return firstNum * secondNum
}
