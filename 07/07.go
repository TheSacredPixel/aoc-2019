//There are five amplifiers connected in series; each one receives an input signal and produces an output signal.
//They are connected such that the first amplifier's output leads to the second amplifier's input,
//the second amplifier's output leads to the third amplifier's input, and so on.
//The first amplifier's input value is 0, and the last amplifier's output leads to your ship's thrusters.

//The Elves have sent you some Amplifier Controller Software (your puzzle input), a program that should run on your existing Intcode computer.
//Each amplifier will need to run a copy of the program.

//When a copy of the program starts running on an amplifier, it will first use an input instruction
//to ask the amplifier for its current phase setting (an integer from 0 to 4).
//Each phase setting is used exactly once.
//The program will then call another input instruction to get the amplifier's input signal,
//compute the correct output signal, and supply it back to the amplifier with an output instruction.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	intCodes       []int
	paramCount     = []int{0, 3, 3, 1, 1, 2, 2, 3, 3}
	modeParamCount = []int{0, 2, 2, 0, 1, 2, 2, 2, 2}
	results        []int
	wg             sync.WaitGroup
)

func main() {
	//1. Try every combination of phase settings on the amplifiers.
	//   What is the highest signal that can be sent to the thrusters?
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

	fmt.Println(getBestSignal([5]int{0, 1, 2, 3, 4}))
}

func getBestSignal(phases [5]int) int {
	wg.Add(120)
	generate(5, &phases)

	wg.Wait()
	highest := 0
	for _, res := range results {
		if res > highest {
			highest = res
		}
	}
	return highest
}

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generate(k int, phases *[5]int) {
	if k == 1 {
		go testPhases(*phases)
		return
	}
	generate(k-1, phases)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			swap(phases, i, k-1)
		} else {
			swap(phases, 0, k-1)
		}
		generate(k-1, phases)
	}
}

func swap(arr *[5]int, pos1, pos2 int) {
	hold := (*arr)[pos1]
	(*arr)[pos1] = (*arr)[pos2]
	(*arr)[pos2] = hold
}

func testPhases(phases [5]int) {
	ph := &phases

	signal := 0
	for _, phase := range (*ph) {
		_, signal = run([]int{phase, signal})
	}

	results = append(results, signal)
	wg.Done()
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

func run(inputs []int) ([]int, int) {
	codes := make([]int, len(intCodes))
	copy(codes, intCodes)
	inputCount := 0

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
			codes[codes[i+1]] = inputs[inputCount]
			inputCount++
		case 4:
			return codes, vals[0]
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
	return codes, 0
}
