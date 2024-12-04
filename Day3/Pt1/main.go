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
	for scanner.Scan() {
		total += LineMulValue(scanner.Text())
	}
	fmt.Println(total)

}


func LineMulValue(line string) int {
	pattern := regexp.MustCompile(`(mul\()(\d{1,3},\d{1,3})\)`)
	matches := pattern.FindAllString(line, -1)
	sum := 0
	for _, rawMatch := range matches {
		match := strings.TrimPrefix(strings.TrimSuffix(rawMatch, ")"), "mul(")

		numbers := strings.Split(match, ",")
		firstNum, _ := strconv.Atoi(numbers[0])
		secondNum, _ := strconv.Atoi(numbers[1])
		sum += firstNum * secondNum
	}
	return sum
}
