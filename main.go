package test1

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome! new version", name)
	return message
}
