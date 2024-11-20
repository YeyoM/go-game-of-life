# Terminal-based Game of Life

A Go implementation of Conway's Game of Life, rendered directly in your terminal. This project brings the fascinating world of cellular automata to life through a simple yet powerful command-line interface.

## About Conway's Game of Life

Conway's Game of Life is a cellular automaton devised by mathematician John Conway in 1970. It's not a traditional game - rather, it's a zero-player game that evolves based on its initial state. The game demonstrates how complex patterns can emerge from simple rules.

### Rules

The universe of the Game of Life is an infinite two-dimensional grid of cells, where each cell is either alive or dead. The game evolves in discrete time steps, where each cell's fate is determined by its neighbors. Each cell interacts with its eight neighbors (horizontal, vertical, and diagonal) according to these rules:

1. Any live cell with fewer than two live neighbors dies (underpopulation)
2. Any live cell with two or three live neighbors lives on to the next generation
3. Any live cell with more than three live neighbors dies (overpopulation)
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction)

## Features

- Terminal-based visualization
- Custom pattern loading from files
- Responsive to terminal resizing
- Multiple initialization options (random, pre-defined patterns, custom patterns)

## Installation

```bash
# Clone the repository
git clone https://github.com/YeyoMoreno/go-game-of-life.git

# Navigate to the project directory
cd go-game-of-life

# Install dependencies
go mod download
```

## Dependencies

- [tcell/v2](https://github.com/gdamore/tcell) - Terminal handling
- [keybd_event](https://github.com/micmonay/keybd_event) - Keyboard event handling

## Usage

```bash
# Run with a custom pattern file
go run main.go pattern.txt
```

### Controls

- `Esc` or `Ctrl+C` - Exit the program

## Pattern Files

You can create custom patterns using text files. The format is:

```
.O.  # O represents a live cell
O.O  # . represents a dead cell
.O.
```

_Do not use spaces between cells._

## Understanding Cellular Automata

Cellular automata are discrete models studied in computability theory and mathematics. They consist of a regular grid of cells, each in one of a finite number of states. The grid can be in any finite number of dimensions. For each cell, a set of cells called its neighborhood is defined relative to that cell.

Some famous patterns in Conway's Game of Life include:

- **Still lifes**: Patterns that don't change
- **Oscillators**: Patterns that return to their initial state after a finite number of steps
- **Spaceships**: Patterns that translate across the grid

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
