package netease

import (
	"bufio"
	"fmt"
	"github.com/winterssy/music-get/config"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
	"time"
)

func isAuthenticated() bool {
	for _, i := range config.M.Cookies {
		if strings.ToUpper(i.Name) == "MUSIC_U" && i.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

func login() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Phone Number: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	password := string(bytePassword)

	req := NewLoginRequest(strings.TrimSpace(phone), strings.TrimSpace(password))
	return req.Do()
}
