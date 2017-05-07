package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
)

var username *string = flag.StringP("username", "u", "", "Username to be used in .htpasswd")
var passwd *string = flag.StringP("passwd", "p", "", "Password to be used in .htpasswd")
var algo *string = flag.StringP("algo", "a", "MD5", "Algorithm to use for hashing passwords. Default: SHA1")
var versionf *bool = flag.BoolP("version", "v", false, "Prints information about the tool")

var (
	version   = "N/A"
	buildTime = "N/A"
	commitId  = "N/A"
)

func main() {

	flag.Parse()
	if *versionf {
		printVersion()
		os.Exit(0)
	}

	if len(*username) == 0 {
		fmt.Fprint(os.Stderr, "Username cannot be empty\n")
		os.Exit(1)
	}

	if len(*passwd) == 0 {
		fmt.Fprint(os.Stderr, "Password cannot be empty\n")
		os.Exit(1)
	}

	hash, err := NewHash(*algo)

	if err != nil {
		panic(err)
	}

	fmt.Println(hash.FormattedAuth(*username, *passwd))
}

func printVersion() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("BuildTime: %s\n", buildTime)
	fmt.Printf("CommitId: %s\n", commitId)
}
