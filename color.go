package xpretty

import "github.com/fatih/color"

var (
	RedPrintf    = color.New(color.FgRed, color.Bold).PrintfFunc()
	CyanPrintf   = color.New(color.FgCyan, color.Bold).PrintfFunc()
	YellowPrintf = color.New(color.FgYellow, color.Bold).PrintfFunc()
	GreenPrintf  = color.New(color.FgGreen, color.Bold).PrintfFunc()
)

var (
	RedPrintln    = color.New(color.FgRed, color.Bold).PrintlnFunc()
	CyanPrintln   = color.New(color.FgCyan, color.Bold).PrintlnFunc()
	YellowPrintln = color.New(color.FgYellow, color.Bold).PrintlnFunc()
	GreenPrintln  = color.New(color.FgGreen, color.Bold).PrintlnFunc()
)

var Warn = color.New(color.FgYellow, color.BgBlack, color.Bold, color.Italic)

var (
	Red     = color.New(color.FgRed, color.Bold).SprintFunc()
	Redf    = color.New(color.FgRed, color.Bold).SprintfFunc()
	Green   = color.New(color.FgGreen, color.Bold).SprintFunc()
	Greenf  = color.New(color.FgGreen, color.Bold).SprintfFunc()
	White   = color.New(color.FgHiWhite, color.Bold).SprintFunc()
	Whitef  = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
	Yellow  = color.New(color.FgYellow, color.Bold).SprintFunc()
	Yellowf = color.New(color.FgYellow, color.Bold).SprintfFunc()
	Cyan    = color.New(color.FgCyan, color.Bold).SprintFunc()
	Cyanf   = color.New(color.FgCyan, color.Bold).SprintfFunc()
)

var Warnf = color.New(color.FgYellow, color.BgBlack, color.Bold, color.Italic).SprintfFunc()

// Redfu means red foreground with bold underline
var (
	Redfu = color.New(color.FgRed, color.Bold, color.Underline).SprintfFunc()
	Redfc = color.New(color.FgRed, color.Bold, color.CrossedOut).SprintfFunc()
)

var (
	Faint = color.New(color.Faint, color.FgWhite).SprintFunc()
	Info  = color.New(color.FgYellow, color.BgBlack, color.Bold).SprintFunc()
)

func SetNoColor(b bool) {
	color.NoColor = b
}
