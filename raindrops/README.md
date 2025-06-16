# Raindrops

A fun Go implementation of the classic **Raindrops** coding challenge — now with a custom `intToString` function instead of importing `strconv`.

## Description

The `Raindrops` program converts a number into a string that contains raindrop sounds depending on the number’s factors:

- If the number is divisible by **3**, it adds `"Pling"` to the result.
- If the number is divisible by **5**, it adds `"Plang"` to the result.
- If the number is divisible by **7**, it adds `"Plong"` to the result.
- If the number is not divisible by **3**, **5**, or **7**, the number itself is returned as a string using a custom conversion function.

## Project Structure

```
raindrops_project_custom_itoa/
├── go.mod
├── main.go
├── raindrops/
│   ├── raindrops.go
│   └── raindrops_test.go
└── README.md
```

## Getting Started

### Run the program
```bash
go run .
```

### Run tests
```bash
go test ./raindrops
```

### Run benchmark
```bash
go test -bench=. ./raindrops
```

## No strconv Used

Instead of importing `strconv`, this project defines a custom `intToString` function for integer-to-string conversion.

## Author

**Mokwa Moffat**
Full-stack Developer | Zone01 Apprentice
[GitHub](https://github.com/mokwathedeveloper)