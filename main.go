package main

import (
	"fmt"
	"github.com/maheshkumaarbalaji/goconcepts/lib/playground"
)


func main() {
	fmt.Println("###############  Go Concepts Start  ###############")
	
	
	fmt.Println("###############  Go Concepts End  ###############")
}

func CheckCustomError() {
	err := playground.SampleFunction()
	err, ok := err.(*playground.CustomError)
	if !ok {
		fmt.Println("Expected a custom type error but got something else")
	} else {
		fmt.Println("Got the expected error")
		fmt.Println(err.Error())
	}
}