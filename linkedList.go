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

func createLinkedList() *LinkedList {
	// адрес последнего добавленного элемента связанного списка будет лежать в переменной
	var linked *List
	// связанный список
	return &LinkedList{
		linked: linked,
		size: 0,
	}
}

func (l *LinkedList) addElToLinked(key int) {
	// создали новый элемент связанного списка, в нем лежит адрес предыдущего элемента
	newList := List{key, l.linked}
	// положили адрес нового элемента на место прежнего, прежний адрес остался записан в новом элементе в поле prev
	l.linked = &newList
	l.size++
}

func (l *LinkedList) findElToLinked(key int) (*List, error) {
	// ищем элемент по связанному списку
	for j := l.linked; j != nil; j = j.prev {
		if j.key == key {
			return j, nil
		}
	}
	return nil, fmt.Errorf("Элемент %d не найден", key)
}

func (l *LinkedList) findNElToLinked(num int) (*List, error) {
	if num <= 0 {
		return nil, errors.New("Порядковый номер искомого элемента не должен быть меньше 1")
	}
	if num > l.size {
		return nil, fmt.Errorf("Элемент с порядковым номером %d не найден. В связанном списке %d элементов", num, l.size)
	}
	// здесь будет искомый элемент связанного списка
	result := l.linked
	// счетчик
	count := l.size - num
	for count > 0 {
		result = result.prev
		count--
	}
	return result, nil
}

func (l *LinkedList) printElToLinked() {
	if l.linked == nil {
		fmt.Println("Связанный список пуст")
		return
	}
	for i := l.linked; i != nil; i = i.prev {
		fmt.Println(i)
	}
}

func (l *LinkedList) deleteAll() {
	for i := l.linked; i != nil; i = i.prev {
		l.linked = i.prev
		fmt.Println("----")
		l.printElToLinked()
		fmt.Println("----")
	}
}

func (l *LinkedList) deleteElToLinked(key int) (*List, error) {
	var prev *List = nil
	for i := l.linked; i != nil; i = i.prev {
		if i.key == key {
			// если первый же элемент является искомым
			if prev == nil {
				l.linked = i.prev
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
	// связанный список
	linkedList := createLinkedList()
	
	// добавляем элементы связанного списка в связанный список
	for i := 0; i < 10; i++ {
		linkedList.addElToLinked(i)
	}

	findElement := 1

	fmt.Println("****")
	fmt.Printf("Ищем элемент с ключом %d в связанном списке:\n", findElement)
	resFind, err := linkedList.findElToLinked(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", findElement)
	} else {
		fmt.Println(err)
	}
	fmt.Println("Адрес найденного элемента:", resFind)
	fmt.Println("****")

	numEl := 8
	fmt.Printf("Ищем элемент с порядковым номером %d в связанном списке:\n", numEl)
	result, err := linkedList.findNElToLinked(numEl)
	if err == nil {
		fmt.Printf("Адрес элемента: %v\n", result)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")

	fmt.Println("Вывести все ключи связанного списка:")
	linkedList.printElToLinked()
	fmt.Println("****")

	fmt.Printf("Удаляем элемент с ключом %d из связанного списка:\n", findElement)
	deletedEl, err := linkedList.deleteElToLinked(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d был удален из связанного списка\n", deletedEl)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")
	
	fmt.Printf("Проверяем, что элемента с ключом %d нет в связанном списке:\n", findElement)
	res, err := linkedList.findElToLinked(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", res)
	} else {
		fmt.Println(err)
	}
	linkedList.printElToLinked()
	fmt.Println("****")

	fmt.Println("Последовательно удаляем элементы из связанного списка:")
	linkedList.deleteAll()
	fmt.Println("****")
}
