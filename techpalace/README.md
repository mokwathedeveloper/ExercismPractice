# Tech Palace Display Project

This Go project simulates a smart display used in a tech store for personalized welcome messages and marketing content cleanup.

## Features

- `WelcomeMessage(name string)`: Displays a welcome message with the name in uppercase.
- `AddBorder(message string, numStars int)`: Adds a configurable star border around a message.
- `CleanUpMessage(message string)`: Cleans old messages by removing stars and trimming spaces.

## Run

```bash
go run .
```

## Test

```bash
go test
``` 