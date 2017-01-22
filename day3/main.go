package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"../common"
)

func checkTriangle(numbers ...string) bool {
	var els []int
	for _, x := range numbers {
		v, _ := strconv.Atoi(x)
		els = append(els, v)
	}

	return els[0]+els[1] > els[2] && els[0]+els[2] > els[1] && els[1]+els[2] > els[0]
}

func createRegexp() *regexp.Regexp {
	reg, err := regexp.Compile("([0-9]+)")
	if err != nil {
		panic(err)
	}

	return reg
}

func getLineElements(line string, reg *regexp.Regexp) []string {
	return reg.FindAllString(line, -1)
}

func part1() {
	file := common.OpenFile("numbers.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := createRegexp()

	count := 0

	for scanner.Scan() {
		numbers := getLineElements(scanner.Text(), reg)
		if checkTriangle(numbers...) {
			count++
		}
	}

	fmt.Println("count part 1", count)
}

func part2() {
	file := common.OpenFile("numbers.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := createRegexp()

	count := 0
	var numbers []string

	for scanner.Scan() {
		numbers = append(numbers, getLineElements(scanner.Text(), reg)...)
	}

	for i := 0; i < len(numbers)/3; i += 3 {
		for j := 0; j < 3; j++ {
			if checkTriangle(numbers[i*3+j], numbers[i*3+(j+3)], numbers[i*3+(j+6)]) {
				count++
			}
		}
	}

	fmt.Println("Count part 2", count)
}

func main() {
	part1()
	part2()
}
