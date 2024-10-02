package main

import "fmt"

func main() {
	type List struct {
		key int
		prev *List
	}
	// первый элемент связанного списка
	list := List{0, nil}
	// адрес последнего элемента связанного списка будет лежать в переменной
	var linked *List
	// добавили первый элемент связанного списка в связанный список
	linked = &list
	
	for i := 1; i < 10; i++ {
		// создали новый элемент связанного списка, в нем лежит адрес предыдущего элемента
		newList := List{i, linked}
		// положили адрес нового элемента на место прежнего, прежний адрес остался записан в новом элементе в поле prev
		linked = &newList
	}
	findElement := 47
	isFindElement := false
	// ищем элемент по связанному списку
	for j := linked; j != nil; j = j.prev {
		fmt.Println(j)
		if j.key == findElement {
			fmt.Printf("Элемент %d найден\n", j.key)
			isFindElement = true
		}
	}
	if !isFindElement {
		fmt.Printf("Элемент %d не найден\n", findElement)
	}
}
