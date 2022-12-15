package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type sensor struct {
	vector
	beaconX, beaconY,
	distanceToClosestBeacon int
}

type vector struct {
	x, y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sensors := []sensor{}

	for scanner.Scan() {
		var (
			curSensor sensor
		)
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &curSensor.x, &curSensor.y, &curSensor.beaconX, &curSensor.beaconY)
		curSensor.distanceToClosestBeacon = manhattanDistance(curSensor.x, curSensor.y, curSensor.beaconX, curSensor.beaconY)
		sensors = append(sensors, curSensor)
	}

	row := []vector{}

	rows := 4000000

	for y := 0; y < rows; y++ {
		for _, sensor := range sensors {
			distanceFromY := Abs(sensor.y - y)

			if distanceFromY > sensor.distanceToClosestBeacon {
				continue
			}

			distanceDelta := sensor.distanceToClosestBeacon - distanceFromY

			minIndex, maxIndex := sensor.x-distanceDelta, sensor.x+distanceDelta

			if minIndex < 0 {
				minIndex = 0
			}

			if maxIndex > rows {
				maxIndex = rows
			}

			row = append(row, vector{minIndex, maxIndex})
		}

		sort.Slice(row, func(i, j int) bool {
			return row[i].x < row[j].x
		})

		maxIter := len(row)

		for j := 1; j < maxIter; j++ {
			//skip invalid rows
			if row[j].y < row[j-1].y {
				row[j] = row[j-1]
			}
			//Find the gap
			if row[j].x-row[j-1].y == 2 {
				fmt.Println("found x:", row[j].x-1, "y:", y)
				fmt.Println((uint64(row[j].x-1) * 4000000) + uint64(y))
			}
		}

		row = row[:0]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
