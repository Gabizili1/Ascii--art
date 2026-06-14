# ASCII Art Generator

A Go command-line tool that takes a string of text as input and outputs its graphical representation using stylized ASCII fonts. 

This project was built as part of the 01-edu platform curriculum.

## Features
* Supports multiple banner styles (`standard`, `shadow`, and `thinkertoy`).
* Handles special characters, spaces, and newline (`\n`) inputs.
* Built purely in Go without external dependencies.

## How to Run

Make sure you have [Go installed](https://go.dev) on your computer.

### Basic Usage
Run the program by passing your text and the desired banner style as arguments:

```bash
go run main.go "Hello World" standard
```

### Examples

**Standard Style:**
```text
 _    _      _ _         __      __         _     _ 

| |  | |    | | |        \ \    / /        | |   | |
| |__| | ___| | | ___     \ \  / /__  _ __| | __| |
|  __  |/ _ \ | |/ _ \     \ \/ / _ \| '__| |/ _` |
| |  | |  __/ | | (_) |     \  / (_) | |  | | (_| |
|_|  |_|\___|_|_|\___/       \/ \___/|_|  |_|\__,_|
                                                   
```

## Project Structure
* `main.go` - The core application entry point and logic.
* `standard.txt`, `shadow.txt`, `thinkertoy.txt` - The character templates mapping font styles.
* `main_test.go` - Automated test suites ensuring font alignment works correctly.
