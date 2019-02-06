// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns a string, giving one to the name parameter and one to yourself
// If name is blank, "you" is the default
func ShareWith(name string) string {
	if name == "" {
    name = "you"
  }

	return fmt.Sprintf("One for %s, one for me.", name)
}
