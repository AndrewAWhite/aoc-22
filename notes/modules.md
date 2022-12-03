# Go Modules

I'll be splitting the solutions into a separate package, so will make use of go modules. To setup:

At project root:
```bash
$ go mod init github.com/AndrewAWhite/aoc-22
```

Then within the `main.go`, the solutions are imported via:

```go
import (
    "github.com/AndrewAWhite/aoc-22/solutions"
)
```

Which are then called according to the naming convention I've chosen for this project ("Solution_$day_$part") like so:

```go
func main() {
	fmt.Printf("Part 1: %s\n", solutions.Solution_1_1())
	fmt.Printf("Part 2: %s\n", solutions.Solution_1_2())
}
```

