
# Go Programming Language Demonstration

This project demonstrates the Go programming language using examples of basic syntax, concurrency, and sequential execution.

## Requirements

- Go programming language
- Visual Studio Code (recommended)
- Go extension for VS Code (recommended)

## How to Install Go

1. Visit https://go.dev/dl/
2. Download Go for your operating system (Windows, macOS, or Linux).
3. Open the installer or installation package and follow the instructions for your operating system.
4. After installing, open a new terminal.
5. Type the following command:

    go version

If the terminal displays a Go version number, Go was installed correctly.

## How to Run the Programs

1. Open the Go-Lang-Demonstration folder in VS Code.
2. Open the terminal.
3. Make sure the terminal is in the main project folder.
4. Run the following programs one at a time.

### 1. Hello World

Command:

    go run "Simple Syntax/hello.go"

This program prints "Hello, World!" to the terminal.

Expected output:

    Hello, World!

### 2. Simple Go Syntax

Command:

    go run "Simple Syntax/simple.go"

This program demonstrates basic Go syntax using student names and grades.

It includes:
- Structs
- Functions
- Loops
- If-else statements
- Slices

The program displays each student's name, numerical grade, and letter grade.

### 3. Concurrent Execution

Command:

    go run concurrency.go

This program runs three tasks concurrently using goroutines and a channel.

Each task waits for 4 seconds. Since the tasks run concurrently, they finish in approximately 4 seconds total.

Tested execution time: 4.0006136 seconds.

### 4. Sequential Execution

Command:

    go run concurrency1.go

This program runs three tasks one at a time.

Each task waits for 4 seconds. Since the tasks run sequentially, they finish in approximately 12 seconds total.

Tested execution time: 12.0017518 seconds.

## Concurrency vs. Sequential Execution

| Concurrent Execution | Sequential Execution |
|---|---|
| Tasks run concurrently | Tasks run one at a time |
| Uses goroutines | Uses regular function calls |
| Approximately 4 seconds | Approximately 12 seconds |

The concurrency example shows how overlapping tasks can reduce the total execution time.

## Project Structure

    Go-Lang-Demonstration/
    |
    |-- Simple Syntax/
    |   |-- go.mod
    |   |-- hello.go
    |   |-- simple.go
    |
    |-- concurrency.go
    |-- concurrency1.go
    |-- README.md
    |-- .gitignore

## Project Files

- `hello.go` - Prints Hello World.
- `simple.go` - Demonstrates basic Go syntax.
- `concurrency.go` - Demonstrates concurrent execution.
- `concurrency1.go` - Demonstrates sequential execution.
- `main.mod` - Web Application.
- `main.mod` - Defines Methods for Web Application.
- `README.md` - Contains project instructions.

## References

- Official Go Website: https://go.dev/
- Go Installation Guide: https://go.dev/doc/install
- Go Documentation: https://go.dev/doc/
- GORM Library: https://pkg.go.dev/gorm.io/gorm 
- Gorilla/Mux: https://github.com/gorilla/mux?tab=readme-ov-file 
- Go Modules: https://go.dev/blog/using-go-modules 
- Creating GoLang Rest API: https://www.youtube.com/watch?v=KPftgI40WHI 