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
	//file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel int64 = 0
	for scanner.Scan() {
		inputMass, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var moduleFuel int64 = 0
		for inputMass > 0 {
			fuelRequired := int64(math.Floor(float64(inputMass)/3.0)) - 2
			if fuelRequired > 0 {
				//log.Printf("Added %v", fuelRequired)
				moduleFuel += fuelRequired
			}
			inputMass = fuelRequired
		}
		totalFuel += moduleFuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total fuel required: %v", totalFuel)
}
