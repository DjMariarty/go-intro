package main

import "fmt"

func main() {

	// название любимого фильма
	const film string = "Джанго освобожденный"
	fmt.Println(film)

	// приложение, которым чаще всего пользуюсь
	const app string = "Whatsapp"
	fmt.Println(app)

	// ссылка на страницу go в википедии
	const language string = "https://ru.wikipedia.org/wiki/Go"
	fmt.Println(language)

	// текст lorem ipsum в сырой строке
	var text string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	 Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	fmt.Println(text)

}
