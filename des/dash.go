package des

import "fmt"

func Logo() {
	fmt.Println("Welcome to the Game Hub!")

	// ASCII art logo GOgame
	ColorSelect("cyan")
	fmt.Println(`
		_____                   _             
		/ ____|                 | |            
		| |  __  __ _ _ __ ___  | |__   ___ _ __
		| | |_ |/ _` + "`" + ` | '_ ` + "`" + ` _ \ | '_ \ / _ \ '__|
		| |__| | (_| | | | | | | | | | |  __/ |
		 \____|\__,_|_| |_| |_| |_| |_|\___|_|
	`)
	colorRset()
}

// ClearScreen clears the terminal screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// colorRset resets the terminal color
func colorRset() {
	fmt.Print("\033[0m")
}

func ColorSelect(color string) {
	
	fmt.Print("\033[0;")
	switch color {
	case "red":
		fmt.Print("31m")
	case "green":
		fmt.Print("32m")
	case "yellow":
		fmt.Print("33m")
	case "blue":
		fmt.Print("34m")
	case "cyan":
		fmt.Print("36m")
	case "magenta":
		fmt.Print("35m")
	default:
		fmt.Print("0m") // Default color
	}
}

