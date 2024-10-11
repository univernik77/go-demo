package main

import (
	"fmt"
	"math"
)

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиментрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func main() {
	const IMTPower = 2
	userHeight, userKg := getUserInput()
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", IMT)
	fmt.Print(result)
}
