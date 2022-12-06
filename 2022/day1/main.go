// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func getElfCals(arr []byte, idxs []int) {
// 	curCalIdx := 0
// 	for i, idx := range idxs {
// 		fmt.Println("Starting Elf %v", i)
// 		split := bytes.Index(arr[curCalIdx:idx-1], []byte{13, 10})
// 		if split != -1 {
// 			cal := string(arr[curCalIdx : curCalIdx+split])
// 			fmt.Println("Cur cal is: %v", cal)
// 			curCalIdx = split + 2
// 		}
// 	}

// }

// func findElves(arr []byte) []int {
// 	elfIdx := 0
// 	var elvs = []int{}
// 	for ix := 0; ix < len(arr); ix++ {
// 		split := bytes.Index(arr[elfIdx:], []byte{13, 10, 13, 10})

// 		if split != -1 {
// 			split += elfIdx
// 			elvs = append(elvs, split-1)
// 			fmt.Println("Found elf at idx: %v", split)
// 			elfIdx = split + 4
// 		}
// 	}

// 	return elvs
// }

// func main() {

// 	content, err := os.ReadFile("../inputs/day1.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//Find all the elf indices
// 	elvIndices := findElves(content)

// 	getElfCals(content, elvIndices)

// }

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var elves [][]int

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if scanner.Text() == "" {
				elves = append(elves, result)
				result = nil
				continue
			}
			return elves, err
		}
		result = append(result, x)
	}
	return elves, scanner.Err()
}

var elfCalMap map[int]int

func findMaxElfCals() {
	maxCal := 0
	elfCalMap = make(map[int]int)
	for i, cals := range elves {
		result := 0
		for _, v := range cals {
			result += v
		}
		if maxCal < result {
			maxCal = result
		}
		elfCalMap[i] = result
	}

	// fmt.Println("map:", elfCalMap)
	// fmt.Println("Max:", maxCal)

	keys := make([]int, 0, len(elfCalMap))

	for key := range elfCalMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return elfCalMap[keys[i]] < elfCalMap[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, elfCalMap[k])
	}
}

func main() {

	content, err := os.ReadFile("../inputs/day1.txt")
	ints, err := ReadInts(strings.NewReader(string(content)))
	fmt.Println(ints, err)

	findMaxElfCals()
}
