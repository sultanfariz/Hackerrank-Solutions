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
 * Complete the 'minimumGroups' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY predators as parameter.
 */
 
func recursiveTraverse(predators []int32, index int32) int {
    if predators[index] == -1 {
        return 1
    }
    ret := recursiveTraverse(predators, predators[index]) + 1
    return ret
}

func minimumGroups(predators []int32) int32 {
    // Write your code here
    streak := 0
    for i:=0; i<len(predators); i++ {
        tempStreak := recursiveTraverse(predators, int32(i))
        if tempStreak > streak {
            streak = tempStreak
        }
    }
    
    return int32(streak)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    predatorsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var predators []int32

    for i := 0; i < int(predatorsCount); i++ {
        predatorsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        predatorsItem := int32(predatorsItemTemp)
        predators = append(predators, predatorsItem)
    }

    result := minimumGroups(predators)

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
