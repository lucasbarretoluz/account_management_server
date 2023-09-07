package util

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]

	}
	return string(b)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCategory() string {
	categories := []string{"food", "travel", "rent", "salary"}

	n := len(categories)
	return categories[rand.Intn(n)]
}

func RandomDescription() string {
	return RandomString(20)
}

func RandomIsExpense() bool {
	return rand.Intn(2) == 1
}

func RandomUserName() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "BRL"}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomUserEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
