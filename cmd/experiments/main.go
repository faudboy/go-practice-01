package main

import "fmt"

func main() {
	const lang = "Go"
	var name string = "Никита"
	completed := 1
	var ready bool = true
	var hours float64 = 4.23

	fmt.Println("Изучаем", lang)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено %d\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность %.2f часа\n", hours)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)
}
