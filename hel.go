package pkg

import "fmt"

// Hello says hello.
func Hel() error {
	fmt.Println("Hello go mod!")
	return nil
}

// Bye says bye.
func By() error {
	fmt.Println("Bye go mod!")
	return nil
}
