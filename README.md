# GoGame Hub 🎮

A terminal-based game collection written in Go, featuring classic games with an intuitive menu system.

## Features

- **Interactive Menu System**: Navigate through games with a user-friendly console interface
- **Multiple Games**: Collection of classic games implemented in Go
- **Colorful Interface**: Terminal colors and ASCII art for enhanced user experience
- **Cross-platform**: Works on macOS, Linux, and Windows

## Games Included

### 🎯 Tic Tac Toe
- Classic 3x3 grid game
- Play against CPU or another player
- Smart AI opponent with strategic moves

### 🔢 Number Guess
- Guess a random number between 1-100
- Color-coded feedback (too high/too low)
- Tracks number of attempts

### 🐍 Python (Snake Game)
- Terminal-based Snake game
- Real-time keyboard controls (WASD)
- Score tracking and collision detection

## Installation

### Prerequisites
- Go 1.22.0 or higher

### Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/7ras/beeb-ou.git
   cd beeb-ou
   ```

2. **Build and run:**
   ```bash
   go run main.go
   ```

   Or build an executable:
   ```bash
   go build -o gogame
   ./gogame
   ```

## Project Structure

```
beeb-ou/
├── main.go          # Entry point
├── go.mod           # Go module file
├── hub/             # Menu and navigation
│   ├── menu.go      # Main menu and game selection
│   └── settings.go  # Settings (coming soon)
├── games/           # Game implementations
│   ├── tictack.go   # Tic Tac Toe game
│   ├── guess.go     # Number guessing game
│   └── python.go    # Snake game
└── des/             # Display utilities
    └── dash.go      # Colors, logo, and screen management
```

## Usage

1. **Start the application:**
   ```bash
   go run main.go
   ```

2. **Navigate the menu:**
   - Use number keys to select options
   - Follow on-screen prompts
   - Press Enter to confirm selections

3. **Game Controls:**
   - **Tic Tac Toe**: Enter position numbers (1-9)
   - **Number Guess**: Type your guess and press Enter
   - **Snake**: Use WASD keys for movement

## Development

### Adding New Games

1. Create a new file in the `games/` directory
2. Implement your game function with the signature `func GameName()`
3. Add the game to the menu in `hub/menu.go`
4. Import the function in the switch statement

### Code Style

- Follow Go conventions and formatting
- Use `gofmt` for code formatting
- Add comments for exported functions
- Keep functions focused and modular

## Technical Details

- **Language**: Go 1.22.0
- **Dependencies**: Standard library only
- **Platform**: Cross-platform terminal application
- **Input**: Keyboard input via standard input
- **Output**: ANSI color terminal output

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-game`)
3. Commit your changes (`git commit -am 'Add new game'`)
4. Push to the branch (`git push origin feature/new-game`)
5. Create a Pull Request

## License

This project is open source. Feel free to use, modify, and distribute.

## Roadmap

- [ ] Settings menu implementation
- [ ] High score tracking
- [ ] More games (Hangman, Rock Paper Scissors, etc.)
- [ ] Sound effects (terminal beeps)
- [ ] Difficulty levels for existing games
- [ ] Save/load game state

## Author

Created by [7ras](https://github.com/7ras)

---

*Enjoy gaming in your terminal! 🎮*