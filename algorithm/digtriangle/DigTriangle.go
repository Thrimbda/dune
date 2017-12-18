/*
* @Author: Michael
* @Date:   2016-10-19 21:03:39
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-19 23:52:09
 */

package digtriangle

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func run() {
	var digTriangle [][]int
	num, digTriangle := readFromFile()
	fmt.Println(num)
	fmt.Println(digTriangle)
	writeToFile(MaxSum(0, 0, num-1, digTriangle))
}

func MaxSum(i, j, num int, digTriangle [][]int) int {
	if i == num {
		return digTriangle[i][j]
	} else {
		x := MaxSum(i+1, j, num, digTriangle)
		y := MaxSum(i+1, j+1, num, digTriangle)
		return int(math.Max(float64(x), float64(y))) + digTriangle[i][j]
	}
}

func readFromFile() (num int, digTriangle [][]int) {
	// this is multiple return values...with variable name.
	inputFile := "../testfiles/input_digTriangle.txt"
	fin, err := os.Open(inputFile)
	defer fin.Close()
	if err != nil {
		panic(err)
	} else {
		rd := bufio.NewReader(fin)
		for i := 0; ; i++ {
			a, _, c := rd.ReadLine()
			if c == io.EOF {
				break
			} else {
				if i == 0 {
					num, _ = strconv.Atoi(string(a))
					digTriangle = make([][]int, num)
				} else {
					for _, value := range strings.Split(string(a), " ") {
						currNum, _ := strconv.Atoi(value)
						digTriangle[i-1] = append(digTriangle[i-1], currNum)
					}
				}
			}
		}
		return num, digTriangle
	}
}

func writeToFile(num int) {
	outputFile := "../testfiles/output_digTriangle.txt"
	fout, err := os.Create(outputFile)
	defer fout.Close()
	if err != nil {
		panic(err)
	} else {
		fout.WriteString(strconv.Itoa(num))
	}
}
