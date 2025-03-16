package hello

import (
		// "fmt"
		"strings"
)

// Say returns a greeting message
func Say(names []string) string {
	if len(names) == 0{
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!" 
}
