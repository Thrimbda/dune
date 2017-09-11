/*
* @Author: Michael
* @Date:   2016-10-20 02:41:14
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 03:07:05
*/

package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
    n := readFromFile()
    fmt.Println(n)
    result := OptimalDecomposition(n)
    fmt.Println(result)
    writeToFile(result)
}

func OptimalDecomposition(n int) int {
    var sum, result, offset, i int
    for i = 2; ; i++ {
        sum += i
        if sum >= n {
            break
        }
    }
    offset = sum - n
    result = 1
    if offset == 1 {
        for j := 3; j < i; j++ {
            result *= j
        } 
        i++
        result *= i
    } else {
        for j := 2; j <= i; j++ {
            if j == offset {
                continue
            }
            result *= j
        }
    }
    return result
}

func readFromFile() int {
    inputFile := "../testfiles/input_optimalDecomposition.txt"
    fin, err := os.Open(inputFile)
    defer fin.Close()
    if err != nil {
        panic(err)
    } else {
        buf := make([]byte, 64)
        list, _ := fin.Read(buf)
        if err != nil {
            panic(err)
        } else {
            strings.Split(string(buf[:list]), " ")
            n, _ := strconv.Atoi(string(buf[:list]))
            return n
        }
    }
}

func writeToFile(num int) {
    outputFile := "../testfiles/output_optimalDecomposition.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.Itoa(num))
    }
}
