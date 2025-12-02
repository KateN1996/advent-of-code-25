package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	decodePassword2("input.txt")
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

func decodePassword2(fileName string) int {
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

		oldVal := curVal
		if direction == 'L' {
			curVal -= num
		} else {
			curVal += num
		}

		if direction == 'L' {
			numOfZeros += floorDiv(oldVal-1, 100) - floorDiv(curVal-1, 100)
		} else {
			numOfZeros += floorDiv(curVal, 100) - floorDiv(oldVal, 100)
		}
		fmt.Printf("zeroes are now %d\n\n", numOfZeros)

		curVal = ((curVal % 100) + 100) % 100
	}

	fmt.Printf("\n\nthis is the password: %d\n", numOfZeros)
	return numOfZeros
}

func floorDiv(a, b int) int {
	q := a / b
	r := a % b
	if (r != 0) && ((r < 0) != (b < 0)) {
		q--
	}
	return q
}
