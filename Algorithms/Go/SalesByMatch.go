package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
    // Write your code here
    sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
    count, pair := 1, 0
    for i:=1; i<len(ar); i++ {
        if (i == len(ar) - 1) && (ar[i] == ar[i-1]) {
            count++
            pair += count/2
            count = 1
        } else if ar[i] == ar[i-1] {
            count++
        } else {
            pair += count/2
            count = 1
        }
    }
    return int32(pair)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

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
