package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.8
	userKg := 100.0
	fmt.Scan(&userHeight)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Println(IMT)
}
