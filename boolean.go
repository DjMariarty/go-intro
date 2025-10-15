package main

import "fmt"

func boolean() {

	// Bool — логический тип данных (булев тип, булевый тип, от англ. Boolean) в информатике.
	//  Присутствует в большинстве языков программирования как самостоятельная сущность или реализуется через численный тип данных

	// прямое присваивание bool
	var name bool = true
	fmt.Println(name)

	// прямое присваивание else
	var zero bool = false
	fmt.Println(zero)

	// вычисление значения bool
	var score bool = 5 == 0
	var score1 bool = 10 < 5

	fmt.Println(score)
	fmt.Println(score1)

	// константы с подходящими значениями
	const I_Am_A_Student bool = false
	const Is_Spring bool = false
	const gols_beauty bool = true
	const const_can_be_changed bool = false
	const i_am_from_Grozny bool = true

	fmt.Println(I_Am_A_Student)
	fmt.Println(Is_Spring)
	fmt.Println(gols_beauty)
	fmt.Println(const_can_be_changed)
	fmt.Println(i_am_from_Grozny)

}

// func main() {
// 	boolean()
// }
