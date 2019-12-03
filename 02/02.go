//An Intcode program is a list of integers separated by commas (like 1,0,0,3,99).
//To run one, start by looking at the first integer (called position 0).
//Here, you will find an opcode - either 1, 2, or 99. The opcode indicates what to do;
//for example, 99 means that the program is finished and should immediately halt.
//Encountering an unknown opcode means something went wrong.

//Opcode 1 adds together numbers read from two positions and stores the result in a third position. The three integers immediately after the opcode tell you these three positions - the first two indicate the positions from which you should read the input values, and the third indicates the position at which the output should be stored.

//For example, if your Intcode computer encounters 1,10,20,30, it should read the values at positions 10 and 20, add those values, and then overwrite the value at position 30 with their sum.

//Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. Again, the three integers after the opcode indicate where the inputs and outputs are, not their values.

//Once you're done processing an opcode, move to the next one by stepping forward 4 positions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sum      = 0
	intCodes []int
)

func main() {
	//Restore the gravity assist program (your puzzle input) to the "1202 program alarm" state
	//it had just before the last computer caught fire.
	//What value is left at position 0 after the program halts?
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
		if scanner.Text() != "," {
			currInt += scanner.Text()
		} else {
			i, _ := strconv.Atoi(currInt)
			intCodes = append(intCodes, i)
			currInt = ""
		}
	}
	//run with 1202 state
	memory := run(intCodes, 12, 2)
	fmt.Println(memory[0])

	//To complete the gravity assist, you need to determine what pair of inputs produces the output 19690720.
	//The inputs should still be provided to the program by replacing the values at addresses 1 and 2, just like before.
	//In this program, the value placed in address 1 is called the noun, and the value placed in address 2 is called the verb.
	//Each of the two input values will be between 0 and 99, inclusive.
loop:
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			memory = run(intCodes, n, v)
			if memory[0] == 19690720 {
				fmt.Println(100*n + v)
				break loop
			}
		}
	}
}

func run(intCodes []int, noun int, verb int) []int {
	codes := make([]int, len(intCodes))
	copy(codes, intCodes)

	codes[1] = noun
	codes[2] = verb

	for i := 0; i < len(codes); i += 4 {
		switch codes[i] {
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
		case 99:
			break
		default:
			fmt.Printf("unknown opcode %d\n", codes[i])
			os.Exit(1)
		}
	}
	return codes
}
