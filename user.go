package main

import (
	"fmt"
	"os"
	"os/user"
)

func getAuthor() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s@%s", u.Username, hostname)
}
