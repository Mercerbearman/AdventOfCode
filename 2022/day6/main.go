package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var chars = make(map[byte]int)

func allUnique(line []byte) bool {

	keys := make(map[byte]bool)
	list := []byte{}

	for _, entry := range line {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return len(line) == len(list)

}

func ReadLine(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		//fmt.Println(line)

		for idx := 0; idx < len(line); idx++ {
			//Grab a slice
			chars := line[idx : idx+14]

			if allUnique(chars) {
				print(string(chars), idx+14)
				break
			} else {
				continue
			}

		}
	}
	return result, scanner.Err()
}

func main() {

	//var totalScore int
	runTest := false

	if runTest {
		content, _ := os.ReadFile("../tests/day6.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))

		//fmt.Println(string(content))
	} else {
		content, _ := os.ReadFile("../inputs/day6.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))
		//fmt.Println(string(content))
	}

}
