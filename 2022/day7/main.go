package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type file struct {
	name string
	size int
}

type directoryNode struct {
	name  string
	files []file
}

type directory struct {
	name      string
	childDirs []directoryNode
	totalSize int
}

func ReadLine(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := 0
	//Root
	scanner.Scan()
	root := directory{name: "/"}
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
	return result, scanner.Err()
}

func main() {
	//var totalScore int
	runTest := true

	if runTest {
		content, _ := os.ReadFile("../tests/day7.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))

		//fmt.Println(string(content))
	} else {
		content, _ := os.ReadFile("../inputs/day7.txt")
		_, _ = ReadLine(strings.NewReader(string(content)))
		//fmt.Println(string(content))
	}
}
