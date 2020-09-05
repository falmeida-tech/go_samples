package main

import (
	"fmt"
	"go_samples/go-pkg-sample/stringutils"
	"go_samples/go-pkg-sample/calc"
)

func main(){
	s:="fabio"
	fmt.Println(stringutils.Upper(s))
	t:= "FABIO"
	fmt.Println(stringutils.Lower(t))
	result := calc.Add(2,3)
	println(result)
	println(calc.Minus(7,2))
	printMemoryAddr()
}

func printMemoryAddr(){
	var a = 10
	fmt.Printf("Address of a variable: %d\n", &a  )
}