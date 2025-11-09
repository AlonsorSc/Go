package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover", r)
		}
	}()
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userKG, userHeight := getUserInput()         // присваиваем в переменные вывод из функции getUserInput
		IMT, err := calculateIMT(userKG, userHeight) // после расчёта функции calculateIMT с аргументами userKG и userHeight, которые мы в неё передали, мы получили результат и присвоили его переменной IMT
		if err != nil {
			// fmt.Println("Не заданы параметры для расчёта")
			// continue
			panic("Не заданы параметры для расчёта")
		}
		outputResult(IMT) // Выводим результат IMT благодаря функции outputResult
		// Цикл. Выполнеяем проверку опроса, хочет пользовать продолжить или нет через функцию checkRepeatCalculation. Если данная функция возвращает true после успешного ответа пользователя "y",
		// то здесь ничего не происходит. Однако, если пользователь написал "n", наша функция вернёт во внутрь checkRepeatCalculation false, и выполнится условия if - сработает break
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKG, userHeight float64) (float64, error) {
	if userKG <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKG / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userKG float64
	fmt.Print("Введите рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес в киллограммах: ")
	fmt.Scan(&userKG)
	return userKG, userHeight
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите провести еще один расчёт (y/n) ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
