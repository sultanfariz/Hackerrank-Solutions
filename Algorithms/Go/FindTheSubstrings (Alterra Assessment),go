package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)


/*
 * Complete the 'firstOccurrence' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING x
 */

func firstOccurrence(s string, x string) int {
    // Write your code here
    for i:=0; i<len(s); i++ {
        if s[i] == x[0] {
            for j:=1; j<len(x); j++ {
                if i+j >= len(s) {
                    return -1
                }
                if x[j] == '*' {
                    continue
                }
                if s[i+j] != x[j] {
                    break
                } else {
                    if j == len(x)-1 {
                        return i
                    }
                }
                
            }
        }
    }
    return -1
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    x := readLine(reader)

    result := firstOccurrence(s, x)

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
