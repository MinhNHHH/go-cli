package main

import (
	"cli-list/pkg/fetch"
	"os"
)

func main() {
	var cmd []string
	if len(os.Args) < 2 {
		cmd = []string{""}
	} else {
		cmd = os.Args[1:]
	}
	fetch.HandleClient(cmd)
}
