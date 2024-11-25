package main

import (
	"fmt"
	"github.com/maheshkumaarbalaji/goconcepts/lib/playground"
	"github.com/maheshkumaarbalaji/goconcepts/lib/ds"
)


func main() {
	fmt.Println("###############  Go Concepts Start  ###############")
	TestTrieTree()
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

func TestTrieTree() {
	root := ds.CreateTree()
	ds.AddRouteToTree(root, "/user/mahesh")
	ds.AddRouteToTree(root, "/user/nehru")
	ds.AddRouteToTree(root, "/routes/addition")
	ds.AddRouteToTree(root, "/routes/division")
	routeParts := root.Print()
	for _, routePart := range routeParts {
		fmt.Println(routePart)
	}
	fmt.Println("============================================")

}