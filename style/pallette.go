package style

// Several common styles for ease of use.

var None = New()

var Bold = New(BOLD)
var Underline = New(UNDERLINE)
var Title = New(BOLD, UNDERLINE)

var BoldSuccess = New(BOLD, GREEN)
var Success = New(GREEN)

var BoldWarning = New(BOLD, YELLOW)
var Warning = New(YELLOW)

var BoldError = New(BOLD, RED)
var Error = New(RED)

var BoldCreate = New(BOLD, GREEN)
var Create = New(GREEN)

var BoldDelete = New(BOLD, RED)
var Delete = New(RED)

var BoldInfo = New(BOLD, CYAN)
var Info = New(CYAN)
