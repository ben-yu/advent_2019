package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Intersection(p1 Point, p2 Point, q1 Point, q2 Point) Point {
	var intersection Point
	pXDir := p2.X - p1.X
	pYDir := p2.Y - p1.Y
	qXDir := q2.X - q1.X
	qYDir := q2.Y - q1.Y

	// p is vertical and q is horizontal
	if p2.X-p1.X == 0 && q2.Y-q1.Y == 0 {
		if pYDir >= 0 && qXDir >= 0 {
			if (p1.X >= q1.X && p1.X <= q2.X) && (q1.Y >= p1.Y && q1.Y <= p2.Y) {
				intersection = Point{p1.X, q1.Y}
			}
		} else if pYDir > 0 && qXDir < 0 {
			if (p1.X <= q1.X && p1.X >= q2.X) && (q1.Y >= p1.Y && q1.Y <= p2.Y) {
				intersection = Point{p1.X, q1.Y}
			}
		} else if pYDir < 0 && qXDir > 0 {
			if (p1.X >= q1.X && p1.X <= q2.X) && (q1.Y <= p1.Y && q1.Y >= p2.Y) {
				intersection = Point{p1.X, q1.Y}
			}
		} else {
			if (p1.X <= q1.X && p1.X >= q2.X) && (q1.Y <= p1.Y && q1.Y >= p2.Y) {
				intersection = Point{p1.X, q1.Y}
			}
		}
		// p is horizontal and q is vertical
	} else if p2.Y-p1.Y == 0 && q2.X-q1.X == 0 {
		if pXDir >= 0 && qYDir >= 0 {
			if (p1.Y >= q1.Y && p1.Y <= q2.Y) && (q1.X >= p1.X && q1.X <= p2.X) {
				intersection = Point{q1.X, p1.Y}
			}
		} else if pXDir > 0 && qYDir < 0 {
			if (p1.Y <= q1.Y && p1.Y >= q2.Y) && (q1.X >= p1.X && q1.X <= p2.X) {
				intersection = Point{q1.X, p1.Y}
			}
		} else if pXDir < 0 && qYDir > 0 {
			if (p1.Y >= q1.Y && p1.Y <= q2.Y) && (q1.X <= p1.X && q1.X >= p2.X) {
				intersection = Point{q1.X, p1.Y}
			}
		} else {
			if (p1.Y <= q1.Y && p1.Y >= q2.Y) && (q1.X <= p1.X && q1.X >= p2.X) {
				intersection = Point{q1.X, p1.Y}
			}
		}
	}
	return intersection
}

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scan input into slice of points
	scanner := bufio.NewScanner(file)
	wires := make([][]Point, 2)

	for x, wire := range wires {
		scanner.Scan()
		s := strings.Split(scanner.Text(), ",")
		origin := Point{0, 0}
		wire = append(wire, origin)

		for _, i := range s {
			dir := string(i[0])
			dist, err := strconv.Atoi(i[1:])
			if err != nil {
				log.Fatal(err)
			}
			var newPoint Point
			if dir == "U" {
				newPoint = Point{origin.X, origin.Y + dist}
			} else if dir == "D" {
				newPoint = Point{origin.X, origin.Y - dist}
			} else if dir == "L" {
				newPoint = Point{origin.X - dist, origin.Y}
			} else if dir == "R" {
				newPoint = Point{origin.X + dist, origin.Y}
			}
			wire = append(wire, newPoint)
			origin = newPoint
			//log.Printf("%v", wire)
		}
		wires[x] = wire
	}

	log.Printf("%v", wires)

	// Compare all line segments for intersection points
	minDist := math.MaxFloat64
	for i := 0; i < len(wires[0])-1; i += 1 {
		for j := 0; j < len(wires[1])-1; j += 1 {
			intersection := Intersection(wires[0][i], wires[0][i+1], wires[1][j], wires[1][j+1])
			if intersection.X == 0 && intersection.Y == 0 {
			} else {
				log.Printf("Intersection: %v", intersection)
				dist := math.Abs(float64(intersection.X)) + math.Abs(float64(intersection.Y))
				if dist < minDist {
					minDist = dist
				}
			}
		}
	}

	log.Printf("%v", minDist)
}
