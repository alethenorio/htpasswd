package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"log"
)

var username *string = flag.StringP("username", "u", "", "Username to be used in .htpasswd")
var passwd *string = flag.StringP("passwd", "p", "", "Password to be used in .htpasswd")
var algo *string = flag.StringP("algo", "a", "MD5", "Algorithm to use for hashing passwords. Default: SHA1")

func main() {

	flag.Parse()

	if len(*username) == 0 {
		log.Fatal("Username cannot be empty")
	}

	if len(*passwd) == 0 {
		log.Fatal("Password cannot be empty")
	}

	hash, err := NewHash(*algo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash.FormattedAuth(*username, *passwd))

}
