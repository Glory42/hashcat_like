package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func main() {
	targetHash := "f4dc5569d42fa90ae423e0930e2e3555"
	wordlistPath := `C:\Users\Glory\wordlists\rockyou.txt`

	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Printf("Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hashedPassword == targetHash {
			fmt.Printf("Password found: %s\n", password)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist: %v\n", err)
		return
	}

	fmt.Println("Password not found in wordlist.")
}
