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
 * Complete the 'maxDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY px as parameter.
 */

func maxDifference(px []int32) int32 {
    // Write your code here
    var min int32 = px[0]
    var diff int32 = -1
    var maxDiff int32 = -1
    for _, val := range px {
        if val < min{
            min = val
        }
        diff = val - min
        if diff > maxDiff {
            fmt.Print(val, diff, min)
            maxDiff = diff
        }
    }
    if maxDiff == 0 {
        return -1
    }
    return maxDiff
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    pxCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var px []int32

    for i := 0; i < int(pxCount); i++ {
        pxItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        pxItem := int32(pxItemTemp)
        px = append(px, pxItem)
    }

    result := maxDifference(px)

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
