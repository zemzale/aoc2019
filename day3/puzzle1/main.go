package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cord struct {
	X, Y int
}

type Direction int

const (
	Up    Direction = 1
	Right           = 1
	Down            = -1
	Left            = -1
)

type Move struct {
	Direction Direction
	Amount    int
	Start     Cord
	End       Cord
}

func NewMove(code string, start Cord) Move {
	var dir Direction
	amount, err := strconv.Atoi(code[1:])
	if err != nil {
		log.Fatal(err)
	}
	var (
		x = start.X
		y = start.Y
	)
	switch string([]rune(code)[0]) {
	case "U":
		dir = Up
		y = +(amount * int(dir))
	case "R":
		dir = Right
		x = +(amount * int(dir))
	case "D":
		dir = Down
		y = +(amount * int(dir))
	case "L":
		dir = Left
		x = +(amount * int(dir))
	}

	end := Cord{
		X: x,
		Y: y,
	}

	return Move{
		Direction: dir,
		Amount:    amount,
		Start:     start,
		End:       end,
	}
}

func main() {

	wire1, _ := getInput("input")

	input1 := make([]Move, len(wire1))

	var cord = Cord{}
	for k, v := range wire1 {
		input1[k] = NewMove(v, cord)
		cord = input1[k].End
	}
	PrettyPrint(input1)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func getInput(path string) ([]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strings.Split(lines[0], ","), strings.Split(lines[1], ",")
}
