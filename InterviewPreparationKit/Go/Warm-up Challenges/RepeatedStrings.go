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
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
    // Write your code here
    aFreq := 0
    arrLen := len(s)
    
    for i:=0; i<arrLen; i++ {
        if s[i] != 'a' {
            break
        } else if i+1 == arrLen {
            return n
        }
    }
    
    div := (int(n)/arrLen) //div = 3
    mult := (div * arrLen) //mult = 9
    
    for i:=0; i<arrLen; i++ {
        // fmt.Printf("%c", s[i])
        if s[i] != 'a' {
            continue
        }
        
        if mult < (int(n)-i) {
            aFreq += div + 1
            fmt.Printf("anying")
        } else {
            fmt.Printf("anying")
            aFreq += div
        }
    }
    return int64(aFreq)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

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
