package util

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToBool converts a string to a boolean value
func StringToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true", "t", "yes", "y", "1":
		return true, nil
	case "false", "f", "no", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("cannot convert %q to boolean", s)
	}
}

func StringToBoolAsInt(value string) int {
	if value == "" || value == "<nil>" {
		return -1
	}
	isDeleted, err := strconv.ParseBool(value)
	if err != nil {
		// Handle error, maybe return a default value or panic
		fmt.Println("err", err)
		return -1
	}
	if isDeleted {
		return 1
	}
	return 0
}

func UintToBool(u uint) bool {
	return u != 0
}
