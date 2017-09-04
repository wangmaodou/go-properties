package main

import (
	"./property"
	"fmt"
)

func main() {

	p:=property.FromFile("./property/test.properties")
	fmt.Println("name=",p.Get("name"))
	fmt.Println("country=",p.Get("country"))
}
