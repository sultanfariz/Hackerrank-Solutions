package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

/*
 * Complete the 'minimumAbsoluteDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
 
func absDiff(a, b int32) int32 {
    diff := int32(0)
    if a < 0 && b < 0 {
        diff = a - b
    } else if a < 0 {
        diff = b - a
    } else {
        diff = a - b
    }
    
    if diff < 0 {
        diff *= -1
    }
    fmt.Println(diff)
    return diff
}

func minimumAbsoluteDifference(arr []int32) int32 {
    // Write your code here
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    min := absDiff(arr[0], arr[1])
    for i, v := range arr {
        if i == len(arr) - 1 {
            break
        }
        temp := absDiff(v, arr[i+1])
        if temp < min {
            min = temp
        }
    }
    return min
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

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := minimumAbsoluteDifference(arr)

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
