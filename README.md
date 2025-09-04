# Quiz Game Project

A simple quiz game written in Go.

## Features

- Read questions from a CSV file
- Timed questions with configurable timeout
- Score calculation at the end of the game
- Support for multiple-choice and true/false questions

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) version 1.16 or higher installed

### Installing the Project

1. Clone the repository:

```bash
git clone https://github.com/your-username/quiz-game.git
cd quiz-game
```

2. Build the program:

```bash
go build -o quiz-game
```

## Usage

### Basic Game Execution

```bash
./quiz-game
```

### Using a Custom Questions File

```bash
./quiz-game -csv=custom_questions.csv
```

### Setting Time Limit for Answers (seconds)

```bash
./quiz-game -timeout=30
```

### Combining Parameters

```bash
./quiz-game -csv=math_questions.csv -timeout=15
```

## Questions File Format

The questions file should be in CSV format with two columns:

1. Question
2. Correct Answer

Example:

```
"5 + 7 = ?",12
"What is the capital of Iran?","Tehran"
"Is Go a programming language?",true
```

## Building Executables

To build executables for different operating systems:

### Linux

```bash
GOOS=linux GOARCH=amd64 go build -o quiz-game-linux
```

### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o quiz-game.exe
```

### macOS

```bash
GOOS=darwin GOARCH=amd64 go build -o quiz-game-mac
```

## Testing

To run tests:

```bash
go test
```

## Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## Developers

- Your Name ([your.email@example.com](mailto:your.email@example.com))

## Support

If you encounter any issues or have questions, please open an Issue on the project's GitHub page.
