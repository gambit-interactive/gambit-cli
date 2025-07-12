package utils

import "github.com/fatih/color"

// PrintLogo displays the Gambit logo in the terminal with colors.
func PrintLogo() {
	cyan := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	// ASCII Art for "Gambit"
	logo := `
██████╗  █████╗ ███╗   ███╗██████╗ ██╗████████╗
 ██╔════╝ ██╔══██╗████╗ ████║██╔══██╗██║╚══██╔══╝
 ██║  ███╗███████║██╔████╔██║██████╔╝██║   ██║   
 ██║   ██║██╔══██║██║╚██╔╝██║██╔══██╗██║   ██║   
 ╚██████╔╝██║  ██║██║ ╚═╝ ██║██████╔╝██║   ██║   
  ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚═════╝ ╚═╝   ╚═╝
`
	println(cyan(bold(logo)))
}
