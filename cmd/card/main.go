package main

import "fmt"

func main() {
	var name string = "Никита"
	group := "Учебная-01"
	const course string = "Бэкенд-разработка на языке Go"
	var tasks int = 1
	interest := true
	fmt.Printf("Участник: %s\nГруппа: %s\nКурс: %s\nЗавершено работ: %d\nИнтерес к бэкенду: %t\n", name, group, course, tasks, interest)
}
