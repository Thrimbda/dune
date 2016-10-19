/*
* @Author: Michael
* @Date:   2016-10-19 22:35:50
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 00:01:37
*/

package main

import (
        "fmt"
        "strings"
        "math"
        "strconv"
        "os"
        "io"
        "bufio"
)

func main() {
    n, r := readFromFile()
    fmt.Println(n, r)
    min := MinCost(n, r)
    fmt.Println(min)
    writeToFile(min)
}

func MinCost(n int, r [][]int) int {
    for l := 2; l < n; l++ {
        for i := 0; i < n - l; i++ {
            j := i + l
            for k := i + 1; k < j; k++ {
                r[i][j] = int(math.Min(float64(r[i][j]), float64(r[i][k] + r[k][j])))
            }
        }
    }
    return r[0][n - 1]
}

func readFromFile() (num int, r [][]int) {
    // this is multiple return values...with variable name.
    inputFile := "../testfiles/input_yachtLend.txt"
    fin, err := os.Open(inputFile)
    defer fin.Close()
    if err != nil {
        panic(err)
    } else {
        rd := bufio.NewReader(fin)
        for i := 0; ; i++{
            a, _, c := rd.ReadLine()
            if c == io.EOF {
                break
            } else {
                if i == 0 {
                    num, _ = strconv.Atoi(string(a))
                    r = make([][]int, num - 1)
                } else {
                    r[i - 1] = make([]int, num + 1)
                    for k, value := range strings.Split(string(a), " ") {
                        j := i + k
                        currNum, _ := strconv.Atoi(value)
                        r[i - 1][j] = currNum
                    }
                }
            }
        }
        return num, r
    }
}

func writeToFile(num int) {
    outputFile := "../testfiles/output_yachtLend.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.Itoa(num))
    }
}
