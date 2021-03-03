package main

import "fmt"

var startingOrderMap = make(map[float64]int)

type Driver struct {
	DriverName           string
	CarNumber            int
	DriverPreviousFinish float64
	OwnerLastFinish      float64
	OwnerPoints          float64
	PreviousFastLap      float64
}

func qualifyingOrder(driver Driver) float64 {
	positionRank := 0.0
	positionRank += driver.DriverPreviousFinish * .25
	positionRank += driver.OwnerLastFinish * .25
	positionRank += driver.OwnerPoints * .35
	positionRank += driver.PreviousFastLap * .15
	return positionRank
}

func mapStartingPositionToCarNumber(carNumber int, startingOrderRank float64) {
	startingOrderMap[startingOrderRank] = carNumber
}

func main() {
	var bradK Driver
	bradK.CarNumber = 2
	bradK.DriverName = "Brad Keselowski"
	bradK.DriverPreviousFinish = 17
	bradK.OwnerLastFinish = 17
	bradK.OwnerPoints = 10
	bradK.PreviousFastLap = 8

	startingOrderRank := qualifyingOrder(bradK)

	mapStartingPositionToCarNumber(bradK.CarNumber, startingOrderRank)

	var michaelM Driver
	michaelM.CarNumber = 34
	michaelM.DriverName = "Michael M"
	michaelM.DriverPreviousFinish = 6
	michaelM.OwnerLastFinish = 6
	michaelM.OwnerPoints = 4
	michaelM.PreviousFastLap = 2

	startingOrderRank = qualifyingOrder(michaelM)

	mapStartingPositionToCarNumber(michaelM.CarNumber, startingOrderRank)

	fmt.Println(startingOrderMap)
}
