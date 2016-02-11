package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codeskyblue/go-sh"
)

func main() {
	cmd := "kubectl"
	if !exists(cmd) {
		fmt.Println("kubectl does not exist")
		os.Exit(1)
	}

	fmt.Println("kubectl exists!")
}

func exists(cmd string) bool {
	return sh.Test("x", which(cmd))
}

func which(cmd string) string {
	which, err := sh.Command("which", cmd).Output()
	if err != nil {
		fmt.Println("An error occured when trying to find " + cmd)
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Trim(string(which), "\n")
}
