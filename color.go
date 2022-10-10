package xpretty

import "github.com/fatih/color"

var RedPrintf = color.New(color.FgRed, color.Bold).PrintfFunc()
var CyanPrintf = color.New(color.FgCyan, color.Bold).PrintfFunc()
var YellowPrintf = color.New(color.FgYellow, color.Bold).PrintfFunc()
var GreenPrintf = color.New(color.FgGreen, color.Bold).PrintfFunc()

var Red = color.New(color.FgRed, color.Bold).SprintFunc()
var Redf = color.New(color.FgRed, color.Bold).SprintfFunc()
var Green = color.New(color.FgGreen, color.Bold).SprintFunc()
var Greenf = color.New(color.FgGreen, color.Bold).SprintfFunc()
var White = color.New(color.FgHiWhite, color.Bold).SprintFunc()
var Whitef = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
var Yellow = color.New(color.FgYellow, color.Bold).SprintFunc()
var Yellowf = color.New(color.FgYellow, color.Bold).SprintfFunc()

var Warnf = color.New(color.FgYellow, color.BgBlack, color.Bold, color.Italic).SprintfFunc()

// Redfu means red foreground with bold underline
var Redfu = color.New(color.FgRed, color.Bold, color.Underline).SprintfFunc()
var Redfc = color.New(color.FgRed, color.Bold, color.CrossedOut).SprintfFunc()

var Faint = color.New(color.Faint, color.FgWhite).SprintFunc()
var Info = color.New(color.FgYellow, color.BgBlack, color.Bold).SprintFunc()

func ToggleColor(b bool) {
	color.NoColor = !b
}
