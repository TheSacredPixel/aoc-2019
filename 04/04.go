//A password must meet the following criteria:
//-It is a six-digit number.
//-The value is within the range given in your puzzle input.
//-Two adjacent digits are the same (like 22 in 122345).
//-Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	val, end := parseRange(scanner.Text())

	//1. How many different passwords within the range given in your puzzle input meet the criteria?
	//2. How many different passwords within the range given in your puzzle input meet all of the criteria?
	success, singleSuccess := 0, 0

	for val <= end {
		digits := intToArray(val)

		pos, num := validateDigits(digits)
		if pos != -1 { //not valid
			//eg 478399 -> 478888
			for ; pos < 6; pos++ {
				digits[pos] = num
			}
			val = arrayToInt(digits)
		} else { //valid
			if hasDuplicate(digits) {
				success++
			}
			if hasSinglePair(digits) {
				singleSuccess++
			}
			val = arrayToInt(digits)
			val++
		}
	}
	fmt.Println(success)
	fmt.Println(singleSuccess)
}

func hasDuplicate(digits [6]int) bool {
	found := [10]bool{}
	for _, dig := range digits {
		if found[dig] {
			return true
		}
		found[dig] = true
	}
	return false
}

func hasSinglePair(digits [6]int) bool {
	found := [10]int{}
	for _, dig := range digits {
		found[dig]++
	}
	for _, num := range found {
		if num == 2 {
			return true
		}
	}
	return false
}

//convert int to array of digits
func intToArray(val int) [6]int {
	digits := [6]int{}
	valString := strconv.Itoa(val)
	valSplit := strings.Split(valString, "")
	for i, dig := range valSplit {
		digits[i], _ = strconv.Atoi(dig)
	}
	return digits
}

func arrayToInt(digits [6]int) int {
	str := ""
	for _, dig := range digits {
		str += strconv.Itoa(dig)
	}
	n, _ := strconv.Atoi(str)
	return n
}

//if invalid, returns digit position and what it should be
func validateDigits(digits [6]int) (int, int) {
	prevN := -1
	for i, n := range digits {
		if n < prevN {
			return i, prevN
		}
		prevN = n
	}
	return -1, -1
}

func parseRange(text string) (int, int) {
	vals := strings.Split(text, "-")
	start, _ := strconv.Atoi(vals[0])
	end, _ := strconv.Atoi(vals[1])
	return start, end
}
