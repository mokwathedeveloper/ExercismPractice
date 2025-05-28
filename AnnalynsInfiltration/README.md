# Introduction

Booleans in Go are represented by the predeclared boolean type `bool`, which can hold the values `true` or `false`. It is a defined type.

```go
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false'
```

Go supports three logical operators to evaluate expressions to Boolean values:

| Operator | Meaning                                           |    |                                                        |
| -------- | ------------------------------------------------- | -- | ------------------------------------------------------ |
| `&&`     | Logical AND: true if **both** statements are true |    |                                                        |
| \`       |                                                   | \` | Logical OR: true if **at least one** statement is true |
| `!`      | Logical NOT: true if the statement is false       |    |                                                        |

# Instructions

In this exercise, you'll implement quest logic for a new RPG game in development. The story follows Annalyn, a brave girl with a loyal dog. Her best friend has been kidnapped while foraging in the forest. Annalyn sets out to rescue them, possibly with the help of her dog.

Annalyn discovers the kidnappers' camp where her friend is held. Two enemies are present: a **knight** and an **archer**.

Annalyn can take several actions based on the state of the enemies and her dog:

### 1. Fast Attack

* **Condition:** The knight is asleep.
* **Reasoning:** The knight takes time to don armor and is vulnerable while sleeping.

### 2. Spy

* **Condition:** At least one of the knight or archer is awake.
* **Reasoning:** If everyone is asleep, spying is pointless.

### 3. Signal Prisoner

* **Condition:** The prisoner is awake **and** the archer is asleep.
* **Reasoning:** Archers are trained to intercept bird calls.

### 4. Free Prisoner

* **With dog:** Can free the prisoner if the archer is asleep (the knight fears the dog).
* **Without dog:** Can free the prisoner if:

  * The prisoner is awake, and
  * Both the knight and archer are asleep.
* **Note:** If the prisoner is asleep, rescue fails — the sudden approach could alert the enemies.

You will implement four functions to determine if these actions are possible based on the state of the characters.

---

### Function: `CanFastAttack()`

Determine if a fast attack can be made.

**Signature:**

```go
func CanFastAttack(knightIsAwake bool) bool
```

**Example:**

```go
var knightIsAwake = true
fmt.Println(CanFastAttack(knightIsAwake))
// Output: false
```

---

### How to Debug

When a test fails, a descriptive message will be shown. You can also use console output for debugging:

```go
import "fmt"
fmt.Println("Debug message")
```

> ⚠️ Note: In in-browser editors, avoid using `fmt.Print` without a newline (`\n`) to prevent test output issues.

---

