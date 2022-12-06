package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var m = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

func pointsForResult(oppMove, myMove string) int {

	if (oppMove == "A" && myMove == "Y") ||
		(oppMove == "B" && myMove == "Z") ||
		(oppMove == "C" && myMove == "X") {
		return m[myMove] + 6
	} else if m[oppMove] == m[myMove] {
		return 3 + m[myMove]
	} else {
		return m[myMove]
	}
}

func chooseResult(oppMove, myMove string) string {

	if myMove == "X" {
		switch oppMove {
		case "A":
			return "Z"
		case "B":
			return "X"
		case "C":
			return "Y"
		}
	}
	if myMove == "Y" {
		switch oppMove {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	}
	if myMove == "Z" {
		switch oppMove {
		case "A":
			return "Y"
		case "B":
			return "Z"
		case "C":
			return "X"
		}
	}
	return "A"
}

func scoreRound(round []byte) int {
	oppMove := string(round[0])
	myMove := string(round[2])

	myMove = chooseResult(oppMove, myMove)

	return pointsForResult(oppMove, myMove)

}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	idx := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		result += scoreRound(line)
		fmt.Printf("Day %v = %v\n", idx, result)
		idx++
		// x, err := strconv.Atoi(scanner.Text())
		// if err != nil {
		// 	if scanner.Text() == "" {
		// 		result = nil
		// 		continue
		// 	}
		// 	return result, err
		// }
		//result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {

	content, _ := os.ReadFile("../inputs/day2.txt")
	totalScore, _ := ReadInts(strings.NewReader(string(content)))
	fmt.Println("Total score:", totalScore)

}
