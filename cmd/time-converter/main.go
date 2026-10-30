package main

import "fmt"

func main() {
	var input int
	const minh int = 60
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ошибка: введите целое число минут")
		return
	}
	if input < 0 {
		fmt.Println("Ошибка: оличество минут не может быть отрицательным")
	}
	fmt.Printf("Полных часов: %d\nОставшихся минут: %d\n", input/minh, input%minh)
}
