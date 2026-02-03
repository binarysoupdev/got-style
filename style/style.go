// Provides semantic types and functions for styling terminal outputs using ANSI escape sequences.
package style

import (
	"fmt"
	"strings"
)

// The Style type is a slice of strings representing the ANSI style modes.
type Style []string

// ANSI style mode constants.
const (
	BOLD      = "1"
	DIM       = "2"
	UNDERLINE = "4"

	BLACK   = "30"
	WHITE   = "37"
	RED     = "31"
	GREEN   = "32"
	BLUE    = "34"
	YELLOW  = "33"
	CYAN    = "36"
	MAGENTA = "35"
)

// Creates a new style using the provided modes.
// Please use the constants instead of plain strings.
func New(mod ...string) Style {
	return mod
}

// Returns whether the style is empty (has no modes).
func (sty Style) IsEmpty() bool {
	return len(sty) == 0
}

// Format the provided string using the style.
// Return as a new string.
func (sty Style) Format(s string) string {
	if sty.IsEmpty() {
		return s
	}

	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", strings.Join(sty, ";"), s)
}

// First format using fmt package formatting, then format using the style.
// Return as a new string.
func (sty Style) FormatF(format string, a ...any) string {
	return sty.Format(fmt.Sprintf(format, a...))
}

// Print the provided string to the console using the style.
func (sty Style) Print(s string) {
	fmt.Print(sty.Format(s))
}

// Print the provided string (with new line) to the console using the style.
func (sty Style) Println(s string) {
	fmt.Println(sty.Format(s))
}

// First format using fmt package formatting, then print to the console using the style.
func (sty Style) PrintF(format string, a ...any) {
	fmt.Print(sty.FormatF(format, a...))
}
