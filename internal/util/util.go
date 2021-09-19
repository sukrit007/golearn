package util

import "fmt"

func DoInternal() string {
	return "test"
}

func Hello(name string) string {
	retValue := fmt.Sprintf("Hello %v", name)
	return retValue
}
