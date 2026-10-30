package main

import "fmt"

func main() {
	var t1, t2, t3 int = 20, 25, 35
	fmt.Printf("Сумма: %d\n", t1+t2+t3)
	var f1, f2, f3 = float64(t1), float64(t2), float64(t3)
	fmt.Printf("Среднее значение: %.2f\n", (f1+f2+f3)/3)
}
