package main
import (
    "fmt"
    "strconv"
    "strings"
)

func HappyMath(p float64, r float64) (float64, float64) {
    if r == 0 {
        return p, 0
    }
    b := p / ((1 / r) + 1)
    a := b / r
    return a, b
}

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var str string
    fmt.Scan(&str)
    splitStr := strings.Split(str, "/")
    p, _ := strconv.ParseFloat(splitStr[0], 64)
    r, _ := strconv.ParseFloat(splitStr[1], 64)
    resA, resB := HappyMath(float64(p), float64(r))
    retA := fmt.Sprintf("%.10f", resA)
    retB := fmt.Sprintf("%.10f", resB)
    fmt.Printf("%s/%s", retA, retB)
}
