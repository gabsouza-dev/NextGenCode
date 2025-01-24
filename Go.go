package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}

package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(100) + 1)
}

package main
import (
    "fmt"
    "time"
)
func main() {
    fmt.Println(time.Now())
}
