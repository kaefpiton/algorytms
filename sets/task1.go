//Множества (задание 1)
package main

import "fmt"

/*
Задача 1

Реализовать множество, которое хранит в себе числа
на основе хэш функции. В качестве хэш функции взять
остаток от деления текущего числа на 10

реализовать следующие функции:
- добавить новое число в множество
- найти число в множестве
- удалить число из множества

*/

//множество с размером 10
var SetSize = 10
var Set = make([][] int, SetSize)
//Создает множество
func generateSet(){
	for i := range Set{
		Set[i] = make([]int,0)
	}
}

//находит элемент в множестве
func find(x int) (int, bool) {
	i := x % 10

	for i, value := range Set[i] {
		if value == x{
			return i, true
		}
	}
	return -1, false
}


// Добавляет элемент в множество
func add(x int)  {
	//Без конструкции if было бы реализовано мультимножество
	//(Т.е. элементы входили бы по нескольку раз)
	_,isfound := find(x)
	if !isfound{
		i := x % 10
		Set[i] = append(Set[i], x)
		fmt.Println("add ", x, " in set!")
	}else {
		fmt.Println( x, "exist in set!")
	}

}


//удаляет элемент с сохранением порядка оставшихся элементов
//(стоит отметить, что операция не эффективная)
func deleteOrder(x int){
	pos,isfound := find(x)
	if isfound{
		i := x % 10
		//операция удаления с сохранением порядка
		//(склеиваются 2 слайса без найденного элемента)
		Set[i] = append(Set[i][ :pos], Set[i][pos+1: ]...)
		fmt.Println( x, "was remove from set!")
	}else {
		fmt.Println( x, "not exist in set!")
	}
}

//удаление элемента с изменением порядка
//зато эффективно
func delete(x int){
	pos,isfound := find(x)
	if isfound {
		i := x % 10

		//Ставим найденный элемент в конец и
		//создаем новый срез без конечного элемента
		endEl := len(Set[i]) - 1
		Set[i][pos] = Set[i][endEl]
		Set[i] = Set[i][:endEl]
	}else {
		fmt.Println( x, "not exist in set!")
	}

}

func main() {

	generateSet()
	add(3)
	add(13)
	add(33)
	add(53)
	add(63)
	delete(33)
	delete(3)
	delete(13)


	for i := range Set {
		fmt.Println(i, Set[i])
	}


}
