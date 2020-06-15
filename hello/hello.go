package hello

import "fmt"

// Hello says hello.
func HelloA() error {
	fmt.Println("Hello go mod!")
	return nil
}

// Bye says bye.
func ByeA() error {
	fmt.Println("Bye go mod!")
	return nil
}
