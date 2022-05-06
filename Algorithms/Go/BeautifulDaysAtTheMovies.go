package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func calcDif(a, b int32) float64 {
    temp := float64(a - b)
    if temp < 0 {
        return -1 * temp
    }
    return temp
}

func reverseNum(n, res int32) int32 {
    if n < 10 {
        return n + res * 10
    }
    return reverseNum(n/10, res*10 + n%10)
}

func beautifulDays(i int32, j int32, k int32) int32 {
    // Write your code here
    beautifulNumCount := int32(0)
    for {
        if i > j {
            break
        }
        fmt.Println((calcDif(i, reverseNum(i, 0)) / float64(k)))
        if math.Mod((calcDif(i, reverseNum(i, 0)) / float64(k)), 1.0) == 0 {
            beautifulNumCount++
        }
        i++
    }
    return beautifulNumCount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    i := int32(iTemp)

    jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    j := int32(jTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := beautifulDays(i, j, k)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
