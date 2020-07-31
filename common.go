package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// useful is it's going pretty bad or is malicious around some time, somehow
func logMessage(s string) {

	fmt.Println("------------------------")
	fmt.Println(time.Now())
	fmt.Println(s)
}

func cleanInput(message, prefix string) string {

	message = strings.TrimPrefix(message, prefix)
	message = strings.TrimSpace(message)
	return message
}

func validID(id string) bool {

	i, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	if i < 0 {
		return false
	}
	return true
}
