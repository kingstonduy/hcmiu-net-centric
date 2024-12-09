package main

import (
	"fmt"
	"strings"
	"unicode"

	"math/rand"
)

func IsValidLuhn(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	if len(number) <= 1 {
		return false
	}
	for _, r := range number {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	sum := 0
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}

func main() {
	validNumber := "4539 3195 0343 6467"
	fmt.Printf("credit card= %s. Is valid=%v\n", validNumber, IsValidLuhn(validNumber))

	invalidNumber := "8273 1232 7352 0569"
	fmt.Printf("credit card= %s. Is valid=%v\n", invalidNumber, IsValidLuhn(invalidNumber))

	for i := 0; i < 1000; i++ {
		cardNumber := GenerateRandomCreditCard()
		fmt.Printf("credit card= %s. Is valid=%v\n", cardNumber, IsValidLuhn(cardNumber))
	}
}

func GenerateRandomCreditCard() string {

	cardNumber := make([]byte, 19)

	for i := 0; i < 19; i++ {
		if i == 4 || i == 9 || i == 14 {
			cardNumber[i] = ' '
		} else {
			cardNumber[i] = byte(rand.Intn(10) + '0')
		}
	}

	return string(cardNumber)
}
