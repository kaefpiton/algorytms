package main

import "fmt"

/*
Задача

Реализовать множество, которое хранит в себе числа
на основе хэш функции. В качестве хэш функции взять
остаток от деления текущего числа на 10

реализовать следующие функции:
- добавить новое число в множество
- найти число в множестве
- удалить число из множества

*/

func add()  {
	fmt.Println("add")
}

func find()  {
	fmt.Println("find")
}

func delete(){
	fmt.Println("delete")
}

func main() {

	nums := []int{9}
	nums = append(nums,9,4 )
	fmt.Println(len(nums),	nums)
	//find()
}
