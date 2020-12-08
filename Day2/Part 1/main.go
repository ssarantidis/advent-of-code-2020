package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var occurences int = 0

func readFilePerLine(file string) int {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// string to int
		i := scanner.Text()
		fmt.Println("Processing:", i)
		if isPasswordValid(i) {
			occurences++
			fmt.Println("The current count is : [", occurences, "]")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return occurences
}

func isPasswordValid(s string) bool {
	slicelvl1 := strings.Split(s, ": ")
	password := slicelvl1[1]
	fmt.Println("The password is:[", password, "]")

	slicelvl2 := strings.Split(slicelvl1[0], " ")
	policyLetter := slicelvl2[1]
	fmt.Println("The policyLetter is:[", policyLetter, "]")

	slicelvl3 := strings.Split(slicelvl2[0], "-")
	minOccur, err := strconv.Atoi(slicelvl3[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	maxOccur, err := strconv.Atoi(slicelvl3[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("The min  is:[", minOccur, "]")
	fmt.Println("The max  is:[", maxOccur, "]")

	numOccur := strings.Count(password, policyLetter)
	fmt.Println("The numOccur of letter ", policyLetter, " is: [", numOccur, "]")

	if (numOccur >= minOccur) && (numOccur <= maxOccur) {
		return true
	}

	return false
}

func main() {
	//Read the file
	result := readFilePerLine("input.txt")
	fmt.Printf("The final result is = %v", result)
}
