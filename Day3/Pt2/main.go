/*
# What I Learned From Day 3
- Basic RegEx! I've never had a need for it before now and after an hour writing
my own terrible implementation for parsing the file I decided it was time to learn.
Showed me the importance of the right tool for the job isn't neccesarrily the one
youre most experienced with.
- Get a clear and accurate understanding of the requirements. Most of my time spent
was wasted not realizing that the do() or don't() state carries between lines,
something I'd known if I'd read
*/

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
