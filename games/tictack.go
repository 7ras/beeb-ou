package games

import (
	"GoGame/des"
	"fmt"
	"math/rand"
	"time"
)

func TicTacToe() {
	des.ClearScreen()
	fmt.Println("=== Tic Tac Toe ===")

	// Board uses 9 cells; empty cells are ' '
	board := make([]rune, 9)
	for i := range board {
		board[i] = ' '
	}

	rand.Seed(time.Now().UnixNano())

	// Mode selection
	vsCPU := selectMode()

	current := 'X'
	human := 'X'
	cpu := 'O'

	for {
		des.ClearScreen()
		fmt.Println("=== Tic Tac Toe ===")
		drawBoard(board)

		if vsCPU && current == cpu {
			// CPU move
			time.Sleep(400 * time.Millisecond)
			pos := cpuMove(board, cpu, human)
			if !updateBoard(board, pos, current) {
				// Fallback: find first empty
				for i := range board {
					if board[i] == ' ' {
						updateBoard(board, i, current)
						break
					}
				}
			}
		} else {
			// Human move
			pos := promptMove(current, board)
			if !updateBoard(board, pos, current) {
				// Shouldn't happen due to prompt validation, but stay safe
				fmt.Println("Unexpected invalid move. Try again.")
				time.Sleep(800 * time.Millisecond)
				continue
			}
		}

		if checkWin(board, current) {
			des.ClearScreen()
			fmt.Println("=== Tic Tac Toe ===")
			drawBoard(board)
			if vsCPU && current == cpu {
				fmt.Println("CPU wins!")
			} else {
				fmt.Print("Player ")
				printColoredSymbol(current)
				fmt.Println(" wins!")
			}
			fmt.Println("Press Enter to continue...")
			fmt.Scanln(new(string))
			return
		}
		if checkDraw(board) {
			des.ClearScreen()
			fmt.Println("=== Tic Tac Toe ===")
			drawBoard(board)
			fmt.Println("It's a draw!")
			fmt.Println("Press Enter to continue...")
			fmt.Scanln(new(string))
			return
		}

		if current == 'X' {
			current = 'O'
		} else {
			current = 'X'
		}
	}
}

func BoardAnimation() {
	for i := 0; i < 5; i++ {
		fmt.Println("+---+---+---+")
		fmt.Println("|   |   |   |")
		fmt.Println("+---+---+---+")
		fmt.Println("|   |   |   |")
		fmt.Println("+---+---+---+")
		fmt.Println("|   |   |   |")
		fmt.Println("+---+---+---+")
		time.Sleep(400 * time.Millisecond)
		des.ClearScreen()

		fmt.Println("+---+---+---+")
		fmt.Println("| 1 | 2 | 3 |")
		fmt.Println("+---+---+---+")
		fmt.Println("| 4 | 5 | 6 |")
		fmt.Println("+---+---+---+")
		fmt.Println("| 7 | 8 | 9 |")
		fmt.Println("+---+---+---+")
		time.Sleep(400 * time.Millisecond)
		des.ClearScreen()

	}

}

func Board() {
	// empty board
	fmt.Println("+---+---+---+")
	fmt.Println("|   |   |   |")
	fmt.Println("+---+---+---+")
	fmt.Println("|   |   |   |")
	fmt.Println("+---+---+---+")
	fmt.Println("|   |   |   |")
	fmt.Println("+---+---+---+")

}

func updateBoard(board []rune, position int, player rune) bool {
	// position is zero-based index (0-8)
	if position < 0 || position >= len(board) {
		return false
	}
	if board[position] != ' ' {
		return false
	}
	board[position] = player
	return true
}

func PlayerInput(input int) {
	// take input from player
	fmt.Print("Enter your move (1-9): ")
	fmt.Scanln(&input)
	// validate input
	if input < 1 || input > 9 {
		fmt.Println("Invalid input. Please try again.")
		PlayerInput(input)
	}

}

// Helpers used by TicTacToe
func selectMode() bool {
	for {
		fmt.Println("Choose mode:")
		fmt.Println("1) Two players")
		fmt.Println("2) Vs CPU")
		fmt.Print("Enter 1 or 2: ")
		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}
		if choice == 1 {
			return false
		}
		if choice == 2 {
			return true
		}
		fmt.Println("Please enter 1 or 2.")
	}
}

func drawBoard(board []rune) {
	// Shows numbers for empty cells to guide input
	fmt.Println("+---+---+---+")
	for r := 0; r < 3; r++ {
		fmt.Print("|")
		for c := 0; c < 3; c++ {
			i := r*3 + c
			if board[i] == ' ' {
				// Display cell number (1-9) when empty
				fmt.Printf(" %d |", i+1)
			} else {
				fmt.Print(" ")
				printColoredSymbol(board[i])
				fmt.Print(" |")
			}
		}
		fmt.Println()
		fmt.Println("+---+---+---+")
	}
}

func promptMove(player rune, board []rune) int {
	for {
		var input int
		fmt.Print("Player ")
		printColoredSymbol(player)
		fmt.Print(", enter your move (1-9): ")
		if _, err := fmt.Scanln(&input); err != nil {
			// Clear bad input line if needed
			fmt.Println("Invalid input. Please enter a number 1-9.")
			continue
		}
		if input < 1 || input > 9 {
			fmt.Println("Out of range. Choose 1-9.")
			continue
		}
		idx := input - 1
		if board[idx] != ' ' {
			fmt.Println("Cell already taken. Choose another.")
			continue
		}
		return idx
	}
}

func checkWin(b []rune, p rune) bool {
	wins := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // cols
		{0, 4, 8}, {2, 4, 6}, // diags
	}
	for _, w := range wins {
		if b[w[0]] == p && b[w[1]] == p && b[w[2]] == p {
			return true
		}
	}
	return false
}

func checkDraw(b []rune) bool {
	for _, v := range b {
		if v == ' ' {
			return false
		}
	}
	return true
}

func printColoredSymbol(symbol rune) {
	switch symbol {
	case 'X':
		des.ColorSelect("red")
		fmt.Print("X")
		des.ColorSelect("reset")
	case 'O':
		des.ColorSelect("blue")
		fmt.Print("O")
		des.ColorSelect("reset")
	default:
		fmt.Printf("%c", symbol)
	}
}

// Basic CPU move: win, block, center, corners, sides
func cpuMove(board []rune, cpu, human rune) int {
	// Try winning move
	for i := 0; i < 9; i++ {
		if board[i] == ' ' {
			board[i] = cpu
			if checkWin(board, cpu) {
				board[i] = ' '
				return i
			}
			board[i] = ' '
		}
	}
	// Try to block human winning move
	for i := 0; i < 9; i++ {
		if board[i] == ' ' {
			board[i] = human
			if checkWin(board, human) {
				board[i] = ' '
				return i
			}
			board[i] = ' '
		}
	}
	// Center
	if board[4] == ' ' {
		return 4
	}
	// Corners
	corners := []int{0, 2, 6, 8}
	var availCorners []int
	for _, i := range corners {
		if board[i] == ' ' {
			availCorners = append(availCorners, i)
		}
	}
	if len(availCorners) > 0 {
		return availCorners[rand.Intn(len(availCorners))]
	}
	// Sides
	sides := []int{1, 3, 5, 7}
	var availSides []int
	for _, i := range sides {
		if board[i] == ' ' {
			availSides = append(availSides, i)
		}
	}
	if len(availSides) > 0 {
		return availSides[rand.Intn(len(availSides))]
	}
	// Fallback: first empty
	for i := 0; i < 9; i++ {
		if board[i] == ' ' {
			return i
		}
	}
	return 0
}
