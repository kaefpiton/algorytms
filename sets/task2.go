package main

import "fmt"

/*
Задача 2

Дана последовательность положительныхN чисел. Найти такие
2 различных числа из последовательности N, чтоб
их сумма удовлетворяла заданному Х

 */

//Решим задачу плохим по скорости способом ( за O(N^2) )
func badSolution(sequence[]int , x int){

	for i := 0 ; i < len(sequence) - 1; i++{
		for j := i + 1 ; j < len(sequence) ; j++{
			if sequence[i] + sequence[j] == x{
				fmt.Println( sequence[i], "+" ,sequence[j], "=" , x)
			}
		}
	}

}

func findEl(set[]int, x int) bool{

	for _, value := range set {
		if value == x{
			return true
		}
	}
	return false
}

//Идея состоит в том чтобы вычетать элемент из
//исходного X и сравнивать с составленным множеством
func goodSolution(sequence[]int , x int){
	set := []int{}
	for _,val := range sequence{
		if findEl(set,x - val)  {
			fmt.Println(val, "+", x - val, "=", x)
		}else {
			set = append(set,val)
		}
	}
}


func main()  {
	fmt.Println("start")
	sequence := []int{3,5,7,11,5}
	goodSolution(sequence, 10)

}
