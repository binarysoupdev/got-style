package style

import (
	"fmt"
	"strings"
)

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
