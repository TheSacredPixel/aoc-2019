//Fuel required to launch a given module is based on its mass.
//Specifically, to find the fuel required for a module, take its mass,
//divide by three, round down, and subtract 2.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sum     = 0
	modules []int
)

func main() {
	//What is the sum of the fuel requirements for
	//all of the modules on your spacecraft?
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		modules = append(modules, i)
	}
	for _, n := range modules {
		sum += calcFuel(n)
	}
	fmt.Println(sum)

	//What is the sum of the fuel requirements for all of the modules on your spacecraft
	//when also taking into account the mass of the added fuel?
	sum = 0
	for _, n := range modules {
		fuel := calcFuel(n)
		for fuel > 0 {
			sum += fuel
			fuel = calcFuel(fuel)
		}
	}
	fmt.Println(sum)

}

func calcFuel(mass int) int {
	return mass/3 - 2
}
