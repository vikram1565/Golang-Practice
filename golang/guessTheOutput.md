# **Guess the output**

## *What is the output of following program*

```go
package main
import "fmt"

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

func main() {
    var n number
    defer n.print()
    defer n.pprint()
    defer func() { n.print() }()
    defer func() { n.pprint() }()
    n = 3
}

```

## Answer

```markdown
3
3
3
0
```

Explanation -

defer works in LIFO manner and evaluates the values when defined. The first defer after main will be executed at last and print the value 0. The second defer will print 3 as it is a pointer to the value of n which is 3 at the time of execution. In case of third and fourth defer, we are passing the methods with wrapped anonymous function. This will print the value of 'n' at the time of execution.
