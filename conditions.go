package main

import "fmt"

func conditions() {

	// задание 1
	const userAge = 20

	/*
	   напиши ниже условие, которое выведет "ты пацан"
	   если userAge меньше 35
	*/

	// тут твой код
	if userAge < 35 {
		fmt.Println("Ты пацан")
	}

	// задание 2
	const isAdmin = false

	/*
	   напиши ниже условие, которое выведет "доступ запрещен",
	   если значение константы isAdmin ложно
	*/
	// тут твой код

	if isAdmin == false {
		fmt.Println("доступ запрещен")
	}

	// задание 3
	const password = "sherlock"
	const realPassword = "holmes"

	/*
	   напиши ниже условие, которое выведет "правильный пароль"
	   только если значения password и realPassword совпадают.
	*/
	// тут твой код

	if password == realPassword {
		fmt.Println("Правильный пароль")

	}

	// задание 4
	const x = 2.7
	const y = -6.0

	/*
	   напиши ниже условие, которое выведет "икс меньше игрик"
	   только если x < y
	*/
	// тут твой код

	if x < y {
		fmt.Println("икс меньше игрик")
	}
}

// func main() {

// 	conditions()
// }
