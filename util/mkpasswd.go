// Copyright 2015 Apcera Inc. All rights reserved.
// +build ignore

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func usage() {
	log.Fatalf("Usage: mkpasswd -p <stdin password> \n")
}

const (
	PasswordLength = 22
	Cost           = 11
)

func main() {
	var pw = flag.Bool("p", false, "Input password via stdin")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	var password string

	if *pw {
		fmt.Printf("Enter Password: ")
		bytePassword, _ := terminal.ReadPassword(0)
		fmt.Printf("\nReenter Password: ")
		bytePassword2, _ := terminal.ReadPassword(0)
		if !bytes.Equal(bytePassword, bytePassword2) {
			log.Fatalf("Error, passwords do not match\n")
		}
		password = string(bytePassword)
		fmt.Printf("\n")
	} else {
		password = genPassword()
		fmt.Printf("pass: %s\n", password)
	}

	cb, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	if err != nil {
		log.Fatalf("Error producing bcrypt hash: %v\n", err)
	}
	fmt.Printf("bcrypt hash: %s\n", cb)
}

func genPassword() string {
	var ch = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@$#%^&*()")
	b := make([]byte, PasswordLength)
	max := big.NewInt(int64(len(ch)))
	for i := range b {
		ri, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Fatalf("Error producing random integer: %v\n", err)
		}
		b[i] = ch[int(ri.Int64())]
	}
	return string(b)
}
