package random

import (
	"math/rand"
)

func RandomString(length int) string {
	symbols := make([]rune, 0)
	func() {
		for i := range 128 + 1 {
			if (48 <= i && i <= 57) || (65 <= i && i <= 90) || (97 <= i && i <= 122) {
				symbols = append(symbols, rune(i))
			}
		}
	}()

	result := make([]rune, 0, length)
	for range length {
		result = append(result, symbols[rand.Intn(len(symbols))])
	}
	return string(result)
}
