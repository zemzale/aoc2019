package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers = getInput("input")
	var sum int

	for _, number := range numbers {
		sum += calculateFuel(number)
	}
	fmt.Printf("Sum is : %d \n", sum)
}

func calculateFuel(weight int) int {
	var sum int
	for {
		weight = weight/3 - 2
		if weight <= 0 {
			return sum
		}
		sum += weight
	}
}

func getInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
