package main

import "fmt"

// try the one in which interface is implemented 

type Car struct {
	Brand  string
	Model  string
	Color  string
	Engine string
}

type ICarBuilder interface {
	color(c string)
	brand(b string)
	model(m string)
	Engine(e string)
}

type CarBuilder struct {
	car *Car
}


func NewCarBuilder () *CarBuilder{
	return &CarBuilder{car: &Car{}} // empty struct
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.car.Brand = brand
	return cb
}

// SetModel sets the model of the car
func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.car.Model = model
	return cb
}

// SetColor sets the color of the car
func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.car.Color = color
	return cb
}

// SetEngine sets the engine type of the car
func (cb *CarBuilder) SetEngine(engine string) *CarBuilder {
	cb.car.Engine = engine
	return cb
}
// final function that will give me 
func (cb *CarBuilder) Build() *Car {
	return cb.car // finally which will give always the final object 
}
// one function to return builder

func main() {
	// Using the builder to create a car object

	// since each time it is returning the same object again and again this can be done 
	car := NewCarBuilder().
		SetBrand("Tesla").
		SetModel("Model S").
		SetColor("Red").
		SetEngine("Electric").
		Build()

	fmt.Printf("Car Details: %+v\n", car)
}
