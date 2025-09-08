package games

import (
	"GoGame/des"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type point struct{ r, c int }

// PythonGame is a simple turn-based terminal Snake clone
func PythonGame() {
	des.ClearScreen()
	fmt.Println("=== Python (Snake) ===")

	rows, cols := 12, 20
	// initial snake of length 3 in center
	start := point{rows / 2, cols / 2}
	snake := []point{{start.r, start.c - 1}, {start.r, start.c}, {start.r, start.c + 1}}
	dir := point{0, 1} // moving right

	food := spawnFood(rows, cols, snake)
	score := 0

	in := bufio.NewReader(os.Stdin)

	for {
		des.ClearScreen()
		fmt.Println("=== Python (Snake) ===")
		fmt.Printf("Score: %d\n", score)
		drawSnakeBoard(rows, cols, snake, food)
		fmt.Print("Move [W/A/S/D], Q to quit (auto-move after 1 sec): ")

		// Create a channel to receive input
		inputChan := make(chan string, 1)

		// Start a goroutine to read input
		go func() {
			line, _ := in.ReadString('\n')
			inputChan <- strings.TrimSpace(line)
		}()

		// Wait for input or timeout after 1 second
		var line string
		select {
		case line = <-inputChan:
			// Input received
		case <-time.After(1 * time.Second):
			// Timeout - continue with current direction
			line = ""
		}

		if line != "" {
			ch := strings.ToLower(string(line[0]))
			switch ch {
			case "w":
				if dir.r != 1 { // prevent instant reverse
					dir = point{-1, 0}
				}
			case "a":
				if dir.c != 1 {
					dir = point{0, -1}
				}
			case "s":
				if dir.r != -1 {
					dir = point{1, 0}
				}
			case "d":
				if dir.c != -1 {
					dir = point{0, 1}
				}
			case "q":
				return
			}
		}

		// compute next head
		head := snake[len(snake)-1]
		next := point{head.r + dir.r, head.c + dir.c}

		// check collision with walls
		if next.r < 0 || next.r >= rows || next.c < 0 || next.c >= cols {
			gameOverSnake(score)
			return
		}
		// check collision with self
		for _, p := range snake {
			if p == next {
				gameOverSnake(score)
				return
			}
		}

		// move snake
		snake = append(snake, next)
		if next == food {
			score += 1
			food = spawnFood(rows, cols, snake)
		} else {
			// remove tail
			snake = snake[1:]
		}
	}
}

func drawSnakeBoard(rows, cols int, snake []point, food point) {
	// prepare occupancy map
	occ := make(map[point]bool, len(snake))
	for _, p := range snake {
		occ[p] = true
	}

	// top border
	fmt.Print("+")
	for i := 0; i < cols; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for r := 0; r < rows; r++ {
		fmt.Print("|")
		for c := 0; c < cols; c++ {
			p := point{r, c}
			if p == food {
				des.ColorSelect("magenta")
				fmt.Print("*")
				des.ColorSelect("reset")
			} else if occ[p] {
				des.ColorSelect("green")
				fmt.Print("O")
				des.ColorSelect("reset")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}

	// bottom border
	fmt.Print("+")
	for i := 0; i < cols; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func spawnFood(rows, cols int, snake []point) point {
	occ := make(map[point]bool, len(snake))
	for _, p := range snake {
		occ[p] = true
	}
	for {
		p := point{rand.Intn(rows), rand.Intn(cols)}
		if !occ[p] {
			return p
		}
	}
}

func gameOverSnake(score int) {
	des.ClearScreen()
	des.ColorSelect("red")
	fmt.Println("Game Over!")
	des.ColorSelect("reset")
	fmt.Printf("Final score: %d\n", score)
	fmt.Println("Press Enter to continue...")
	fmt.Scanln(new(string))
}
