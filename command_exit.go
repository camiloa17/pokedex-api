package main

import (
	"fmt"
	"syscall"
)

func commandExit(conf *apiConfig, args ...string) error {
	fmt.Println("Goodbye!")
	syscall.Exit(0)
	return nil
}
