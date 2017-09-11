/*
* @Author: Michael
* @Date:   2016-10-20 01:55:56
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 02:13:43
*/

package main

import (
        "fmt"
        "sort"
        "os"
        "io"
        "bufio"
        "strconv"
        "strings"
)

func main() {
    n, s, timeList := readFromFile()
    time := MultiOptimalServeOrder(n, s, timeList)
    fmt.Println(time)
    writeToFile(time)
}

func MultiOptimalServeOrder(n, s int, timeList []int) float64 {
    sumRun := make([]int, s + 1)
    sumWait := make([]int, s + 1)
    sort.Ints(timeList)
    for i := 0; i < n; i++ {
        j := i % s
        sumRun[j] += timeList[i]
        sumWait[j] += sumRun[j]
    }
    var t int
    for i := 0; i < s; i++ {
        t += sumWait[i]
    }
    return float64(t)/float64(n)
}

func readFromFile() (n, s int, timeList []int) {
    // this is multiple return values...with variable name.
    inputFile := "../testfiles/input_multiOptimalServeOrder.txt"
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
                    n, _ = strconv.Atoi(strings.Split(string(a), " ")[0])
                    s, _ = strconv.Atoi(strings.Split(string(a), " ")[1])
                } else {
                    for _, value := range strings.Split(string(a), " ") {
                        currNum, _ := strconv.Atoi(value)
                        timeList = append(timeList, currNum)
                    }
                }
            }
        }
        return n, s, timeList
    }
}

func writeToFile(num float64) {
    outputFile := "../testfiles/output_multiOptimalServeOrder.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.FormatFloat(num, 'f', -1, 64))
    }
}
