package main

import "log"

func errorHandler(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
