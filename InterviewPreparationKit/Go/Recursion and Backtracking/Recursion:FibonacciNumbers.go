package main
import "fmt"

func fibonacci(n int) int {
    // Write your code here.
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    fmt.Println(fibonacci(n))
}
