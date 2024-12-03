package main

import "fmt"

// interacts with factory and factory intrates with others 
// why ? application to be loosely coupled 
// ammending does not affect to others
// how ?? to
type Car interface {
	getCar() string
}
type SedanType struct{
	Name string
}
func getNewSedan() *SedanType{
	return &SedanType{}
}

func (s *SedanType) getCar string{
	return "Honda City"
}
type HactBackType struct{
	Name string
}
func getNewHactBack() *HactBackType{
	return &HactBackType{}
}

func (s *HactBackType) getCar string{
	return "polo"
}


func main(){
	getCarFactory(1)
	getCarFactory(2)
}

func getCarFactory(cartype int) {
	var car Car
	if cartype == 1 {
		car = getNewHactBack()
	}else{
		car = getNewSedan()
	}

	carInfo := car.getCar()
	fmt.Println(carInfo)
}
