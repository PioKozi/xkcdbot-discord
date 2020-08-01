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

func PrepareInput(message, prefix string) string {
	message = strings.TrimPrefix(message, prefix)
	message = strings.TrimSpace(message)
	messageWords := strings.Fields(message)
	message = strings.Join(messageWords, " ")
	message = strings.ToLower(message)
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

func UpdateStringsMap(oldmap, updatemap map[string]string) {
	for key, value := range updatemap {
		oldmap[key] = value // not checking if key is already taken, new results are more important
	}
}
