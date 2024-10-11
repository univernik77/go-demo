package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("_Калькулятор индекса массы тела_")
	fmt.Println("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите ваш вес в килограммах: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", IMT)
	fmt.Print(result)
}
