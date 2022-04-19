package main

import "fmt"

/*
Задача:
Даны 2 (целые положительные)числа X и Y
Проверить, можно ли получить X
из Y путем перестановки цифр в X
 */

func countDigits(num int) [10]int {
	digits := [10]int{}

	for num != 0{
		nextnum := num % 10
		digits[nextnum] = digits[nextnum] + 1
		num = num / 10
	}

	return digits
}

func main(){
	var X  = 1233
	var Y = 1323
	fmt.Println("X = ", X)
	fmt.Println("Y = ", Y)

	xDigits := countDigits(X)
	yDigits := countDigits(Y)

	//fmt.Println(xDigits)
	//fmt.Println(yDigits)

	if xDigits == yDigits{
		fmt.Println("Число Y можно составить из цифр X")
	}else {
		fmt.Println("Число Y нельзя составить из цифр X")
	}

}