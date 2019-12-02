package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := getInput("input")

	instructions[1] = 12
	instructions[2] = 2
	fmt.Printf("result : %d", run(instructions))

}

func run(instructions []int) int {
	for i := 0; i <= len(instructions); i += 4 {
		opcode := instructions[i]
		arg1 := instructions[i+1]
		arg2 := instructions[i+2]
		res := instructions[i+3]
		switch opcode {
		case 1:
			instructions[res] = add(instructions[arg1], instructions[arg2])
		case 2:
			instructions[res] = multply(instructions[arg1], instructions[arg2])
		default:
			return instructions[0]
		}
	}
	return instructions[0]
}

func add(x, y int) int {
	return x + y
}

func multply(x, y int) int {
	return x * y
}

func getInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileBuffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var (
		values       = string(fileBuffer)
		instructions = make([]int, 0)
	)

	for _, number := range strings.Split(values, ",") {
		i, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, i)
	}

	return instructions
}
