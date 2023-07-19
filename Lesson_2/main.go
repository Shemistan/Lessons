package main

// Что это?
// go mod init github.com/Shemistan/Lessons
// go mod tidy

import "fmt"

func main() {
	// Пример 1: Неточный ответ при сложении
	a := 0.1
	b := 0.2
	sum := a + b
	fmt.Println(sum) // Ожидаемый ответ: 0.3, фактический ответ: 0.30000000000000004

	// Пример 2: Неточный ответ при делении
	c := 1.0
	d := 3.0
	quotient := c / d
	fmt.Println(quotient) // Ожидаемый ответ: 0.3333333333333333, фактический ответ: 0.33333333333333337
}
