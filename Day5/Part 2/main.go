package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readAndProcess(file string) []int {
	// Store the seat ids
	var sliceSeatIDs []int

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	// Parse the text by line
	lines := strings.Split(string(rawBytes), "\r\n")

	for _, line := range lines {
		fmt.Println("Processing line: [", line, "]")
		// Convert the line to a binary number
		var s string
		replacer := strings.NewReplacer("B", "1", "F", "0", "R", "1", "L", "0")
		s = replacer.Replace(line)
		fmt.Println("The string representation of line is [", s, "]")
		// Row calculation
		row, err := strconv.ParseInt(s[0:7], 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The int represenation is [", int(row), "]")
		// Column calculation
		col, err := strconv.ParseInt(s[7:], 2, 0)
		if err != nil {
			log.Fatal(err)
		}

		// Seat Id calculation
		sliceSeatIDs = append(sliceSeatIDs, int(row)*8+int(col))
	}

	return sliceSeatIDs
}

func main() {
	result := readAndProcess("input.txt")
	sort.Ints(result)
	max := result[len(result)-1]
	min := result[0]
	sizeResult := len(result)

	// will be the sum of the numbers in the result slice
	sum := 0
	for i := 0; i < len(result); i++ {
		sum += result[i]
	}
	fmt.Println("The sum of seat Ids is [", sum, "]")

	// the total sum of numbers between min and max - arimthmitiki proodos!
	total := ((sizeResult + 1) * (min + max)) / 2
	fmt.Println("The missing number is [", total-sum, "]")

}
