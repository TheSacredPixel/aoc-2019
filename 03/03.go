//Two wires are connected to a central port and extend outward on a grid.
//You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

//The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit,
//you need to find the intersection point closest to the central port.
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

	//1: What is the Manhattan distance from the central port to the closest intersection?
	//2: What is the fewest combined steps the wires must take to reach an intersection?
	scanner := bufio.NewScanner(file)

	wireOne := getWire(scanner)
	wireOnePath, _ := plotPath(wireOne, nil)

	wireTwo := getWire(scanner)
	_, matches := plotPath(wireTwo, wireOnePath)
	closest, shortest := 0, 0
	for locString, steps := range matches {
		pos := strings.Split(locString, ",")
		locX, _ := strconv.Atoi(pos[0])
		locY, _ := strconv.Atoi(pos[1])

		dist := abs(locX) + abs(locY)

		if dist < closest || closest == 0 {
			closest = dist
		}
		if steps < shortest || shortest == 0 {
			shortest = steps
		}
	}
	fmt.Printf("%d\n%d\n", closest, shortest)
}

func plotPath(wire []string, compare map[string]int) (map[string]int, map[string]int) {
	dirs := make(map[byte][]int)
	dirs[byte('L')] = []int{0, -1}
	dirs[byte('R')] = []int{0, 1}
	dirs[byte('D')] = []int{1, -1}
	dirs[byte('U')] = []int{1, 1}

	path := make(map[string]int)
	steps := 0
	matches := make(map[string]int)
	loc := [2]int{0, 0} //XY
	for _, key := range wire {
		//fmt.Printf("parsed %s\n", key)
		isVert := dirs[key[0]][0]
		sign := dirs[key[0]][1]
		walk, err := strconv.Atoi(key[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 1; i <= walk; i++ {
			loc[isVert] += sign
			steps++
			locString := fmt.Sprintf("%d,%d", loc[0], loc[1])
			path[locString] = steps

			if otherSteps, ok := compare[locString]; compare != nil && ok {
				matches[locString] = steps + otherSteps
			}
		}
		//fmt.Printf("current loc %v\n", loc)
		//os.Exit(2)
	}
	//fmt.Println(matches)
	return path, matches
}

func getWire(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
