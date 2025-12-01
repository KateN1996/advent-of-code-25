package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	decodePassword("input.txt")
}

func decodePassword(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numOfZeros := 0
	curVal := 50

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		num, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			curVal -= num
		} else {
			curVal += num
		}

		curVal = ((curVal % 100) + 100) % 100

		if curVal == 0 {
			numOfZeros += 1
		}
	}
	fmt.Printf("\n\nthis is the password: %d\n", numOfZeros)
	return numOfZeros

}
