//Guess - игра, в который игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Получаем текущую дату и в ремя в виде целого чилса
	rand.Seed(seconds)           //Инициализируем генератор случайных чисел
	target := rand.Intn(100) + 1 // Генерируем целое число от 1 до 100
	fmt.Println("Я выбрал случайное число от 1 до 100.")
	fmt.Println("Можете ли вы угадать его?")
	//	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) // Создаём bufio.Reader для чтения ввода с клавиатуры

	success := false // Настроить программу, что бы по умолчанию выводилось сообщение о проигрыше

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У вас осталось", 10-guesses, "догадок.")

		fmt.Print("Сделайте предположение: ") // Запрашиваем число
		input, err := reader.ReadString('\n') // Прочитать данные, введенные пользователем до нажатия Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // Удаляем символ новой строки
		guess, err := strconv.Atoi(input) // Введенная строка преобразуется в целое число
		if err != nil {
			log.Fatal(err)
		}
		if guess < target { //Если введенное значение меньше загаданного, сообщить об этом
			fmt.Println("Ой, ваша догадка была занижена!")
		} else if guess > target { // Если введенное значение больше загаданного, сообщить об этом
			fmt.Println("Ой, ваша догадка была завышенна")
		} else { //В противном случае введенное значение должно быть правильным...
			success = true // Предотвращаем вывод сообщения о проигрыше
			fmt.Println("Хорошая работа! Вы угадали!")
			break // Выход из цикла
		}
	}
	if !success { //Есди переменная success равна false, сообщить игроку загаданное число
		fmt.Println("Извините, вы не угадали число, Это было:", target)
	}
}
