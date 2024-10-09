package main

import (
	"fmt"
	"errors"
)

type List struct {
	key int
	prev *List
}

type LinkedList struct {
	linked *List
	size int
}

func addElToLinked(key int, linkedList *LinkedList) {
	// создали новый элемент связанного списка, в нем лежит адрес предыдущего элемента
	newList := List{key, linkedList.linked}
	// положили адрес нового элемента на место прежнего, прежний адрес остался записан в новом элементе в поле prev
	linkedList.linked = &newList
	linkedList.size++
}

func findElToLinked(key int, linkedList *LinkedList) (*List, error) {
	// ищем элемент по связанному списку
	for j := linkedList.linked; j != nil; j = j.prev {
		if j.key == key {
			return j, nil
		}
	}
	return nil, fmt.Errorf("Элемент %d не найден", key)
}

func findNElToLinked(num int, linkedList *LinkedList) (*List, error) {
	if num <= 0 {
		return nil, errors.New("Порядковый номер искомого элемента не должен быть меньше 1")
	}
	if num > linkedList.size {
		return nil, fmt.Errorf("Элемент с порядковым номером %d не найден. В связанном списке %d элементов", num, linkedList.size)
	}
	// здесь будет искомый элемент связанного списка
	result := linkedList.linked
	// счетчик
	count := linkedList.size - num
	for count > 0 {
		result = result.prev
		count--
	}
	return result, nil
}

func printElToLinked(linkedList *LinkedList) {
	if linkedList.linked == nil {
		fmt.Println("Связанный список пуст")
		return
	}
	for i := linkedList.linked; i != nil; i = i.prev {
		fmt.Println(i)
	}
}

func deleteAll(linkedList *LinkedList) {
	for i := linkedList.linked; i != nil; i = i.prev {
		linkedList.linked = i.prev
		fmt.Println("----")
		printElToLinked(linkedList)
		fmt.Println("----")
	}
}

func deleteElToLinked(key int, linkedList *LinkedList) (*List, error) {
	var prev *List = nil
	for i := linkedList.linked; i != nil; i = i.prev {
		if i.key == key {
			// если первый же элемент является искомым
			if prev == nil {
				linkedList.linked = i.prev
				return i, nil
			}
			// удаляем элемент
			prev.prev = i.prev
			// возвращаем адрес удаленного элемента
			return i, nil
		}
		prev = i
	}
	return nil, fmt.Errorf("Элемент с ключом %d не был найден в связанном списке", key)
}

func main() {
	// адрес последнего добавленного элемента связанного списка будет лежать в переменной
	var linked *List

	// связанный список
	linkedList := LinkedList{linked, 0}
	
	// добавляем элементы связанного списка в связанный список
	for i := 0; i < 10; i++ {
		addElToLinked(i, &linkedList)
	}

	findElement := 1

	fmt.Println("****")
	fmt.Printf("Ищем элемент с ключом %d в связанном списке:\n", findElement)
	resFind, err := findElToLinked(findElement, &linkedList)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", findElement)
	} else {
		fmt.Println(err)
	}
	fmt.Println("Адрес найденного элемента:", resFind)
	fmt.Println("****")

	numEl := 8
	fmt.Printf("Ищем элемент с порядковым номером %d в связанном списке:\n", numEl)
	result, err := findNElToLinked(numEl, &linkedList)
	if err == nil {
		fmt.Printf("Адрес элемента: %v\n", result)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")

	fmt.Println("Вывести все ключи связанного списка:")
	printElToLinked(&linkedList)
	fmt.Println("****")

	fmt.Printf("Удаляем элемент с ключом %d из связанного списка:\n", findElement)
	deletedEl, err := deleteElToLinked(findElement, &linkedList)
	if err == nil {
		fmt.Printf("Элемент с ключом %d был удален из связанного списка\n", deletedEl)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")
	
	fmt.Printf("Проверяем, что элемента с ключом %d нет в связанном списке:\n", findElement)
	res, err := findElToLinked(findElement, &linkedList)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", res)
	} else {
		fmt.Println(err)
	}
	printElToLinked(&linkedList)
	fmt.Println("****")

	fmt.Println("Последовательно удаляем элементы из связанного списка:")
	deleteAll(&linkedList)
	fmt.Println("****")
}
