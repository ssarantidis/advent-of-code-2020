package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readAndProcess(file string) map[int][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	myMap := make(map[int][]string)

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\r\n")

	passportIndex := 0
	for _, line := range lines {
		if line != "" {
			fields := strings.Fields(line)
			myMap[passportIndex] = append(myMap[passportIndex], fields...)
		} else {
			// Found a blank line, increase the passportIndex by 1
			passportIndex++
		}
	}

	return myMap
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(v, str) {
			return true
		}
	}

	return false
}

func main() {
	arrayMandatoryKeysPassport := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validPassports := 0
	result := readAndProcess("input.txt")
	for _, value := range result {
		validPassport := true
		for _, v := range arrayMandatoryKeysPassport {
			if !contains(value, v) {
				validPassport = false
				break
			}
		}
		if validPassport == true {
			validPassports++
		}
	}
	fmt.Println("The total number of valid passports are [", validPassports, "]")
}
