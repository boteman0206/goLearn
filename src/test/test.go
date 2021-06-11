package main

import (
	"fmt"
	symmetricCipher "github.com/legofun/symmetric-cipher"
)

func EnEHRMobile() {
	enMobile, _ := symmetricCipher.SCEncryptString("18205696989", "sfmwohnd", "des")
	reEnMobile, _ := symmetricCipher.SCDecryptString("UHiHEyWWP8BntopDLYt46Q==", "sfmwohnd", "des")

	fmt.Println(enMobile)
	fmt.Println(reEnMobile)
}

func main() {
	EnEHRMobile()
}
