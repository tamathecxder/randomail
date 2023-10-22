package randomail

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

var domains = []string{"example.com", "test.com", "yourdomain.com"}

// GenerateRandomEmailWithCustomDomain generates a random email address with a customised domain.
func GenerateRandomEmailWithCustomDomain(customDomain string) (string, error) {
	if !strings.Contains(customDomain, ".") {
		return "", errors.New("custom domain must contain a dot (.)")
	}

	username := generateRandomString(8)
	email := fmt.Sprintf("%s@%s", username, customDomain)

	return email, nil
}

// `GenerateRandomEmails` generates random email addresses equal to n number of arguments.
func GenerateRandomEmails(n int) []string {
	emails := make([]string, n)
	for i := 0; i < n; i++ {
		emails[i] = GenerateRandomEmail()
	}
	return emails
}

// `GenerateRandomEmail` generates a random email address.
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
