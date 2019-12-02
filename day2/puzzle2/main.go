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
	orignalInstructions := getInput("input")

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			instructions := make([]int, len(orignalInstructions))
			copy(instructions, orignalInstructions)
			instructions[1] = noun
			instructions[2] = verb
			res := run(instructions)
			fmt.Printf("running with noun{%2d} and verb{%2d} | res{%d} \n", noun, verb, res)
			if res == 19690720 {
				fmt.Printf("result : %d", 100*noun+verb)
				return
			}
		}
	}

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
