package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)



/*
 * Complete the 'getActualFloor' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY numbers
 *  2. INTEGER floor
 */
 
func containNumber(number, target int32) bool {
    pangkat := math.Log10(float64(target))
    pangkat = float64(int(pangkat))
    modifier := math.Pow(10, (pangkat+1))
    for {
        mod := number%int32(modifier)
        if mod == target {
            return true
        }
        if number < int32(modifier) {
            return false
        }
        number = number/10
    }
}

func getActualFloor(numbers []int32, floor int32) int32 {
    if(floor<0){
        floor*=-1
    }
    unlucky := 0
    for i:=0; int32(i)<=floor; i++ {
        for j:=0; j<len(numbers); j++ {
            if containNumber(floor, numbers[j]){
                return 0
            }
            if numbers[j] <= floor {   
                if containNumber(int32(i), numbers[j]){
                    unlucky += 1
                    break;
                }
            }
        }
    }
    return floor-int32(unlucky)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var numbers []int32

    for i := 0; i < int(numbersCount); i++ {
        numbersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        numbersItem := int32(numbersItemTemp)
        numbers = append(numbers, numbersItem)
    }

    floorTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    floor := int32(floorTemp)

    result := getActualFloor(numbers, floor)

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
