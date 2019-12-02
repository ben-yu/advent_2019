package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadValue(loc int, arr []int) int {
	if loc < 0 || loc >= len(arr) {
		log.Fatal("Out of bounds index: %v", loc)
	}

	return arr[loc]
}

func saveValue(loc int, arr []int, value int) []int {
	if loc < 0 || loc >= len(arr) {
		log.Fatal("Out of bounds index: %v", loc)
	}

	arr[loc] = value
	return arr
}

func runProgram(noun int, verb int) int {
	file, err := os.Open("./day_2_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scan input into slice
	scanner := bufio.NewScanner(file)
	programArr := make([]int, 0)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		for _, i := range s {
			j, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			programArr = append(programArr, j)
		}
	}

	// Do replacement as asked for
	programArr[1] = noun
	programArr[2] = verb

	// Iterate through opcodes
	for i := 0; i < len(programArr); i += 4 {
		opCode := programArr[i]
		if opCode == 1 {
			a := loadValue(programArr[i+1], programArr)
			b := loadValue(programArr[i+2], programArr)
			programArr = saveValue(programArr[i+3], programArr, a+b)
		} else if opCode == 2 {
			a := loadValue(programArr[i+1], programArr)
			b := loadValue(programArr[i+2], programArr)
			programArr = saveValue(programArr[i+3], programArr, a*b)
		} else if opCode == 99 {
			return programArr[0]
		} else {
			log.Fatal("Unsupported opCode: %v", opCode)
		}
		log.Printf("%v", programArr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return -1
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := runProgram(noun, verb)
			if result == 19690720 {
				log.Printf("Noun: %v, Verb: %v", noun, verb)
				os.Exit(0)
			}
		}
	}
}
