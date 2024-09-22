package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	var number = rand.Intn(10)
	fmt.Println(number)
	fmt.Println("Привет,это игра 'Угадай число'.У тебя 3 попытки:) ")
	fmt.Println("Введите ваше число")
	count := 0
	player_number1 := -1

	for count < 3 && number != player_number1 {

		fmt.Scan(&player_number1)
		if count >= 1 && number != player_number1 {
			if number < player_number1 {
				fmt.Println("Загаданное число меньше")
			} else if number > player_number1 {
				fmt.Println("Загаданное число больше")
			}
		} else if count == 0 {
			fmt.Println("Подумай еще")
		}
		count++

	}
	if number == player_number1 {
		fmt.Println("Верно!")
	} else {
		fmt.Println("Упс...ты проиграл")
	}

}
