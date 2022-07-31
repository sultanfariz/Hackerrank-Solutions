package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */
func swap(arr []int32, a, b int) []int32 {
    temp := arr[a]
    arr[a] = arr[b]
    arr[b] = temp
    return arr
}

func countSwaps(a []int32) {
    // Write your code here
    count := 0
    for i:=0; i<len(a); i++ {
        for j, val2 := range a {
            if j < len(a)-1 && val2 > a[j+1] {
                swap(a, j, j+1)
                count++
            }
        }
    }
    fmt.Printf("Array is sorted in %d swaps.\n", count)
    fmt.Printf("First Element: %d\n", a[0])
    fmt.Printf("Last Element: %d\n", a[len(a)-1])
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
