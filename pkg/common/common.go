package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// useful is it's going pretty bad or is malicious around some time, somehow
func LogMessage(s string) {

	fmt.Println("------------------------")
	fmt.Println(time.Now())
	fmt.Println(s)
}

func CleanInput(message, prefix string) string {

	message = strings.TrimPrefix(message, prefix)
	message = strings.TrimSpace(message)
	return message
}

func ValidID(id string) bool {

	i, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	if i < 0 {
		return false
	}
	return true
}
