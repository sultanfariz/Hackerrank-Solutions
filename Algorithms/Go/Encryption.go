package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
    // Write your code here
    strLen := len(s)
    sqrtRes := math.Sqrt(float64(strLen))
    rowNum, colNum := math.Floor(sqrtRes), math.Ceil(sqrtRes)
    if int(rowNum * colNum) < strLen {
        if rowNum > colNum {
            colNum = rowNum
        } else {
            rowNum = colNum
        }
    }
    
    retString := ""
    strMap := make(map[int]string)
    for idx, val := range s {
        strMap[idx] = string(val)
    }
    for i:=0; i<int(colNum); i++ {
        for j:=0; j<int(strLen); j+=int(colNum) {
            retString = fmt.Sprintf("%s%s", retString, strMap[j+i])
        }
        retString = fmt.Sprintf("%s%s", retString, " ")
    }
    return retString
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

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
