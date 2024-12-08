package utils

import (
	"time"
	"golang.org/x/exp/rand"
)

func GenerateProductCodeValue(existingCodes map[string]bool) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for {
		code := ""
		for i := 1; i <= 7; i++ {
			code += string(alphanum[rand.Intn(len(alphanum))])
		}

		if !existingCodes[code] {
			existingCodes[code] = true
			return code
		}
	}
}
