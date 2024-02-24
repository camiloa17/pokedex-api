package main

import (
	"fmt"
	"syscall"
)

func commandExit() error {
	fmt.Println("Goodbye!")
	syscall.Exit(0)
	return nil
}
