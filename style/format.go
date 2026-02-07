package style

import (
	"fmt"
	"strings"
)

// Format the provided string using the style.
// Return as a new string.
func (s Style) Sprint(str string) string {
	if s.IsEmpty() {
		return str
	} else {
		return fmt.Sprintf("\x1B[%sm%s\x1B[0m", strings.Join(s, ";"), str)
	}
}

// First format using fmt package formatting, then format using the style.
// Return as a new string.
func (s Style) Sprintf(format string, a ...any) string {
	return s.Sprint(fmt.Sprintf(format, a...))
}

// Write the provided string to stdout using the style.
func (s Style) Print(str string) {
	fmt.Print(s.Sprint(str))
}

// Write the provided string (with new line) to stdout using the style.
func (s Style) Println(str string) {
	fmt.Println(s.Sprint(str))
}

// First format using fmt package formatting, then write to stdout using the style.
func (s Style) Printf(format string, a ...any) {
	fmt.Print(s.Sprintf(format, a...))
}

//============================================

// Deprecated: use Sprint instead.
func (s Style) Format(str string) string {
	return s.Sprint(str)
}

// Deprecated: use Sprintf instead.
func (s Style) FormatF(format string, a ...any) string {
	return s.Sprintf(format, a...)
}

// Deprecated: use Printf instead.
func (s Style) PrintF(format string, a ...any) {
	s.Printf(format, a...)
}
