package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	mostCalories := day_one()
	fmt.Println("MOST_CALORIES: ", mostCalories)

}

type Elf struct {
	calories int
	numItems int
}

func day_one() int {
	data, err := os.ReadFile("./inputs/day_one_part_one.txt")
	if err != nil {
		errMsg := fmt.Errorf("Error reading input file: %w", err)
		log.Panic(errMsg)
	}
	split := strings.Split(string(data), "\n")
	elfs := []Elf{}
	//
	elfIdx := 0
	for i := range split {
		if len(elfs) == 0 {
			elfs = append(elfs, Elf{calories: 0, numItems: 0})
		}
		if split[i] == "" {
			// NEW ELF
			elfIdx += 1
			elfs = append(elfs, Elf{calories: 0, numItems: 0})
		}
		if split[i] != "" {
			numCals, err := strconv.Atoi(split[i])
			if err != nil {
				formatted := fmt.Errorf("Error converting to int: %w", err)
				fmt.Println(formatted)
			}
			elfs[elfIdx].calories += numCals
			elfs[elfIdx].numItems += 1

		}

	}

	sumCalories := [3]int{0, 0, 0}
	for i := 0; i < 3; i++ {
		for j := range elfs {
			if elfs[j].calories > sumCalories[i] {

				if !contains(sumCalories, elfs[j].calories) {
					sumCalories[i] = elfs[j].calories
				}
			}

		}
	}
	fmt.Println(sumCalories)

	// return sumCalories
	sumCals := 0
	for i := 0; i < len(sumCalories); i++ {
		sumCals += sumCalories[i]
	}

	return sumCals

}
func contains(s [3]int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
