package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println("Hello world")
	euler()
	triangle()
	enums()
}

func euler(){
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
}

func triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func enums(){
	const(
		cpp=100
		java=101
		python=iota
		golang
	)
	fmt.Println(cpp,java,python,golang)
}
