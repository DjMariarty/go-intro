package main

import "fmt"

func all_my_skills() {
	// переменные
	var text string
	var score int
	var drob float64

	const tride string = "Hello"
	const punk int = 5
	name := "Shadid"
	age := 27

	// вывод в консоль
	fmt.Println(text)
	fmt.Println(score)
	fmt.Println(drob)
	fmt.Println(tride)
	fmt.Println(punk)
	fmt.Printf("Я %s, мне %d лет \n", name, age)

	// арифметические операторы
	X := text + tride
	Y := score * punk

	fmt.Println(X)
	fmt.Println(Y)

	Stroke := "Привет, мне 25 лет \n и я учусь в интукод"
	help := `Все получится \n главное стараться`

	fmt.Println(Stroke)
	fmt.Println(help)

}

// func main() {

// 	all_my_skills()
// }
