
# Exercise: Numbers and Arithmetic Operators in Go

## Learning Goals

This exercise is designed to deepen your understanding of numeric types and arithmetic operations in Go.

---

## Introduction

### Numbers in Go

Go provides basic numeric types for working with integer and floating-point values. These types vary depending on the size of the value and the architecture of the system (e.g., 32-bit vs 64-bit).

For this exercise, you'll focus on:

* **int**
  Signed integers with a minimum size of 32 bits. On most modern 64-bit systems, `int` is 64 bits in size.
  Examples: `0`, `255`, `2147483647`
  Value range on 64-bit systems: `-9223372036854775808` to `9223372036854775807`

* **float64**
  64-bit floating-point numbers, e.g., `0.0`, `3.14`

* **uint**
  Unsigned integers of the same size as `int`.
  Examples: `0`, `255`
  Value range on 64-bit systems: `0` to `18446744073709551615`

You can convert between numeric types using type conversion.

---

### Arithmetic Operators in Go

Go supports common arithmetic operators:

| Operator | Example   | Result                                 |
| -------- | --------- | -------------------------------------- |
| `+`      | `4 + 6`   | `10`                                   |
| `-`      | `15 - 10` | `5`                                    |
| `*`      | `2 * 3`   | `6`                                    |
| `/`      | `13 / 3`  | `4` (integer division drops remainder) |
| `%`      | `13 % 3`  | `1`                                    |

Go also supports shorthand assignment operators (e.g., `a += 5` means `a = a + 5`), and increment/decrement statements (`++` and `--`).

---

### Type Conversion

To convert between types, use the type name as a function:

```go
var x int = 42
f := float64(x)  // f is 42.0 of type float64

var y float64 = 11.9
i := int(y)      // i is 11 of type int
```

---

### Arithmetic Operations with Different Types

Unlike some languages, Go does **not** allow arithmetic between mismatched types without explicit conversion:

```go
var x int = 42
value := float32(2.0) * x  // This causes a compile error!

// Correct way:
value := float32(2.0) * float32(x)
```

---

## Exercise Instructions

You will write functions to analyze car production data from an assembly line.

### 1. Calculate Working Cars Produced Per Hour

The assembly line speed controls how many cars are produced each hour, but not all cars are successful. Given the total cars produced per hour and a success rate percentage (0-100), calculate the number of successful cars produced per hour.

**Example:**

```go
CalculateWorkingCarsPerHour(1547, 90) // => 1392.3
```

*Return type: `float64`*

---

### 2. Calculate Working Cars Produced Per Minute

Using the number of cars produced per hour and the success rate, calculate how many successful cars are produced per minute.

**Example:**

```go
CalculateWorkingCarsPerMinute(1105, 90) // => 16
```

*Return type: `int`*

---

### 3. Calculate Production Cost

Each car costs \$10,000 to produce individually, but producing 10 cars together costs \$95,000 (a bulk discount).

For example, producing 37 cars costs:
3 groups of 10 cars × \$95,000 + 7 individual cars × \$10,000 = \$355,000

Implement a function to calculate the total production cost for any number of cars.

**Examples:**

```go
CalculateCost(37) // => 355000
CalculateCost(21) // => 200000
```

*Return type: `uint`*

---

