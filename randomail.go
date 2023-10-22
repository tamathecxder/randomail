package randomail

import (
	"fmt"
	"math/rand"
)

var domains = []string{"example.com", "test.com", "yourdomain.com"}

// GenerateRandomEmails menghasilkan alamat email acak sejumlah n.
func GenerateRandomEmails(n int) []string {
	emails := make([]string, n)
	for i := 0; i < n; i++ {
		emails[i] = GenerateRandomEmail()
	}
	return emails
}

// GenerateRandomEmail menghasilkan alamat email acak.
func GenerateRandomEmail() string {
	username := generateRandomString(8)
	domain := domains[rand.Intn(len(domains))]
	email := fmt.Sprintf("%s@%s", username, domain)
	return email
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
