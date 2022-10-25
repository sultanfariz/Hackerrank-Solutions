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
 * Complete the 'hackerrankInString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func hackerrankInString(s string) string {
    // Write your code here
    hackerrankStr := "hackerrank"
    count := 0
    for _, v := range s {
        if string(v) == string(hackerrankStr[count]) {
            if count == len(hackerrankStr)-1 {
                return "YES"
            }
            count++
        }
    }
    return "NO"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := hackerrankInString(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
