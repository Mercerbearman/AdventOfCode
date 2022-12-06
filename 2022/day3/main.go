package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var priorities = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 1 + 26,
	'B': 2 + 26,
	'C': 3 + 26,
	'D': 4 + 26,
	'E': 5 + 26,
	'F': 6 + 26,
	'G': 7 + 26,
	'H': 8 + 26,
	'I': 9 + 26,
	'J': 10 + 26,
	'K': 11 + 26,
	'L': 12 + 26,
	'M': 13 + 26,
	'N': 14 + 26,
	'O': 15 + 26,
	'P': 16 + 26,
	'Q': 17 + 26,
	'R': 18 + 26,
	'S': 19 + 26,
	'T': 20 + 26,
	'U': 21 + 26,
	'V': 22 + 26,
	'W': 23 + 26,
	'X': 24 + 26,
	'Y': 25 + 26,
	'Z': 26 + 26,
}

var idx = 0

func findItemsinCommon(comp1 string, comp2 string) string {
	retVal := ""
	for _, v := range comp1 {
		if strings.Contains(comp2, string(v)) {
			retVal += string(v)
		}

	}
	return retVal
}

func findItemInThree(line1 string, line2 string, line3 string) byte {

	commonItems := findItemsinCommon(line1, line2)

	token := "\n"
	if strings.Contains(string(commonItems), token) {
		fmt.Println("Found a weird character in my search.")
	}

	return findItemInBoth(line3, commonItems)
}

func findItemInBoth(comp1 string, comp2 string) byte {
	for _, v := range comp1 {
		if strings.Contains(comp2, string(v)) {
			idx++
			return byte(v)
		}

	}
	return 0

}

// func findItemInBoth(comp1 []byte, comp2 []byte) byte {
// 	str := string(comp2)
// 	for _, v := range comp1 {
// 		if strings.Contains(str, string(v)) {
// 			idx++
// 			return v
// 		}

// 	}
// 	return 0

// }

func readLine2(line []byte) ([]byte, []byte) {

	lenBuff := len(line) / 2
	comp1 := line[:lenBuff]
	comp2 := line[lenBuff:]

	fmt.Println("line: ", string(line))
	fmt.Println("Compart Two: ", string(comp1))
	fmt.Println("Compart Two: ", string(comp2))

	return comp1, comp2

}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	// part1 := false
	//idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		// if part1 {
		// 	comp1, comp2 := readLine2(line)
		// 	iteminBoth := findItemInBoth(comp1, comp2)
		// 	result += priorities[iteminBoth]
		// } else {
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		itemsInAll := findItemInThree(line, line2, line3)
		result += priorities[itemsInAll]
		// }
	}
	return result, scanner.Err()
}

func main() {

	content, _ := os.ReadFile("../inputs/day3.txt")
	totalScore, _ := ReadInts(strings.NewReader(string(content)))
	fmt.Println("Total score:", totalScore)

}
