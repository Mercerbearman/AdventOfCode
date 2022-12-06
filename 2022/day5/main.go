package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var stacks [][]string

func moveItems(count, fromArr, toArr int) {

	fromArr -= 1
	toArr -= 1

	part1 := false
	if part1 {
		for i := 0; i < count; i++ {
			len1 := len(stacks[fromArr]) - 1
			itm := stacks[fromArr][len1:]
			stacks[fromArr] = stacks[fromArr][:len1]

			//len2 := len(stacks[toArr])
			stacks[toArr] = append(stacks[toArr], itm[0])
			//stacks[toArr][len2] = itm[0]
		}
	} else {
		// Cranes can move items multiple at a time.
		len1 := len(stacks[fromArr])
		itemsToMove := stacks[fromArr][len1-count:]
		stacks[fromArr] = stacks[fromArr][:len1-count]

		stacks[toArr] = append(stacks[toArr], itemsToMove...)

	}

}

func ReadLine(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		var count, fromArr, toArr int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &fromArr, &toArr)

		fmt.Println(count, fromArr, toArr)

		moveItems(count, fromArr, toArr)

	}
	return result, scanner.Err()
}

// [N]         [C]     [Z]
// [Q] [G]     [V]     [S]         [V]
// [L] [C]     [M]     [T]     [W] [L]
// [S] [H]     [L]     [C] [D] [H] [S]
// [C] [V] [F] [D]     [D] [B] [Q] [F]
// [Z] [T] [Z] [T] [C] [J] [G] [S] [Q]
// [P] [P] [C] [W] [W] [F] [W] [J] [C]
// [T] [L] [D] [G] [P] [P] [V] [N] [R]
//  1   2   3   4   5   6   7   8   9

func initialize() {
	stacks[0] = append(stacks[0], "T", "P", "Z", "C", "S", "L", "Q", "N")
	stacks[1] = append(stacks[1], "L", "P", "T", "V", "H", "C", "G")
	stacks[2] = append(stacks[2], "D", "C", "Z", "F")
	stacks[3] = append(stacks[3], "G", "W", "T", "D", "L", "M", "V", "C")
	stacks[4] = append(stacks[4], "P", "W", "C")
	stacks[5] = append(stacks[5], "P", "F", "J", "D", "C", "T", "S", "Z")
	stacks[6] = append(stacks[6], "V", "W", "G", "B", "D")
	stacks[7] = append(stacks[7], "N", "J", "S", "Q", "H", "W")
	stacks[8] = append(stacks[8], "R", "C", "Q", "F", "S", "L", "V")
}

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

func initializeTest() {
	stacks[0] = append(stacks[0], "Z", "N")
	stacks[1] = append(stacks[1], "M", "C", "D")
	stacks[2] = append(stacks[2], "P")
}

func printEndofEachStack() {
	result := ""
	for x := 0; x < len(stacks); x++ {
		lastItem := len(stacks[x]) - 1

		if lastItem < 0 {
			continue
		}

		result += stacks[x][lastItem]
	}

	fmt.Println(result)
}

func main() {

	//var totalScore int
	runTest := false

	stacks = make([][]string, 9)
	if runTest {
		initializeTest()
		content, _ := os.ReadFile("../tests/day5.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))

		//fmt.Println(string(content))
	} else {
		initialize()
		content, _ := os.ReadFile("../inputs/day5.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))
		//fmt.Println(string(content))
	}

	printEndofEachStack()

}
