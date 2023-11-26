/*
Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел,
а возвращает два массива: один из чётных чисел,
второй из нечётных.
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(divider(s))
}

func divider(s []int) (odd []int, even []int) {
	for i, v := range s {
		if v%2 != 0 {
			odd = append(odd, s[i])
		} else {
			even = append(even, s[i])
		}
	}
	return
}
