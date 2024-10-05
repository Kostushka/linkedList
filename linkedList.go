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

func findNElToLinked(num int, linked *List) *List {
	if num <= 0 {
		fmt.Println("Порядковый номер искомого элемента не должен быть меньше 1")
		return nil
	}
	// размер связанного списка
	linkedSize := 0
	// здесь будет искомый элемент связанного списка
	result := linked
	// счетчик
	count := num
	for i := linked; i != nil; i = i.prev {
		linkedSize++
		if count == 0 {
			// если счетчик доходит до порядкого номера искомого элемента, и это не конец списка, берем предыдущий элемент списка
			result = result.prev
			count++
		}
		count--
	}
	if num > linkedSize {
		fmt.Printf("Элемент с порядковым номером %d не найден. В связанном списке %d элементов\n", num, linkedSize)
		return nil
	}
	return result
}

func printElToLinked(linked *List) {
	if linked == nil {
		fmt.Println("Связанный список пуст")
		return
	}
	for i := linked; i != nil; i = i.prev {
		fmt.Println(i)
	}
}

func deleteAll(linked **List) {
	for i := *linked; i != nil; i = i.prev {
		*linked = i.prev
		fmt.Println("----")
		printElToLinked(*linked)
		fmt.Println("----")
	}
}

func deleteElToLinked(key int, linked **List) *List {
	var prev *List = nil
	for i := *linked; i != nil; i = i.prev {
		if i.key == key {
			if prev == nil {
				*linked = i.prev
				fmt.Printf("Элемент с ключом %d был удален из связанного списка\n", i.key)
				return i
			}
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

	fmt.Println("****")
	fmt.Printf("Ищем элемент с ключом %d в связанном списке:\n", findElement)
	resFind := findElToLinked(findElement, linked)
	fmt.Println("Адрес найденного элемента:", resFind)
	fmt.Println("****")

	numEl := 10
	fmt.Printf("Ищем элемент с порядковым номером %d в связанном списке:\n", numEl)
	fmt.Printf("Адрес элемента: %v\n", findNElToLinked(numEl, linked))
	fmt.Println("****")

	fmt.Println("Вывести все ключи связанного списка:")
	printElToLinked(linked)
	fmt.Println("****")

	fmt.Printf("Удаляем элемент с ключом %d из связанного списка:\n", findElement)
	fmt.Println(deleteElToLinked(findElement, &linked))
	fmt.Println("****")
	
	fmt.Printf("Проверяем, что элемента с ключом %d нет в связанном списке:\n", findElement)
	findElToLinked(findElement, linked)
	printElToLinked(linked)
	fmt.Println("****")

	fmt.Println("Последовательно удаляем элементы из связаннго списка:")
	deleteAll(&linked)
	fmt.Println("****")
}
