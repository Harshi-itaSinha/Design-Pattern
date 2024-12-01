package main

/*
Adapter Design patter is used , when we want to connect already implemented  system 
and can't be changed and to connect he two we form adapter design pattern

to make 2 systems compatible 


*/
import "fmt"

// Target
type mobile interface{
	chargeAppleMobile() // this function is there for both android adapter and apple 
}

// Concrete prototype implementation
type apple struct {}

// Concerete Implementation 
func(a *apple) chargeAppleMobile(){
	fmt.Println("Charging Apple Mobile")
}

type android struct{}

func(a *android) chargeAndroidMobile(){
	fmt.Println("Charging Android Mobile")
}

// Adapter 

type androidAdapter struct{
	android *android
}
func(ad *androidAdapter) chargeAppleMobile(){
	ad.android.chargeAndroidMobile()
}
// client 
type client struct{}
func (c* client) chargeMobile(mob mobile){
	mob.chargeAppleMobile()
}

func main() {
	// initial requirement
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	// extended requirement
	android:= &android{}
	// wrapper to make it compatible
	androidAdapter :=&androidAdapter{
		android:android,
	}
	// same function is passed and only the parameter is changing
	// if i had added 
	client.chargeMobile(androidAdapter)
}