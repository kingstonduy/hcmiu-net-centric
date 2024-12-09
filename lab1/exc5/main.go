package main

import (
	"fmt"

	"math/rand"
)

func isValid(s string) bool {
	stack := []rune{}

	bracketMap := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	for _, char := range s {
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, char)
		} else if char == ']' || char == '}' || char == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	testCases := []string{
		"{[()]}()[]",
		"{[(])}",
		"{(([])[])[]}",
		"{[({})]}",
		"{[(])}",
		"{[}",
		"abc[]{}",
		"[{]}(){}",
		"((([])))",
		"((()))",
		"(",
		"]",
	}

	for _, testCase := range testCases {
		fmt.Printf("Bracket string =%s. Is valid=%+v\n", testCase, isValid(testCase))
	}

	for i := 1; i <= 100; i++ {
		str := GenerateRandomBrackets(50)
		fmt.Printf("Bracket string =%s. Is valid=%+v\n", str, isValid(str))
	}
}

const (
	openingBrackets = "([{"
	closingBrackets = ")]}"
)

func GenerateRandomBrackets(length int) string {
	brackets := make([]rune, length)

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			brackets[i] = rune(openingBrackets[rand.Intn(len(openingBrackets))])
		} else {
			switch brackets[i-1] {
			case '(':
				brackets[i] = ')'
			case '{':
				brackets[i] = '}'
			case '[':
				brackets[i] = ']'
			}
		}
	}

	for i := range brackets {
		j := rand.Intn(i + 1)
		brackets[i], brackets[j] = brackets[j], brackets[i]
	}

	return string(brackets)
}
