package games

import (
	"GoGame/des"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type point struct{ r, c int }

// Terminal control structures for raw input
type termios struct {
	Iflag  uint64
	Oflag  uint64
	Cflag  uint64
	Lflag  uint64
	Cc     [20]uint8
	Ispeed uint64
	Ospeed uint64
}

// enableRawMode puts terminal in raw mode
func enableRawMode() *termios {
	var original termios
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TIOCGETA, uintptr(unsafe.Pointer(&original)))

	raw := original
	raw.Lflag &^= syscall.ECHO | syscall.ICANON
	raw.Cc[syscall.VMIN] = 1
	raw.Cc[syscall.VTIME] = 0

	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TIOCSETA, uintptr(unsafe.Pointer(&raw)))
	return &original
}

// disableRawMode restores original terminal mode
func disableRawMode(original *termios) {
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TIOCSETA, uintptr(unsafe.Pointer(original)))
}

// readChar reads a single character without Enter
func readChar() (rune, error) {
	var buf [1]byte
	n, err := os.Stdin.Read(buf[:])
	if err != nil || n == 0 {
		return 0, err
	}
	return rune(buf[0]), nil
}

func PythonGame() {
	des.ClearScreen()
	fmt.Println("=== Python (Snake) ===")

	// Enable raw mode for single-key input
	original := enableRawMode()
	defer disableRawMode(original)

	rows, cols := 12, 20
	// initial snake of length 3 in center
	start := point{rows / 2, cols / 2}
	snake := []point{{start.r, start.c - 1}, {start.r, start.c}, {start.r, start.c + 1}}
	dir := point{0, 1} // moving right

	food := spawnFood(rows, cols, snake)
	score := 0

	for {
		des.ClearScreen()
		fmt.Println("=== Python (Snake) ===")
		fmt.Printf("Score: %d\n", score)
		drawSnakeBoard(rows, cols, snake, food)
		fmt.Print("Move [W/A/S/D], Q to quit (auto-move after 1 sec): ")

		// Create a channel to receive input
		inputChan := make(chan rune, 1)

		// Start a goroutine to read single character input
		go func() {
			ch, err := readChar()
			if err == nil {
				inputChan <- ch
			}
		}()

		// Wait for input or timeout after 1 second
		var inputReceived bool
		select {
		case ch := <-inputChan:
			// Input received
			inputReceived = true
			key := strings.ToLower(string(ch))
			switch key {
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
		case <-time.After(1 * time.Second):
			// Timeout - continue with current direction
			inputReceived = false
		}

		_ = inputReceived // Mark as used

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
