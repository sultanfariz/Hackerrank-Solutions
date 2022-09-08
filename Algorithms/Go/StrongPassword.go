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
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
 */

func minimumNumber(n int32, password string) int32 {
    // Return the minimum number of characters to make the password strong
    numbers := "0123456789"
    lower_case := "abcdefghijklmnopqrstuvwxyz"
    upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    special_characters := "!@#$%^&*()-+"
    
    lenFlag, numFlag, lowFlag, upFlag, charFlag := 0, 1, 1, 1, 1
    // numFlag, lowFlag, upFlag, charFlag := false
    
    if len(password) < 6 {
        lenFlag = 6 - len(password)
    }
    
    for _, v := range password {
        if strings.Contains(numbers, string(v)) { 
            numFlag = 0
        } else if strings.Contains(lower_case, string(v)) { 
            lowFlag = 0
        } else if strings.Contains(upper_case, string(v)) { 
            upFlag = 0
        } else if strings.Contains(special_characters, string(v)) { 
            charFlag = 0
        }
    }
    
    min := numFlag + lowFlag + upFlag + charFlag
    
    if lenFlag < min {
        fmt.Println(min)
        return int32(min)
    } else {
        return int32(lenFlag)
    }
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

    password := readLine(reader)

    answer := minimumNumber(n, password)

    fmt.Fprintf(writer, "%d\n", answer)

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
