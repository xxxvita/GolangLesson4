package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Захотелось попробовать так запустить функцию получения строки с массивом
// Заранее понимаю, что подход кривущий
func getMethod() func() (string, error) {
	if len(os.Args) > 1 {
		return getDataFromArgs
	} else {
		return getDataFromStdin
	}
}

func getDataFromArgs() (string, error) {
	return strings.Join(os.Args[1:], ","), nil
}

// Чтение строк с массивами чисел до успеха или до обрыва программы
func getDataFromStdin() (string, error) {
	fmt.Printf("Введите строку элементов типа int через запятую:\n")

	txt := ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt = scanner.Text()

		if len(txt) > 1 {
			break
		}
	}

	return txt, nil
}

func StringToArrayInt(str string) (mas []int64, err error) {
	if len(str) == 0 {
		err = errors.New("Длина строки с массивом равна 0")
		return
	}

	var i int64

	array_str := strings.Split(str, ",")
	for _, s := range array_str {
		i, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			mas = nil
			return
		}

		mas = append(mas, i)
	}

	return
}

// Функция для управления вспомогательным выводом
func debug(s string) {
	if 1 == 2 {
		fmt.Println(s)
	}
}

// сортировка массива методом вставок
func sort(input_mas []int64) (result []int64) {
	result = make([]int64, len(input_mas))
	copy(result, input_mas[:])
	debug(fmt.Sprintf("Массив перед сортировкой: %v\n", result))

	for i := 1; i < len(result); i++ {
		debug(fmt.Sprintf("i: %d, массив: %v", i, result))

		// Проверка: стоит ли текущий элемент на своём месте
		// или требуется его вставить в раннее сформированный участок [<i:]
		if result[i] < result[i-1] {
			debug(fmt.Sprintf("Элемент с номером %d (result[i]==%d) меньше предыдущего элемента %d", i, result[i], result[i-1]))

			el := result[i]
			j := 0

			// поиск элемента меньшего или равного i-тому
			for j = i - 1; j >= 0; j-- {
				if result[j] <= el {
					break
				}
			}

			if j >= 0 {
				debug(fmt.Sprintf("Для вставки определён элемент j: %d (result[j]==%d)", j, result[j]))
			} else {
				debug(fmt.Sprintf("Элемент будет добавлен в начало массива"))
			}

			debug(fmt.Sprintf("--------До удаления элемента массив: %v", result))
			// Удаляется элемент el
			result = append(result[:i], result[i+1:]...)
			debug(fmt.Sprintf("--------После удаления элемента массив: %v", result))
			// Слайс сдвигается, создавая дыру для вставки, найденного элемента
			if j < 0 {
				result = append(append([]int64{}, el), result[:]...)
			} else {
				result = append(result[:j+1], append(append([]int64{}, el), result[j+1:]...)...)
			}
			debug(fmt.Sprintf("--------После вставки элемента массив: %v", result))
		}

		debug(fmt.Sprintf("-----------------------------------------------------------------"))
	}

	return
}

func main() {
	fmt.Println("Программа сортирует массив чисел, заданный в консоли (или через аргументы)")

	// Если в программу передали входящие аргументы (массив), то чтение
	// осуществляется из них, если нет, то из os.Stdin
	str, err := getMethod()()
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}

	// Преобразование полученной строки в массив целых чисел
	mas, err := StringToArrayInt(str)
	if err != nil {
		fmt.Printf("Ошибка ввода массива: %s\n", err.Error())
		os.Exit(1)
	}

	// Сортировка массива
	mas = sort(mas)
	// Вывод результата
	fmt.Printf("Отсортированный массив: %v\n", mas)
}
