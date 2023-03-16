package main

import (
	"fmt"
	"strconv"
)

func main() {
	romanMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	romanNumRim := []string{"L", "XL", "X", "IX", "V", "IV", "I"}
	romanNumArr := []int{50, 40, 10, 9, 5, 4, 1}
	var firstNum, secondNum, argument string
	fmt.Print("Введите первое число, знак операции и втрое число: ")
	fmt.Scan(&firstNum, &argument, &secondNum)
	a := romanMap[firstNum]
	b := romanMap[secondNum]
	var arabic bool
	if b == 0 || a == 0 {
		var errNumOne, errNumTwo error
		a, errNumOne = strconv.Atoi(firstNum)
		b, errNumTwo = strconv.Atoi(secondNum)
		arabic = true
		if errNumOne != nil || errNumTwo != nil {
			panic("Используются одновременно разные системы счисления, не целые числа или числа вне доступного диапазона.")
		}
	}
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Одно из чисел меньше 1 или больше 10")
	}
	var res int
	var aRes string
	switch argument {
	case "/":
		res = a / b
	case "*":
		res = a * b
	case "+":
		res = a + b
	case "-":
		res = a - b
	default:
		panic("Не соответствует предоставленные арифметическим операциям (*,/,+,-)")
	}
	if arabic {
		fmt.Println(res)
	} else {
		for i := 0; res > 0; {
			if res != res%romanNumArr[i] {
				aRes += romanNumRim[i]
				res -= romanNumArr[i]
				i += 0
			} else {
				i++
			}
		}
		if aRes == "" {
			fmt.Println("Результатом операций с римскими числами не может быть ноль или отрицательное число")
		}
		fmt.Println(aRes)
	}
}
