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
		sum += number/3 - 2
	}
	fmt.Printf("Sum is : %d \n", sum)
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
