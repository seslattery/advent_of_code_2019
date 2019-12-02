package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fuelRequired(mass int) int {
	return (mass / 3) - 2
}

//Takes a newline seperated list of masses and returns total output
func totalFuelRequired(masses []int) int {
	var totalFuel int
	for _, mass := range masses {
		totalFuel += fuelRequired(mass)
	}
	return totalFuel
}

func readMassesFromFile(fn string) ([]int, error) {
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)

	var line string
	var masses []int
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
		sanitizedLine := strings.TrimSuffix(line, "\n")
		i, err := strconv.Atoi(sanitizedLine)
		if err != nil {
			fmt.Printf("Error converting string to int: %q\n", err)
			break
		}
		masses = append(masses, i)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return masses, err
}

func main() {
	masses, _ := readMassesFromFile("/Users/seanslattery/dev/AoC_2019/Day_01/test.txt")
	total := totalFuelRequired(masses)
	fmt.Printf("The total amount of fuel required is: %d\n", total)
}
