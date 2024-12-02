package main

import (
	//"os"
	"bufio"
	"crypto/x509"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	*/
	file := `1 2 3 4 5
5 6 7 8 9`
	scanner := bufio.NewScanner(strings.NewReader(file))
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		numsRaw := strings.Split(line, " ")

		curNum, _ := strconv.Atoi(numsRaw[0])
		for _, n := range numsRaw[1:] {
			num, _ := strconv.Atoi(n)
			if num != curNum {
				if Abs(num - curNum) < 3 && Abs(num - curNum) > 0 {
					curNum = num
					continue
			}
			}
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
