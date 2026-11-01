// Package greet is a tiny synthetic package for RiskRails runner runs.
package greet

import "strings"

// Hello returns a greeting for name.
func Hello(name string) string {
	return "Hello, " + name
}

// Shout returns the greeting for name in upper case, followed by "!".
func Shout(name string) string {
	return strings.ToUpper(Hello(name)) + "!"
}
