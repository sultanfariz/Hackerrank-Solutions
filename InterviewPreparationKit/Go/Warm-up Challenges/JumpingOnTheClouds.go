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
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
    // Write your code here
    cloudInARow, jumps := 0, 0
    for i:=0; i<len(c); i++ {
        if c[i] == 0 {
            cloudInARow++
        }else if c[i] != 0 {
            jumps += int(cloudInARow/2)
            jumps++
            cloudInARow = 0
        }
        
        if i == (len(c) - 1) {
            jumps += int(cloudInARow/2)
            cloudInARow = 0
        }
    }
    return int32(jumps)
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

    cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    result := jumpingOnClouds(c)

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
