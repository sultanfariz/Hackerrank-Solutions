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
 * Complete the 'gamingArray' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// func recursiveGame(player string, arr []int32, sortedArr []int32) string {
//     if player == "BOB" {
//         player = "ANDY"
//     } else {
//         player = "BOB"
//     }
//     if len(arr) == 0 {
//         return player
//     }
//     max, maxIdx := arr[0], 0
//     for idx, val := range arr {
//         if max < val {
//             max, maxIdx = val, idx
//         }
//     }
//     return recursiveGame(player, arr[:maxIdx], sortedArr)
// }

// func gamingArray(arr []int32) string {
//     // Write your code here
//     player := "BOB"
//     arrIdxMap := make(map[int32]int)
//     for idx, val := range arr {
//         arrIdxMap[val] = idx
//     }
    
//     sortedArr := make([]int32, len(arr))
//     sortedArr = arr
    
//     winner := recursiveGame(player, arr, sortedArr)
//     return winner
// }

func gamingArray(arr []int32) string {
    count, max := 0, arr[0]
    for _, val := range arr {
        if val > max {
            count++
            max = val
        }
    }
    if count%2 == 0 {
        return "BOB"
    }
    return "ANDY"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    g := int32(gTemp)

    for gItr := 0; gItr < int(g); gItr++ {
        arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)

        arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var arr []int32

        for i := 0; i < int(arrCount); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := gamingArray(arr)

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
