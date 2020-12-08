package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFilePerLine(file string) int {
	m := make(map[int]string)

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// string to int
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Processing:", i)
		if check(i, m) {
			return i * (2020 - i)
		}
		m[i] = "DUMMY"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}

func check(i int, m map[int]string) bool {
	if _, ok := m[2020-i]; ok {
		//do something here
		return true
	}
	return false
}

func main() {
	//Read the file
	result := readFilePerLine("input.txt")
	fmt.Printf("The final result is =%v", result)
}
