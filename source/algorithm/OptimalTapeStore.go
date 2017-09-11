/*
* @Author: Michael
* @Date:   2016-10-20 01:23:56
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 01:51:44
*/

package main

import (
        "fmt"
        "os"
        "io"
        "sort"
        "bufio"
        "strings"
        "strconv"
)

type program struct {
    length, probability int
}

func (self program) String() string {
    return fmt.Sprintf("<program(%d, %d)>", self.length, self.probability)
}

func main() {
    n, programs := readFromFile()
    time := optimalTapeStore(n, programs)
    fmt.Println(time)
    writeToFile(time)
}

func optimalTapeStore (n int, programs []program) float64 {
    var sum, time int
    revisedProbability := make([]int, len(programs))
    for i := 0; i < n; i++ {
        revisedProbability[i] = programs[i].length * programs[i].probability
        sum += programs[i].probability
    }
    sort.Ints(revisedProbability)

    for i := 0; i < n; i++ {
        time += (n - i) * revisedProbability[i]
    }
    return float64(time)/float64(sum)
}

func readFromFile() (num int, programs []program) {
    // this is multiple return values...with variable name.
    inputFile := "../testfiles/input_optimalTapeStore.txt"
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
                    programs = make([]program, num)
                } else {
                    start, _ := strconv.Atoi(strings.Split(string(a), " ")[0])
                    end, _ := strconv.Atoi(strings.Split(string(a), " ")[1])
                    programs[i - 1] = program{start, end}
                }
            }
        }
        return num, programs
    }
}

func writeToFile(num float64) {
    outputFile := "../testfiles/output_optimalTapeStore.txt"
    fout, err := os.Create(outputFile)
    defer fout.Close()
    if err != nil {
        panic(err)
    } else {
        fout.WriteString(strconv.FormatFloat(num, 'f', -1, 64))
    }
}
