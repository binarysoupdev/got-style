// Provides semantic types and functions for styling terminal outputs using ANSI escape sequences.
package style

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
