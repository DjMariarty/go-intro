package main

import "fmt"

func numbers() {

	// мой возраст
	my_years := 27

	// год рождения
	year_of_birth := 1998

	// число рождения
	The_number_of_horns := 17

	// количество братьев
	My_brothers := 1

	// количество сестер
	my_sisters := 0

	// количество семьи
	my_family := 6

	// стоимость проезда до места обучения
	the_fare := 43.00

	// текущий год
	current_year := 2025

	// курс доллара
	dollar_exchange_rate := 83.00

	// курс евро
	the_euro_exchange_rate := 90

	// курс биткоина
	bitcoin_exchange_rate := 123986

	// вывод в консоль
	fmt.Println(my_years)
	fmt.Println(year_of_birth)
	fmt.Println(The_number_of_horns)
	fmt.Println(My_brothers)
	fmt.Println(my_sisters)
	fmt.Println(my_family)
	fmt.Println(the_fare)
	fmt.Println(current_year)
	fmt.Println(dollar_exchange_rate)
	fmt.Println(the_euro_exchange_rate)
	fmt.Println(bitcoin_exchange_rate)
	fmt.Println("")

	// Tекущий год минус год рождения
	Q := current_year - year_of_birth

	// Текущий год минус мой возраст
	W := current_year - my_years

	// количество братьев минус количество сестер
	E := My_brothers - my_sisters

	// курс евро, умноженный на 1000
	R := the_euro_exchange_rate * 1000

	// курс доллара, умноженный на 2.5
	T := dollar_exchange_rate * 2.5

	// курс биткоина, разделённый на 10000
	Y := bitcoin_exchange_rate / 10000

	// стоимость проезда до места обучения, разделённая на курс доллара
	U := the_fare / dollar_exchange_rate

	// количество человек в семье минус количество братьев минус количество сестёр
	I := my_family - My_brothers - my_sisters

	// 0, делённый на мой возраст
	O := 0 / my_years

	// 35 минус мой возраст
	P := 35 / my_years

	// вывод в консоль
	fmt.Println(Q)
	fmt.Println(W)
	fmt.Println(E)
	fmt.Println(R)
	fmt.Println(T)
	fmt.Println(Y)
	fmt.Println(U)
	fmt.Println(I)
	fmt.Println(O)
	fmt.Println(P)

}

// func main() {

// 	numbers()
// }
