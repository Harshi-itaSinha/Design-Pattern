package main

import "fmt"

// Car interface
type Car interface {
	getCar() string
}

// SedanType struct implementing Car interface
type SedanType struct {
	Name string
}

func getNewSedan() Car {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Honda City"
}

// HactBackType struct implementing Car interface
type HactBackType struct {
	Name string
}

func getNewHactBack() Car {
	return &HactBackType{}
}

func (h *HactBackType) getCar() string {
	return "Polo"
}

// Main function
func main() {
	getCarFactory(1) // Should return "Polo"
	getCarFactory(2) // Should return "Honda City"
}

// Factory function
func getCarFactory(cartype int) {
	var car Car
	if cartype == 1 {
		car = getNewHactBack()
	} else {
		car = getNewSedan()
	}

	carInfo := car.getCar()
	fmt.Println(carInfo)
}
