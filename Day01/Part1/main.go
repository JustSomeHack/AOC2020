package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numbers = []int{}

func main() {
	readFile("./input.txt")

	found := false
	for i1, n1 := range numbers {
		for i2, n2 := range numbers {
			if i1 == i2 {
				continue
			}
			if n1+n2 == 2020 {
				fmt.Printf("Found %d + %d == %d\n", n1, n2, n1*n2)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, line := range text {
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Can't convert %v", line)
			continue
		}
		numbers = append(numbers, i)
	}
}
