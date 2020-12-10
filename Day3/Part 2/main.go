package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readAndProcess(file string, rightHop int, downHop int) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentPosition := 0
	var resultArray []string

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")

	for i := downHop; i < len(lines); i += downHop {
		lineLength := len(strings.TrimRight(lines[i], "\r\n"))
		fmt.Println("Index #", i)
		fmt.Println("Processing:", lines[i])
		if (currentPosition + rightHop) <= (lineLength - 1) {
			// shift position in the same line
			currentPosition += rightHop
			fmt.Println("The current position is: ", currentPosition)
		} else {
			// shift position to the next line
			currentPosition = (currentPosition + rightHop) - lineLength
		}
		// Save each character in the resultArray
		resultArray = append(resultArray, string([]rune(lines[i])[currentPosition]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return resultArray
}

func main() {
	finalResult := 1
	downHopsArray := []int{1, 1, 1, 1, 2}
	rightHopsArray := []int{1, 3, 5, 7, 1}
	var treesOccurArray []int

	for repetition := 0; repetition < len(downHopsArray); repetition++ {
		result := readAndProcess("input.txt", rightHopsArray[repetition], downHopsArray[repetition])
		resultString := strings.Join(result, "")
		treesOccurArray = append(treesOccurArray, strings.Count(resultString, "#"))
	}
	fmt.Println("The treesOccurArray are: ", treesOccurArray)

	for _, value := range treesOccurArray {
		finalResult *= value
	}

	fmt.Println("The final result is: ", finalResult)
}
