package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var valid int = 0

func main() {
	readFile("./input.txt")

	fmt.Printf("\n\nFound %d valid passwords\n\n", valid)
}

func processLine(line string) {
	line = strings.ReplaceAll(line, ":", "")
	parts := strings.Split(line, " ")

	nums := strings.Split(parts[0], "-")
	min, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}

	key := []rune(parts[1])[0]
	password := parts[2]

	minPos := []rune("test")[0]
	maxPos := []rune("test")[0]

	for i, c := range password {
		if (i + 1) == min {
			minPos = c
			continue
		}
		if (i + 1) == max {
			maxPos = c
			continue
		}
	}

	if (minPos == key && maxPos != key) || (minPos != key && maxPos == key) {
		fmt.Printf("%s is valid\n", password)
		valid++
		return
	}
	fmt.Printf("%s is not valid\n", password)
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
		processLine(line)
	}
}
