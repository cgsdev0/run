package color

import "github.com/charmbracelet/lipgloss"

// https://ethanschoonover.com/solarized/#the-values
var (
	Yellow  = lipgloss.Color("#B58900")
	Orange  = lipgloss.Color("#CB4B16")
	Red     = lipgloss.Color("#DC322F")
	Magenta = lipgloss.Color("#D33682")
	Violet  = lipgloss.Color("#6C71C4")
	Blue    = lipgloss.Color("#268BD2")
	Cyan    = lipgloss.Color("#2AA198")
	Green   = lipgloss.Color("#859900")

	XXXLight = lipgloss.Color("7") // base3
	XXLight  = lipgloss.Color("7") // base2
	XLight   = lipgloss.Color("7") // base1
	Light    = lipgloss.Color("7") // base0
	Dark     = lipgloss.Color("7") // base00
	XDark    = lipgloss.Color("7") // base01
	XXDark   = lipgloss.Color("7") // base02
	XXXDark  = lipgloss.Color("7") // base03
)
