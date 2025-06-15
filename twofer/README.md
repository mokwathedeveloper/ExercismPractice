# Two Fer

A simple Go package that implements a sharing message generator following the "one for X, one for me" pattern.

## Description

This package provides a function `ShareWith` that generates messages in the format "One for X, one for me" where X is either a provided name or "you" if no name is given.

## Installation

To use this package in your Go project:

```bash
go get github.com/yourusername/twofer
```

## Usage

Here's how to use the package in your Go code:

```go
package main

import (
    "fmt"
    "twofer"
)

func main() {
    // With a name
    fmt.Println(twofer.ShareWith("Alice"))
    // Output: One for Alice, one for me.

    // Without a name (empty string)
    fmt.Println(twofer.ShareWith(""))
    // Output: One for you, one for me.
}
```

## Examples

The package includes an example program in the `cmd/twofer` directory that demonstrates its usage.

To run the example:

```bash
cd cmd/twofer
go run main.go
```

## Testing

The package includes comprehensive tests. To run them:

```bash
go test -v
```

## Function Documentation

### `ShareWith(name string) string`

Returns a sharing message based on the provided name.

Parameters:
- `name`: A string representing the name of the person to share with.
  - If empty, defaults to "you"

Returns:
- A string in the format "One for X, one for me."

Examples:
```go
ShareWith("Alice") // returns "One for Alice, one for me."
ShareWith("")      // returns "One for you, one for me."
ShareWith("Bob")   // returns "One for Bob, one for me."
```

## Project Structure

```
twofer/
├── twofer.go         # Main package implementation
├── twofer_test.go    # Test suite
├── go.mod           # Module definition
├── README.md        # This documentation
└── cmd/
    └── twofer/
        └── main.go  # Example program
```

## License

This project is open source and available under the [MIT License](LICENSE). 