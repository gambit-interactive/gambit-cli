package utils

import "github.com/fatih/color"

// PrintLogo displays the Gambit logo in the terminal with colors.
func PrintLogo() {
	cyan := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	// ASCII Art for "Gambit"
	logo := `
  ________              .__    .__  __ 
 /  _____/  ____   ____ |  |__ |__|/  |_
/   \  ___ /  _ \ /  _ \|  |  \|  \   __\
\    \_\  (  <_> |  <_> )   Y  \  ||  |
 \______  /\____/ \____/|___|  /__||__|
        \/                  \/             
`
	println(cyan(bold(logo)))
}
