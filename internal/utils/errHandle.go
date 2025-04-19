package utils

import (
	"fmt"
	"os"
)

func CheckErr(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s: %v\n", msg, err)
		os.Exit(1)
	}
}

func Warn(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[warning] %s: %v\n", msg, err)
	}
}
