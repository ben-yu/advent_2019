package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day_1_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel int64 = 0
	for scanner.Scan() {
		inputMass, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fuelRequired := int64(math.Floor(float64(inputMass)/3.0)) - 2
		totalFuel += fuelRequired
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total fuel required: %v", totalFuel)
}
