package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hgfischer/go-otp"
)

var (
	version string = "1.0"
	gitrev  string = "aasdasd"
)

func init() {
	//version = "1.0"
}

func main() {
	totp := &otp.TOTP{Secret: strings.ToUpper(os.Getenv("GITHUB_TOTP_SECRET")), IsBase32Secret: true}
	fmt.Printf("version: %v-%v\n", version, gitrev)
	fmt.Println(totp.Get())
}
