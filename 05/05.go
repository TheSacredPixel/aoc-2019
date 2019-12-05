//Fisrt, add to your Intcode computer two new instructions:
//	Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
//	Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.

//Second, you'll need to add support for parameter modes:
//	Parameter mode 0, position mode, which causes the parameter to be interpreted as a position -
//		if the parameter is 50, its value is the value stored at address 50 in memory.
//	Now, your ship computer will also need to handle parameters in mode 1, immediate mode.
//		In immediate mode, a parameter is interpreted as a value - if the parameter is 50, its value is simply 50.

//ABCDE
// 1002

//DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
//                                  omitted due to being a leading zero
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	input = 5
)

var (
	intCodes       []int
	paramCount     = []int{0, 3, 3, 1, 1, 2, 2, 3, 3}
	modeParamCount = []int{0, 2, 2, 0, 1, 2, 2, 2, 2}
)

func main() {
	//1. After providing 1 to the only input instruction and passing all the tests,
	//what diagnostic code does the program produce?

	//2. Your computer is only missing a few opcodes:
	//	Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	//  Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	//  Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	//  Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	// What is the diagnostic code for system ID 5?
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	currInt := ""
	for scanner.Scan() {
		if scanner.Text() == "," || scanner.Text() == "\n" {
			i, _ := strconv.Atoi(currInt)
			intCodes = append(intCodes, i)
			currInt = ""
		} else {
			currInt += scanner.Text()
		}
	}

	_ = run(intCodes)
}

func interpretParameters(codes []int, opcode int, i int) []int {
	modes := strconv.Itoa(codes[i])

	if len(modes) == 1 {
		modes = ""
	} else {
		modes = modes[:len(modes)-2]
	}
	values := []int{}

	for j := 1; j <= modeParamCount[opcode]; j++ {
		pos := len(modes) - 1*j
		if pos < 0 || modes[pos] == '0' {
			values = append(values, codes[codes[i+j]])
		} else {
			values = append(values, codes[i+j])
		}
	}
	return values
}

func run(intCodes []int) []int {
	codes := make([]int, len(intCodes))
	copy(codes, intCodes)

	for i := 0; i < len(codes); {
		if codes[i] == 99 {
			break
		}
		opcode := codes[i] % 100
		vals := interpretParameters(codes, opcode, i)
		switch opcode {
		case 1:
			codes[codes[i+3]] = vals[0] + vals[1]
		case 2:
			codes[codes[i+3]] = vals[0] * vals[1]
		case 3:
			codes[codes[i+1]] = input
		case 4:
			fmt.Println(vals[0])
		case 5:
			if vals[0] != 0 {
				i = vals[1]
				continue
			}
		case 6:
			if vals[0] == 0 {
				i = vals[1]
				continue
			}
		case 7:
			if vals[0] < vals[1] {
				codes[codes[i+3]] = 1
			} else {
				codes[codes[i+3]] = 0
			}
		case 8:
			if vals[0] == vals[1] {
				codes[codes[i+3]] = 1
			} else {
				codes[codes[i+3]] = 0
			}
		default:
			fmt.Printf("unknown opcode %d\n", codes[i])
			os.Exit(1)
		}
		i += paramCount[opcode] + 1
	}
	return codes
}
