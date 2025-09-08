package games

import (
    "GoGame/des"
    "fmt"
    "math/rand"
    "time"
)

// GuessGame is a simple number guessing game
func GuessGame() {
    des.ClearScreen()
    fmt.Println("=== Number Guess ===")
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1 // 1-100

    attempts := 0
    for {
        attempts++
        var g int
        fmt.Print("Enter a number (1-100): ")
        if _, err := fmt.Scanln(&g); err != nil {
            fmt.Println("Please enter a valid number.")
            continue
        }
        if g < 1 || g > 100 {
            fmt.Println("Out of range. Choose 1-100.")
            continue
        }
        if g < target {
            des.ColorSelect("blue")
            fmt.Println("Too low!")
            des.ColorSelect("reset")
        } else if g > target {
            des.ColorSelect("yellow")
            fmt.Println("Too high!")
            des.ColorSelect("reset")
        } else {
            des.ColorSelect("green")
            fmt.Printf("Correct! You guessed it in %d attempts.\n", attempts)
            des.ColorSelect("reset")
            fmt.Println("Press Enter to continue...")
            fmt.Scanln(new(string))
            return
        }
    }
}

