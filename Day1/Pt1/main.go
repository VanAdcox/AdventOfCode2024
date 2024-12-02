package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"slices"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list_a []int
	var list_b []int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])
		list_a = append(list_a, a)
		list_b = append(list_b, b)
	}

	slices.Sort(list_a)
	slices.Sort(list_b)

	sum := 0
	for i, a := range list_a {
		sum += Abs(a - list_b[i])
	}
	fmt.Print(sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
