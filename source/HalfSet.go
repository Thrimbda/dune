/*
* @Author: Michael
* @Date:   2016-09-27 11:31:16
* @Last Modified by:   64509
* @Last Modified time: 2016-09-29 03:45:55
*/

package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    n := readFromFile()
    a := make([]int, 1024)
    sum := HalfSet(n, a)
    fmt.Printf("%d\n", sum)
    writeToFile(sum)
}

func HalfSet(n int, array []int) int {
    if array[n] > 0 {
        return array[n]
    } else {
        array[n] = 1
        for i := 1; i <= n / 2; i++ {
            array[n] += HalfSet(i, array)
        }
        return array[n]
    }
}

func readFromFile() int {
    inputFile := "../testfiles/input.txt"
    fin, err := os.Open(inputFile)
    defer fin.Close()
    if err != nil {
        panic(err)
    } else {
        buf := make([]byte, 32)
        num, _ := fin.Read(buf)
        fmt.Println(string(buf[0]))
        num, err := strconv.Atoi(string(buf[0]))
        if err != nil {
            panic(err)
        } else {
            return num
        }
    }
}

func writeToFile(num int) {
    outputFile := "../testfiles/output.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.Itoa(num))
    }
}