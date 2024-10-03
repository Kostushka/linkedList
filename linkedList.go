package main

import "fmt"

type List struct {
		key int
		prev *List
}

func addElToLinked(key int, linked **List) {
	// создали новый элемент связанного списка, в нем лежит адрес предыдущего элемента
	newList := List{key, *linked}
	// положили адрес нового элемента на место прежнего, прежний адрес остался записан в новом элементе в поле prev
	*linked = &newList
}

func findElToLinked(key int, linked *List) *List {
	// ищем элемент по связанному списку
	for j := linked; j != nil; j = j.prev {
		if j.key == key {
			fmt.Printf("Элемент с ключом %d найден\n", j.key)
			return j
		}
	}
	fmt.Printf("Элемент %d не найден\n", key)
	return nil
}

func printKeys(linked *List) {
	for i := linked; i != nil; i = i.prev {
		fmt.Println(i)
	}
}

func deleteElToLinked(key int, linked *List) *List {
	var prev *List
	for i := linked; i != nil; i = i.prev {
		if i.key == key {
			// удаляем элемент
			prev.prev = i.prev
			fmt.Printf("Элемент с ключом %d был удален из связанного списка\n", i.key)
			// возвращаем адрес удаленного элемента
			return i
			
		}
		prev = i
	}
	return nil
}

func main() {
	// адрес последнего добавленного элемента связанного списка будет лежать в переменной
	var linked *List
	// добавляем элементы связанного списка в связанный список
	for i := 0; i < 10; i++ {
		addElToLinked(i, &linked)
	}

	findElement := 1

	fmt.Printf("Ищем элемент с ключом %d в связанном списке:\n", findElement)
	resFind := findElToLinked(findElement, linked)
	fmt.Println("Адрес найденного элемента:", resFind)

	fmt.Println("Вывести все ключи связанного списка:")
	printKeys(linked)

	fmt.Printf("Удаляем элемент с ключом %d из связанного списка:\n", findElement)
	fmt.Println(deleteElToLinked(findElement, linked))
	
	fmt.Printf("Проверяем, что элемента с ключом %d нет в связанном списке:\n", findElement)
	findElToLinked(findElement, linked)
	printKeys(linked)
}
