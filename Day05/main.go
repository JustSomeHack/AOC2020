package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

const rows int = 128
const seats int = 8

func main() {
	readFile("./input.txt")
}

func processRow(row string) int {
	fRow := 0
	bRow := rows - 1

	for _, c := range row {
		diff := (bRow - fRow + 1) / 2
		char := string(c)
		switch char {
		case "F":
			bRow -= diff
			break
		case "B":
			fRow += diff
			break
		}
	}

	if fRow != bRow {
		log.Fatalf("Error getting row")
	}

	return fRow
}

func processSeat(row string) int {
	fSeat := 0
	bSeat := seats - 1

	for _, c := range row {
		diff := (bSeat - fSeat + 1) / 2
		char := string(c)
		switch char {
		case "L":
			bSeat -= diff
			break
		case "R":
			fSeat += diff
			break
		}
	}

	if fSeat != bSeat {
		log.Fatalf("Error getting row")
	}

	return fSeat
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

	ids := make([]int, 0)

	for _, line := range text {
		row := processRow(line)
		seat := processSeat(line)
		seatID := row*8 + seat
		ids = append(ids, seatID)
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] > ids[j]
	})

	currentID := 0
	for _, id := range ids {
		if currentID == 0 {
			currentID = id
			continue
		}
		if currentID-1 != id {
			log.Printf("Missing %d", currentID-1)
		}
		currentID = id
	}
}
