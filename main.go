package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tg123/go-htpasswd"
	"golang.org/x/crypto/ssh/terminal"
)

//go:generate go run cmd/gen/gen.go

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your username: ")
	user, _ := reader.ReadString('\n')
	// convert CRLF to LF
	user = strings.Replace(user, "\n", "", -1)
	// blackmagic for windows
	user = strings.Replace(user, "\r\n", "", -1)

	fmt.Println("password!")
	password, _ := terminal.ReadPassword(0)

	myauth, err := htpasswd.NewFromReader(strings.NewReader(htpasswdData), htpasswd.DefaultSystems, nil)
	if err != nil {
		fmt.Println(err)
	}

	ok := myauth.Match(user, string(password))

	fmt.Println(ok)

}
