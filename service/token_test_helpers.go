package service

import "os"

func GenerateToken(userId int) (string, error) {
	// Set the token secret to a known value for testing
	os.Setenv("TOKEN_SECRET", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	return generateToken(userId)
}
