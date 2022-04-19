package main

import "fmt"

func sort(slice []int)[]int{
	if len(slice) < 2{
		return slice
	}

	pivot := slice[0]
	var greaterpivot = []int{}
	var lesspivot = []int{}

	for i := 1; i < len(slice); i++ {
		if slice[i] <= pivot{
			lesspivot = append(lesspivot, slice[i])
		}else {
			greaterpivot = append(greaterpivot, slice[i])
		}
	}

	return append(append(sort(lesspivot),pivot),sort(greaterpivot)...)
}



func main() {
	fmt.Println("Quick sort start")
	slice := []int{1,5,2,3}

	fmt.Println(sort(slice))
}
