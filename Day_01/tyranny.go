package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type mass struct {
	value int
	next  *mass
}

type Stack struct {
	size int
	top  *mass
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value int) {
	stack.top = &mass{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value int) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	return 0
}

func fuelRequired(mass int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		fuel = 0
	}
	return fuel
}

//Takes a newline seperated list of masses and returns total output
func totalFuelRequired(masses []int) int {
	var totalFuel int
	for _, mass := range masses {
		totalFuel += fuelRequired(mass)
	}
	return totalFuel
}

func totalFuelRequiredPartII(masses []int) int {
	var totalFuel int
	stack := new(Stack)
	for _, mass := range masses {
		stack.Push(mass)
	}

	for {
		//fmt.Printf("Stack size: %d\n", stack.Len())
		//fmt.Printf("Total Fuel: %d\n", totalFuel)
		if stack.Len() > 0 {
			mass := stack.Pop()
			//	fmt.Printf("Mass: %d\n", mass)
			fuel := fuelRequired(mass)
			//	fmt.Printf("Fuel: %d\n", fuel)
			if fuel != 0 {
				stack.Push(fuel)
			}
			totalFuel += fuel
		} else {
			//	fmt.Println("Loop should end")
			break
		}
	}
	return totalFuel
}

func readMassesFromFile(fn string) ([]int, error) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Error opening file %q", err)
	}
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
	masses, _ := readMassesFromFile("/Users/seanslattery/dev/AoC_2019/Day_01/masses.txt")
	total := totalFuelRequired(masses)
	fmt.Printf("The total amount of fuel required for part I is: %d\n", total)
	total2 := totalFuelRequiredPartII(masses)
	fmt.Printf("The total amount of fuel required for part II is: %d\n", total2)
}
