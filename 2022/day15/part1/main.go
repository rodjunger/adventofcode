package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type sensor struct {
	x, y,
	beaconX, beaconY,
	distanceToClosestBeacon int
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

	y := 2000000

	//This is the lazy way, math would be way faster (keeping a list of min/maxIndexes that were painted)
	//Logic is actually pretty simple too, I should probably do it
	occupyedSpots := make(map[int]struct{})

	for _, sensor := range sensors {
		distanceFromY := Abs(sensor.y - y)

		//fmt.Printf("Sensor x:%d, y: %d, distanceToClosestBeacon: %d, distanceFromY: %d\n",
		//	sensor.x, sensor.y, sensor.distanceToClosestBeacon, distanceFromY)

		if distanceFromY > sensor.distanceToClosestBeacon {
			continue
		}

		distanceDelta := sensor.distanceToClosestBeacon - distanceFromY

		//fmt.Println("sensor is close enough, marking", (distanceDelta*2)+1, "indices")

		minIndex, maxIndex := sensor.x-distanceDelta, sensor.x+distanceDelta

		//fmt.Println("marking indexes", minIndex, "to", maxIndex)
		for i := minIndex; i <= maxIndex; i++ {
			occupyedSpots[i] = struct{}{}
		}
	}

	for _, sensor := range sensors {
		if sensor.beaconY == y {
			delete(occupyedSpots, sensor.beaconX)
		}
	}

	fmt.Println(len(occupyedSpots))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
