package sdk

import "fmt"

func NewChat(address, username, _, _ string) string {
	fmt.Println("start connection")
	return "connection established"
}
