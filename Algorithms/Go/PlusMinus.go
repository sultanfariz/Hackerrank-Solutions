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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var pos,neg,zero int32 = 0, 0, 0
    var n int32 = int32(len(arr))
    for i:=0; i<len(arr); i++{
        if arr[i]>0 {
            pos++
        }else if arr[i]<0{
            neg++
        }else {
            zero++
        }
    }
    posProp, negProp, zeroProp := float64(pos)/float64(n), float64(neg)/float64(n), float64(zero)/float64(n)
    fmt.Printf("%.6f\n", posProp)
    fmt.Printf("%.6f\n", negProp)
    fmt.Printf("%.6f\n", zeroProp)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

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

    plusMinus(arr)
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
