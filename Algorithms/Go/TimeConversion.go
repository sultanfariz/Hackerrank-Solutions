package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Write your code here
    stringArr := strings.Split(s, ":")
    thirdEl := strings.Split(stringArr[2], "")
    seconds := thirdEl[0] + thirdEl[1]
    meridiem := thirdEl[2] + thirdEl[3]
    
    hour, err := strconv.Atoi(stringArr[0])
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    
    var hourString string
    if meridiem == "PM" {
        if hour != 12 {
            hour = int(hour) + 12
        }
        hourString = strconv.Itoa(hour)
    }else if meridiem == "AM" {
        if hour == 12 {
            hour = int(hour) - 12
        }
        hourString = strconv.Itoa(hour)
        if hour < 10 {
            hourString = "0" + hourString
        }
    }
    
    return hourString + ":" + stringArr[1] + ":" + seconds
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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
