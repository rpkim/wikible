package util

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func GetCredentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	fmt.Println("")

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

func GetPromptString(question string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(question)

	answer, _ := reader.ReadString('\n')

	return strings.TrimSpace(answer)
}