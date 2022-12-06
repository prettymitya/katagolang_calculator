package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
reading from bufio string
from int to string
delete all spaces
*/
func main() {
	var num1, num2, operator string
	count := 0
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, " ", "")
	for j := 0; j < len(str); j++ {
		y := string(str[j])
		if y == "+" || y == "-" || y == "/" || y == "*" {
			count++
		}
	}
	if count != 1 {
		fmt.Println("Вывод ошибки, так как знак не соответствует условию")
		os.Exit(1)
	}
	for i := 0; i < len(str); i++ {
		x := string(str[i])
		if x == "+" || x == "-" || x == "/" || x == "*" {
			operator = x
			break
		}
		num1 += x
	}
	for j := len(str) - 1; j != 0; j-- {
		x := string(str[j])
		if x == "+" || x == "-" || x == "/" || x == "*" {
			break
		}
		num2 = x + num2
	}

	rome1, rome2 := RtoI(num1), RtoI(num2)
	if rome1 == 0 || rome2 == 0 {
		num1ar, _ := strconv.Atoi(num1)
		num2ar, _ := strconv.Atoi(num2)
		error1(num1ar, num2ar)
		var sumarab1 = operations(operator, num1ar, num2ar)
		fmt.Println(sumarab1)
	} else {
		var sumrome1 = operations(operator, rome1, rome2)
		if sumrome1 <= 0 {
			fmt.Println("Вывод ошибки, так как в Римской системе нет отрицательных")
			os.Exit(1)
		}
		error1(rome1, rome2)
		sumrome2 := resultRtoI(sumrome1)
		fmt.Println(sumrome2)
	}
}

/*
Rome to Arabic
*/
func RtoI(str string) int {
	switch {
	case str == "X":
		return 10
	case str == "IX":
		return 9
	case str == "VIII":
		return 8
	case str == "VII":
		return 7
	case str == "VI":
		return 6
	case str == "V":
		return 5
	case str == "IV":
		return 4
	case str == "III":
		return 3
	case str == "II":
		return 2
	case str == "I":
		return 1
	default:
		return 0
	}
}

func resultRtoI(sum int) string {
	result := []struct {
		begin int
		end   string
	}{
		{100, "C"},
		{90, "XC"},
		{80, "LXXX"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var resultati string
	for _, result1 := range result {
		for sum >= result1.begin {
			resultati += result1.end
			sum -= result1.begin
		}
	}
	return resultati
}

func operations(operator string, num1, num2 int) int {
	var sum int
	switch {
	case operator == "+":
		sum = num1 + num2
	case operator == "-":
		sum = num1 - num2
	case operator == "*":
		sum = num1 * num2
	case operator == "/":
		sum = num1 / num2
	}
	return sum
}
func error1(num1ar, num2ar int) {
	if num1ar <= 0 || num1ar > 10 || num2ar <= 0 || num2ar > 10 {
		fmt.Println("Вывод ошибки, так как пример не соответствует условию")
		os.Exit(1)
	}
}
