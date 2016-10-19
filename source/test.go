/*
* @Author: Michael
* @Date:   2016-10-19 22:07:52
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 01:02:28
*/

package main

import (
        "fmt"
        "math/rand"
        "sort"
)

type activity struct {
    start, end int
}

type activities [] *activity

func (p activities) Len() int {
    return len(p)
}

func (p activities) Swap(i int, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p activities) Less(i int, j int) bool {
    return p[i].end > p[j].end
}

func main() {
    a := make(activities, 8)
    for i := 0; i < len(a); i++ {
        a[i] = &activity{i + 1, rand.Intn(1000)}
    }
    printA(a)
    sort.Sort(a)
    printA(a)
}

func printA (a activities) {
    for _, value := range a {
        fmt.Print(value.end, " ")
    }
    fmt.Println(" ")
}