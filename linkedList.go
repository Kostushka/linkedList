package main

import (
	"fmt"
	"errors"
)

type Item struct {
	key int
	prev *Item
}

type LinkedList struct {
	linked *Item
	size int
}

func newLinkedList() *LinkedList {
	// связанный список
	return &LinkedList{}
}

func (l *LinkedList) addEl(key int) {
	// создали новый элемент связанного списка, в нем лежит адрес предыдущего элемента
	// положили адрес нового элемента на место прежнего, прежний адрес остался записан в новом элементе в поле prev
	l.linked = &Item{
		key: key, 
		prev: l.linked,
	}
	l.size++
}

func (l *LinkedList) findEl(key int) (*Item, error) {
	// ищем элемент по связанному списку
	for n := l.linked; n != nil; n = n.prev {
		if n.key == key {
			return n, nil
		}
	}
	return nil, fmt.Errorf("Элемент %d не найден", key)
}

func (l *LinkedList) findNEl(num int) (*Item, error) {
	if num <= 0 {
		return nil, errors.New("Порядковый номер искомого элемента не должен быть меньше 1")
	}
	if num > l.size {
		return nil, fmt.Errorf("Элемента с порядковым номером %d в списке нет. В связанном списке %d элементов", num, l.size)
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

func (l *LinkedList) printEl() {
	if l.linked == nil {
		fmt.Println("Связанный список пуст")
		return
	}
	for n := l.linked; n != nil; n = n.prev {
		fmt.Println(n)
	}
}

func (l *LinkedList) deleteAll() {
	for n := l.linked; n != nil; n = n.prev {
		l.linked = n.prev
		fmt.Println("----")
		l.printEl()
		fmt.Println("----")
	}
}

func (l *LinkedList) deleteEl(key int) (*Item, error) {
	var prev *Item = nil
	for n := l.linked; n != nil; n = n.prev {
		if n.key == key {
			// если первый же элемент является искомым
			if prev == nil {
				l.linked = n.prev
				return n, nil
			}
			// удаляем элемент
			prev.prev = n.prev
			// возвращаем адрес удаленного элемента
			return n, nil
		}
		prev = n
	}
	return nil, fmt.Errorf("Элемент с ключом %d не был найден в связанном списке", key)
}

func main() {
	// связанный список
	linkedList := newLinkedList()
	
	// добавляем элементы связанного списка в связанный список
	for i := 0; i < 10; i++ {
		linkedList.addEl(i)
	}

	findElement := 1

	fmt.Println("****")
	fmt.Printf("Ищем элемент с ключом %d в связанном списке:\n", findElement)
	resFind, err := linkedList.findEl(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", findElement)
	} else {
		fmt.Println(err)
	}
	fmt.Println("Адрес найденного элемента:", resFind)
	fmt.Println("****")

	numEl := 8
	fmt.Printf("Ищем элемент с порядковым номером %d в связанном списке:\n", numEl)
	result, err := linkedList.findNEl(numEl)
	if err == nil {
		fmt.Printf("Адрес элемента: %v\n", result)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")

	fmt.Println("Вывести все ключи связанного списка:")
	linkedList.printEl()
	fmt.Println("****")

	fmt.Printf("Удаляем элемент с ключом %d из связанного списка:\n", findElement)
	deletedEl, err := linkedList.deleteEl(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d был удален из связанного списка\n", deletedEl)
	} else {
		fmt.Println(err)
	}
	fmt.Println("****")
	
	fmt.Printf("Проверяем, что элемента с ключом %d нет в связанном списке:\n", findElement)
	res, err := linkedList.findEl(findElement)
	if err == nil {
		fmt.Printf("Элемент с ключом %d найден\n", res)
	} else {
		fmt.Println(err)
	}
	linkedList.printEl()
	fmt.Println("****")

	fmt.Println("Последовательно удаляем элементы из связанного списка:")
	linkedList.deleteAll()
	fmt.Println("****")
}
