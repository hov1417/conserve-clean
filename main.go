package main

import (
	"github.com/hov1417/conserve-clean/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
