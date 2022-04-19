package main

import (
	"fmt"
	"math"
)
/*
 Сортировка подсчетом

 Данна последовательность оценок: 5453215

 Создадим множество значений данной
 последовательности(0 входит в данное
 множество для удобства индексации)

 Далее выводим по очереди элементы множества
 с учетом их количества

 P.S. Данный метод сортировки хорошо подходи
	  при небольшой разнице между минимальным
	  и максимальным значением в последовательности
 */

func findMin(array []int) int {
	min := array[0]

	for _,val := range array{
		if val < min{
			min = val
		}
	}

	return min
}
func findMax(array []int) int {
	max := array[0]

	for _,val := range array{
		if val > max{
			max = val
		}
	}

	return max
}


func countSort(seq []int)  {
	minval := findMin(seq)
	maxval := findMax(seq)
	fmt.Println("min  = ", minval)
	fmt.Println("max  = ", maxval)

	//модуль принимает и возвращает только float64
	//как варинт можно написать свою функцию
	//определения модуля или расширить Math
	k := int(math.Abs(float64(maxval - minval))) + 1
	count := make([]int,k)

	for _,val := range(seq){
		count[val - minval] += 1
	}
	newpos := 0 // нужен для добавления каждого нового элемента
	fmt.Println(count)
	for i:= 0; i < k; i++{
		for j := 0; j < count[i]; j++{
			seq[newpos] = i + minval
			newpos++
		}
	}

	fmt.Println(seq)
}


func main() {
	fmt.Println("start")

	seq := []int{5,4,5,3,2,1,5}
	countSort(seq)
}
