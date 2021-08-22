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
 * Complete the 'minimizeBias' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ratings as parameter.
 */

func minimizeBias(ratings []int32) int32 {
    // Write your code here
    var bias int32 = 0
    sort.Slice(ratings, func(i, j int) bool { return ratings[i] < ratings[j] })
    fmt.Println(ratings)
    for i:=0; i<len(ratings); i+=2 {
        bias += ratings[i+1] - ratings[i]
    }
    return bias
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    ratingsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var ratings []int32

    for i := 0; i < int(ratingsCount); i++ {
        ratingsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        ratingsItem := int32(ratingsItemTemp)
        ratings = append(ratings, ratingsItem)
    }

    result := minimizeBias(ratings)

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
