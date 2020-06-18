//This is hello package1
//This is hello package2
package pkg

import "fmt"

// DisableModifiers will disable the modifier syntax
var DisableModifiers = false

// Result represents a json value that is returned from Get().
type Result struct {
	// Type is the json type
	Type Type
	// Raw is the raw json
	Raw string
	// Str is the json string
	Str string
	// Num is the json number
	Num float64
	// Index of raw value in original json, zero means index unknown
	Index int
}

// Type is Result type
type Type int

// Result represents a json value that is returned from Get().
type Result1 struct {
	// Type is the json type
	Type1 Type
	// Raw is the raw json
	Raw1 string
	// Str is the json string
	Str1 string
	// Num is the json number
	Num1 float64
	// Index of raw value in original json, zero means index unknown
	Index1 int
}
// Hello function for say hello.
func Hel(t Type) error {
	//print hello go mod message
	fmt.Println("Hello go mod!")
	//return nil
	return nil
}

// Bye function for say bye.
func By() error {
	//Print bye go mod message
	fmt.Println("Bye go mod!")
	//return nil
	return nil
}
