package hub

import (
    "GoGame/des"
    "GoGame/games"
    "fmt"
)

func MainMenu() {
    for {
        des.ClearScreen()
        des.Logo()
        fmt.Println("=== Game Hub Main Menu ===")
        fmt.Println("1. Games")
        fmt.Println("2. Settings")
        fmt.Println("3. Exit")
        fmt.Print("Enter your choice: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            GameMenu()
        case 2:
            fmt.Println("Settings coming soon...")
            pause()
        case 3:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
            pause()
        }
    }
}

func GameMenu() {
    for {
        des.ClearScreen()
        fmt.Println("=== Game Menu ===")
        fmt.Println("1. Tic Tac Toe")
        fmt.Println("2. Number Guess")
        fmt.Println("3. Python (Snake)")
        fmt.Println("4. Back to Main Menu")
        fmt.Print("Enter your choice: ")

        var choice int
        fmt.Scanln(&choice)
        switch choice {
        case 1:
            fmt.Println("Starting Tic Tac Toe...")
            games.TicTacToe()
        case 2:
            fmt.Println("Starting Number Guess...")
            games.GuessGame()
        case 3:
            fmt.Println("Starting Python (Snake)...")
            games.PythonGame()
        case 4:
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
            pause()
        }
    }
}

func pause() {
    fmt.Println("Press Enter to continue...")
    fmt.Scanln(new(string))
}

