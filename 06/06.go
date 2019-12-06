//You download a map of the local orbits (your puzzle input).
//Except for the universal Center of Mass (COM), every object in space is in orbit around exactly one other object.
//To verify maps, the Universal Orbit Map facility uses orbit count checksums -
//the total number of direct orbits (like the one shown above) and indirect orbits.
//If A orbits B, B orbits C, and C orbits D, then A indirectly orbits D.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//1. What is the total number of direct and indirect orbits in your map data?
	//2. What is the minimum number of orbital transfers required to move from
	//   the object YOU are orbiting to the object SAN is orbiting?
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	orbits, revOrbits := make(map[string][]string), make(map[string]string)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ")")
		orbits[text[0]] = append(orbits[text[0]], text[1])
		revOrbits[text[1]] = text[0]
	}

	fmt.Println(getOrbitSum(orbits))

	fmt.Println(findHopsBetween(orbits, revOrbits, revOrbits["YOU"], revOrbits["SAN"]))
}

func getAllDeps(revOrbits map[string]string, object string) []string {
	deps := []string{}
	for currObj := object; ; {
		dep := revOrbits[currObj]
		deps = append(deps, dep)
		if dep == "COM" {
			break
		} else {
			currObj = dep
		}
	}
	return deps
}

func findHopsBetween(orbits map[string][]string, revOrbits map[string]string, objectOne string, objectTwo string) int {
	//get all dependencies of both objects
	objOneDeps := getAllDeps(revOrbits, objectOne)
	objTwoDeps := getAllDeps(revOrbits, objectTwo)

	for i, depOne := range objOneDeps {
		for j, depTwo := range objTwoDeps {
			if depOne == depTwo {
				return i + j + 2
			}
		}
	}
	return -1
}

func getOrbitSum(orbits map[string][]string) int {
	objects, nextObjects := []string{}, []string{"COM"}
	orbitSum := 0
	for depth := 0; nextObjects != nil; {
		objects = nextObjects
		nextObjects = nil
		depth++

		for _, obj := range objects {
			for _, next := range orbits[obj] {
				nextObjects = append(nextObjects, next)
				orbitSum += depth
			}
		}
	}
	return orbitSum
}
