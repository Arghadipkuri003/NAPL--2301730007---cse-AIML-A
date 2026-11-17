
package myutils

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CountVowels(s string) int {
	count := 0
	for _, ch := range strings.ToLower(s) {
		if strings.ContainsRune("aeiou", ch) {
			count++
		}
	}
	return count
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is undefined for negative numbers")
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents are not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}
