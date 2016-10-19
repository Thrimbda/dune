/*
* @Author: Michael
* @Date:   2016-10-20 02:16:00
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 02:38:14
*/

package main

import (
        "fmt"
        "os"
        "io"
        "bufio"
        "strconv"
        "strings"
)

func main() {
    n, k, distance := readFromFile()
    fmt.Println(n, k, distance)
    time := CarRefuel(n, k, distance)
    fmt.Println(time)
    writeToFile(time)
}

func CarRefuel(n, k int, distance []int) int {
    var time int
    for i := 0; i < k + 1; i++ {
        if distance[i] > n {
            return -1
        }
    }
    var s int
    for i := 0; i < k + 1; i++ {
        s += distance[i]
        if s > n {
            time++
            s = distance[i]
        }
    }
    return time
}

func readFromFile() (n, k int, distance []int) {
    // this is multiple return values...with variable name.
    inputFile := "../testfiles/input_carRefuel.txt"
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
                    k, _ = strconv.Atoi(strings.Split(string(a), " ")[1])
                } else {
                    for _, value := range strings.Split(string(a), " ") {
                        currNum, _ := strconv.Atoi(value)
                        distance = append(distance, currNum)
                    }
                }
            }
        }
        return n, k, distance
    }
}

func writeToFile(num int) {
    outputFile := "../testfiles/output_carRefuel.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.Itoa(num))
    }
}
