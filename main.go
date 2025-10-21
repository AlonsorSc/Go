package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println("__ Калькулятор индекса массы тела __")
    userKG, userHeight := getUserInput()
    IMT := calculateIMT(userKG, userHeight)
    outputResult(IMT)
}

func outputResult(imt float64) {
    result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
    fmt.Print(result)
}
func calculateIMT(userKG, userHeight float64) float64 {
    const IMTPower = 2
    IMT := userKG / math.Pow(userHeight/100, IMTPower)
    return IMT
}

func getUserInput() (float64, float64) {
    var userHeight, userKG float64
    fmt.Print("Введите вес в киллограммах: ")
    fmt.Scan(&userKG)
    fmt.Print("Введите рост в метрах: ")
    fmt.Scan(&userHeight)
    return userKG, userHeight
}
