package style

import (
	"fmt"
	"strings"
)

// Format as a string, first using fmt, then using the style.
func (s Style) Sprint(a ...any) string {
	str := fmt.Sprint(a...)

	if s.IsEmpty() {
		return str
	} else {
		return fmt.Sprintf("\x1B[%sm%s\x1B[0m", strings.Join(s, ";"), str)
	}
}

// Format as a string, first using fmt, then using the style.
func (s Style) Sprintf(format string, a ...any) string {
	return s.Sprint(fmt.Sprintf(format, a...))
}

// Format to stdout, first using fmt, then using the style.
func (s Style) Print(a ...any) {
	fmt.Print(s.Sprint(a...))
}

// Format to stdout (with new line), first using fmt, then using the style.
func (s Style) Println(a ...any) {
	fmt.Println(s.Sprint(a...))
}

// Format to stdout, first using fmt, then using the style.
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
