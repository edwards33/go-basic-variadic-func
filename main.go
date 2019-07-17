package main

import (
	"fmt"
)

func main(){
	num := average(21, 23, 24, 34, 45, 54)

	fmt.Println(num)
}

func average(args ...float64) float64 {
	fmt.Printf("%T \n", args)
	var sum float64

	for _, v := range args {
		sum += v
	}

	return sum / float64(len(args))
} 