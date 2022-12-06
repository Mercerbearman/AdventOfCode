package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func splitLine2Ways(line, substring string) (one, two string) {
	strs := strings.Split(line, substring)
	one = strs[0]
	two = strs[1]
	return one, two
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		//Now we need to split the line 2 ways.
		tasks1, tasks2 := splitLine2Ways(line, ",")
		taskAminS, taskAmaxS := splitLine2Ways(tasks1, "-")
		taskBminS, taskBmaxS := splitLine2Ways(tasks2, "-")

		taskAmin, _ := strconv.Atoi(taskAminS)
		taskAmax, _ := strconv.Atoi(taskAmaxS)
		taskBmin, _ := strconv.Atoi(taskBminS)
		taskBmax, _ := strconv.Atoi(taskBmaxS)

		// lenA := taskAmax - taskAmin
		// lenB := taskBmax - taskBmin

		// part1 := false

		if taskAmin >= taskBmin && taskAmin <= taskBmax {
			result += 1
		} else if taskBmin >= taskAmin && taskBmin <= taskAmax {
			result += 1
		}

	}
	return result, scanner.Err()
}

func main() {

	runTest := false

	if runTest {
		content, _ := os.ReadFile("../tests/day4.txt")
		totalScore, _ := ReadInts(strings.NewReader(string(content)))

		expectedScore := 2

		if totalScore != expectedScore {
			fmt.Printf("Test failed. Actual %v, expected %v\n", totalScore, expectedScore)

		} else {
			fmt.Printf("Test passed. Actual %v, expected %v\n", totalScore, expectedScore)
		}
	} else {
		content, _ := os.ReadFile("../inputs/day4.txt")
		totalScore, _ := ReadInts(strings.NewReader(string(content)))

		fmt.Println("Total score:", totalScore)
	}

}
