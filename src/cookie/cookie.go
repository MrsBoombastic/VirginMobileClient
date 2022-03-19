package cookie

import (
	"log"
	"os"
)

func SaveCookie(s string) error {
	cookie := []byte(s)
	err := os.WriteFile("cookie.txt", cookie, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadCookie() (string, error) {
	dat, err := os.ReadFile("cookie.txt")
	if err != nil {
		log.Println("Error reading cookie.txt. If this file does not exist, please create new one!")
		return "", err
	}
	return string(dat), nil
}
