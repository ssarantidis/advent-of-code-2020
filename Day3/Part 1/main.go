package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readAndProcess(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentPosition := 0
	var resultArray []string
	scanner.Scan() // skip first line
	for scanner.Scan() {
		// string to int
		i := scanner.Text()
		lineLength := len(i)
		fmt.Println("Processing:", i)
		if (currentPosition + 3) <= (lineLength - 1) {
			// shift position in the same line
			currentPosition += 3
			fmt.Println("The current position is: ", currentPosition)
		} else {
			// shift position to the next line
			currentPosition = (currentPosition + 3) - lineLength
			fmt.Println("The current position after abs is: ", currentPosition)
		}
		// Save each character in the resultArray
		resultArray = append(resultArray, string([]rune(i)[currentPosition]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return resultArray
}

func main() {
	//Read the file
	result := readAndProcess("input.txt")
	resultString := strings.Join(result, "")
	fmt.Println("The number of trees I would encounter is: ", strings.Count(resultString, "#"))
}
