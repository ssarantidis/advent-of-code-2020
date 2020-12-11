package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

// contains checks if a string is present in a slice and valid
func containsAndValid(passportFields []string, key string, regex string) bool {
	isValid := true
	isPresent := false
	for _, v := range passportFields {
		if strings.Contains(v, key) {
			isPresent = true
			value := strings.Split(v, ":")
			fmt.Println("#########################")
			fmt.Println("The value to check is: ", value[1])
			match, _ := regexp.MatchString(regex, value[1])
			if !match {
				isValid = false
			}
			fmt.Println(value[0], "is valid? ", isValid)
		}
	}
	return isPresent && isValid
}

func main() {
	var keysToRegexMap map[string]string = map[string]string{"byr": "^(19[2-9][0-9]|20[0][0-2])$", "iyr": "^(201[0-9]|2020)$", "eyr": "^(202[0-9])|(2030)$", "hgt": "^(1[5-8][0-9]|19[0-3])cm$|^([5][9]|[6][0-9]|[7][0-6])in$", "hcl": "^#([0-9a-f]{6})$", "ecl": "^(amb|blu|brn|gry|grn|hzl|oth)$", "pid": "^([0-9]{9})$"}
	validPassports := 0
	result := readAndProcess("input.txt")
	for _, value := range result {
		validPassport := true
		for key, regex := range keysToRegexMap {
			if !containsAndValid(value, key, regex) {
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
