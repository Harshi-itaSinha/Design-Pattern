package main

import (
	"fmt"
	"sync"
)

// when we need only one obj like connection
// config etc
//

var mutex =&sync.Mutex{}

type config struct {

	// various things can be here in this struct 
	// i think go can compete a lot 

	// config variable
}

var counter =1;
var singleConfiguration *config

func getConfigInstance() * config{

	if singleConfiguration == nil{
		mutex.Lock() // does mutex work?? why it the go routine executing
		// till down??

		defer mutex.Unlock()
		if singleConfiguration == nil {
			fmt.Println("creating single",counter)
			singleConfiguration = &config{}
			counter++
		}else{
			// this called many a times
			fmt.Println("already single -1 instance")
		}
	}else{
		fmt.Println("already single -2instance")
	}
	return singleConfiguration
}

func main(){
	for i:=0 ; i<100;i++{
		go getConfigInstance()
	}
	fmt.Scanln()
}