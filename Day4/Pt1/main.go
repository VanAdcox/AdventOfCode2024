package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	originalCrossword := strings.Split(input, "\n")
	/*
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {

		}
	*/

	total := 0
	crossword := make([]string, len(originalCrossword))
	copy(crossword, originalCrossword)

	for y := 0; y < len(originalCrossword); y++ {
		for x := 0; x < len(originalCrossword[y]); x++ {
			newCrossword, wordFound := wasWordFound(crossword, string(crossword[y][x]), y, x)
			if wordFound {
				total += 1
			}
			crossword = newCrossword
		}
	}
	fmt.Println(total)
}
func wasWordFound(crossword []string, curLetter string, y, x int) ([]string, bool) {
	
}
