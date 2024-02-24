package main

import (
	"fmt"
	"syscall"
)

func commandExit(conf *apiConfig) error {
	fmt.Println("Goodbye!")
	syscall.Exit(0)
	return nil
}
