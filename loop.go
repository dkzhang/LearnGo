package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func main() {
	apply(func(i int, i2 int) int {
		return i * i2
	}, 3,4)

}

func convertToBin(n int) string{
	result := ""
	if n==0 {
		return "0"
	}

	for ; n>0; n /=2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func div(a,b int) (q,r int){
	return a/b, a%b
}

func apply(op func(int,int) int, a, b int)int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" +
		"(%d, %d) = %d \n", opName, a, b, op(a,b))
	return op(a,b)
}

func myPlus(a,b int) int{
	return a+b
}
