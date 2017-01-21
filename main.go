package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkTriangle(els []int) bool {
	return els[0]+els[1] > els[2] && els[0]+els[2] > els[1] && els[1]+els[2] > els[0]
}

func part1() {
	file, err := os.Open("numbers.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	reg, e := regexp.Compile("([0-9]+)")
	if e != nil {
		fmt.Println(e)
		return
	}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		var els []int
		for _, x := range reg.FindAllString(line, -1) {
			v, _ := strconv.Atoi(x)
			els = append(els, v)
		}

		if checkTriangle(els) {
			count++
		}
	}

	fmt.Println("count part 1", count)
}

func main() {
	part1()
}
