package hello

import "fmt"

// Say returns a greeting message
func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
