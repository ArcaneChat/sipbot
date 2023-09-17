package main

import (
	"encoding/base64"

	"github.com/sethvargo/go-password/password"
)

// convert email address to SIP user name
func addr2user(addr string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(addr))
}

// generate a random password
func genPassword() string {
	pass, err := password.Generate(30, 5, 5, false, true)
	if err != nil {
		panic(err)
	}
	return pass
}
