package style

// Several common styles for ease of use.

var None = New()

var Bold = New(BOLD)             // (Bold)
var Underline = New(UNDERLINE)   // (Underline)
var Title = New(BOLD, UNDERLINE) // (Bold, Underline)

var BoldSuccess = New(BOLD, GREEN) // (Bold, Green)
var Success = New(GREEN)           // (Green)

var BoldWarning = New(BOLD, YELLOW) // (Bold, Yellow)
var Warning = New(YELLOW)           // (Yellow)

var BoldError = New(BOLD, RED) // (Bold, Red)
var Error = New(RED)           // (Red)

var BoldCreate = New(BOLD, GREEN) // (Bold, Green)
var Create = New(GREEN)           // (Green)

var BoldDelete = New(BOLD, RED) // (Bold, Red)
var Delete = New(RED)           // (Red)

var BoldInfo = New(BOLD, CYAN) // (Bold, Cyan)
var Info = New(CYAN)           // (Cyan)
