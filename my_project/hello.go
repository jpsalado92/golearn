package my_project

import (
	"fmt"
	"strings"
)

func Say(names []string)string {
	if len(names) == 0 {
		return "Hello, world!"
	}
	return fmt.Sprintf("Hello, %s.", strings.Join(names, ", ") + "!")
}
