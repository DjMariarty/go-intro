package main

import (
	"fmt"
	"strings"
)

func conditions_with_else() {
	// задание 1
	const userAge = 20

	/*
	   напиши ниже условие, которое выведет "ты пацан"
	   если userAge меньше 35, иначе выведет "не по-пацански"
	*/
	// тут твой код

	if userAge < 35 {
		fmt.Println("Ты пацан")
	} else {
		fmt.Println("не по-пацански")
	}

	// задание 2
	const isUser = false

	/*
	   напиши ниже условие, которое выведет "привет, юзер", если
	   значение isUser истинно, иначе — "не пользователь ты мне"
	*/
	// тут твой код

	if isUser == true {
		fmt.Println("Привет юзер")
	} else {
		fmt.Println("не пользователь ты мне")
	}

	// задание 3
	const password = "sherlock"
	const realPassword = "holmes"

	/*
	   напиши ниже условие, которое выведет "правильный пароль"
	   если password == realPassword,
	   иначе — "неверный пароль".
	*/
	// тут твой код

	if password == realPassword {
		fmt.Println("правильный пароль")
	} else {
		fmt.Println("неверный пароль")
	}

	// задание 4 (со звёздочкой)
	const email = "admin@intocode.ru"

	/*
	   выведи "неверный эмайл", если в email нет символа "@",
	   иначе выведи "ты зареган".
	   Подсказка: используй strings.Contains(email, "@"). Выше есть ссылка на материал
	*/
	// тут твой код

	if strings.Contains(email, "@") {
		fmt.Println("неверный эмайл")
	} else {
		fmt.Println("ты зареган")
	}

	// задание 5
	const x = 2.7
	const y = -6.0

	/*
	   выведи значение переменной x, если оно больше y,
	   иначе выведи значение y
	*/
	// тут твой код

	if x > y {
		fmt.Println(x)

	} else {
		fmt.Println(y)
	}

}

func main() {
	conditions_with_else()
}
