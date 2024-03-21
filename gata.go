package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var romanMap = []struct {
	decVal int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	{3, "III"}, {2, "II"}, {6, "VI"}, {7, "VII"}, {8, "VIII"},
}

// функция для перевода из арабских чисел в римские
func decimalToRomanIterative(num int) string {
	result := ""
	for _, pair := range romanMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

// фунция для проверки является ли строка арабским числом
func CheckingForNumbers(n string) bool {
	for i := 0; i < len(n); i++ {
		if unicode.IsDigit([]rune(n)[i]) != true {
			return false
		}
	}
	return true
}

// функция для перевода строки в число
func Numbers(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

// функция для перевода числа в строку
func String(st int) string {
	line := strconv.Itoa(st)
	return line
}

// функция проверят является ли первый элемент массива римским числом. Если "нет" - выводит 0, если "да" - выводит число
func CheakingRomanNumerals1(crn []string) int {
	number1 := 0
	for _, c := range crn[0] {
		k1 := strconv.Itoa(number1)
		for _, r := range romanMap {
			if r.symbol == string(c) {
				for i := len(k1) - 1; i < len(k1); i++ {
					if string(k1[i]) == "1" && r.symbol != "I" {
						number1 -= 2
					}
				}
				number1 += r.decVal
			}
		}
	}
	return number1
}

// функция проверят является ли третий элемент массива римским числом. Если "нет" - выводит 0, если "да" - выводит число
func CheakingRomanNumerals2(crn []string) int {
	number2 := 0
	for _, c := range crn[2] {
		k2 := strconv.Itoa(number2)
		for _, r := range romanMap {
			if r.symbol == string(c) {
				for i := len(k2) - 1; i < len(k2); i++ {
					if string(k2[i]) == "1" && r.symbol != "I" {
						number2 -= 2
					}
				}
				number2 += r.decVal
			}
		}
	}

	return number2
}

// функция проверят является ли второй элемент массива математической операцией
func MathematicalCalculations(m []string) bool {
	switch m[1] {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	default:
		return false
	}
}

// функция порвепяет сколько математических операций ввердено в строку
func StringsCount(sc string) bool {
	count := 0
	for i := 0; i < len(sc); i++ {
		if string(sc[i]) == "+" || string(sc[i]) == "-" || string(sc[i]) == "*" || string(sc[i]) == "/" {
			count++
		}
	}
	if count == 1 {
		return false
	}
	return true
}

// фунеция проверяется диапозон арабских чисел
func CheckingForARangeArNum(cr []string) bool {
	if Numbers(cr[0]) < 1 || Numbers(cr[0]) > 10 || Numbers(cr[2]) < 1 || Numbers(cr[2]) > 10 {
		return false
	}
	return true
}

// функция проверяет диапозон римских чисел
func CheckingForARangeRomNum(cr []string) bool {
	if CheakingRomanNumerals1(cr) < 0 || CheakingRomanNumerals2(cr) < 0 || CheakingRomanNumerals1(cr) > 10 || CheakingRomanNumerals2(cr) > 10 {
		return false
	}
	return true
}

// функция считает арабские числа
func AnswerArabicNumerals(an []string) string {
	var answer string
	switch an[1] {
	case "+":
		answer1 := Numbers(an[0]) + Numbers(an[2])
		answer = String(answer1)
	case "-":
		answer2 := Numbers(an[0]) - Numbers(an[2])
		answer = String(answer2)
	case "*":
		answer3 := Numbers(an[0]) * Numbers(an[2])
		answer = String(answer3)
	case "/":
		answer4 := Numbers(an[0]) / Numbers(an[2])
		if Numbers(an[2]) != 0 {
			answer = String(answer4)
		}
	}
	return answer
}

// фунция считает римские числа
func AnswerRomanNumbers(arn []string) string {
	var answer string
	switch arn[1] {
	case "+":
		answer1 := CheakingRomanNumerals1(arn) + CheakingRomanNumerals2(arn)
		answer = decimalToRomanIterative(answer1)
	case "-":
		answer2 := CheakingRomanNumerals1(arn) - CheakingRomanNumerals2(arn)
		if answer2 >= 1 {
			answer = decimalToRomanIterative(answer2)
		} else if answer2 == 0 {
			answer = "Выдача паники, так как в римской системе нет числа  '0'!"
		} else {
			answer = "Выдача паники, так как в римской системе нет отрицательных чисел!"
		}
	case "*":
		answer3 := CheakingRomanNumerals1(arn) * CheakingRomanNumerals2(arn)
		answer = decimalToRomanIterative(answer3)
	case "/":
		answer4 := CheakingRomanNumerals1(arn) / CheakingRomanNumerals2(arn)
		if float32(CheakingRomanNumerals1(arn))/float32(CheakingRomanNumerals2(arn)) >= 1 {
			answer = decimalToRomanIterative(int(answer4))
		} else {
			answer = "Выдача паники, так как значение меньше 1!"
		}
	}
	return answer
}

// функция проверят ввод данных. обрабатывает ошибки или выдает ответ, если данные введены корректно
func CheakingForCorrect(sc []string, s string) string {
	st := " "
	if (CheakingRomanNumerals1(sc) == 0 || CheakingRomanNumerals2(sc) == 0) || (CheckingForNumbers(sc[0]) == false || CheckingForNumbers(sc[2]) == false) && len(sc) >= 3 {
		st = "Выдача паники, так как используются неизвестные системы счисления!"
	}
	if (CheckingForARangeArNum(sc) == false && CheckingForNumbers(sc[0]) == true && CheckingForNumbers(sc[2]) == true) || (CheckingForARangeRomNum(sc) == false && CheakingRomanNumerals1(sc) != 0 && CheakingRomanNumerals2(sc) != 0) {
		st = "Выдача паники, число не удовлетворяет диапозону от 1 до 10!"
	}
	if StringsCount(s) == true && len(sc) > 3 {
		st = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)!"
	}
	if MathematicalCalculations(sc) != true && len(sc) == 3 {
		st = "Выдача паники, так как используются неизвестные математические действия!"
	}
	if (CheakingRomanNumerals1(sc) != 0 && CheckingForNumbers(sc[2]) == true) || (CheakingRomanNumerals2(sc) != 0 && CheckingForNumbers(sc[0]) == true) && len(sc) == 3 {
		st = "Выдача паники, так как используются одновременно разные системы счисления!"
	}
	if CheakingRomanNumerals1(sc) != 0 && CheakingRomanNumerals2(sc) != 0 && MathematicalCalculations(sc) == true && len(sc) == 3 && CheckingForARangeRomNum(sc) == true {
		st = AnswerRomanNumbers(sc)
	}
	if CheckingForNumbers(sc[0]) == true && CheckingForNumbers(sc[2]) == true && CheckingForARangeArNum(sc) == true && StringsCount(s) != true && len(sc) == 3 {
		st = AnswerArabicNumerals(sc)
	}
	return st
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	check := strings.Split(strings.Trim(s, "\n"), " ")
	if len(check) < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией!")
	} else if CheakingForCorrect(check, s) != "" && len(check) >= 3 {
		check[2] = check[2][:len(check[2])-1]
		fmt.Println(CheakingForCorrect(check, s))
	}

}
