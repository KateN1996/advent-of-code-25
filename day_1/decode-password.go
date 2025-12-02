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
		fmt.Printf("num is %d, dir is %d\n", num, direction)

		if direction == 'L' {
			numOfZeros += (curVal+1)/100 - (curVal-num+1)/100
			curVal -= num
		} else {
			numOfZeros += (curVal+num)/100 - curVal/100
			curVal += num
		}
		fmt.Printf("curval is %d\n", curVal)

		//numOfZeros += int(math.Abs(float64(curVal)) / 100)
		curVal = ((curVal % 100) + 100) % 100

		if curVal == 0 {
			numOfZeros += 1
			fmt.Printf("curval is 0; zeroes are now %d\n\n", numOfZeros)
		}
	}
	fmt.Printf("\n\nthis is the password: %d\n", numOfZeros)
	return numOfZeros
}
