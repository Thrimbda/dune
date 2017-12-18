/*
* @Author: Michael
* @Date:   2016-10-20 00:43:47
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-20 01:19:40
 */

package activityarrange

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type activity struct {
	start, end int
}

func (self activity) String() string {
	return fmt.Sprintf("<activity(%d, %d)>", self.start, self.end)
}

type activities []*activity

func (self activities) Len() int {
	return len(self)
}

func (self activities) Swap(i int, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self activities) Less(i int, j int) bool {
	return self[i].end < self[j].end
}

func activityArrange(n int, acts activities) int {
	sort.Sort(acts)
	var count int
	currTime := -1
	for i := 0; i < n; i++ {
		if acts[i].start > currTime {
			count++
			currTime = acts[i].end
		}
	}
	return count
}

func run() {
	n, acts := readFromFile()
	count := activityArrange(n, acts)
	fmt.Print(count)
	writeToFile(count)
}

func readFromFile() (num int, acts activities) {
	// this is multiple return values...with variable name.
	inputFile := "../testfiles/input_activityArrange.txt"
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
					acts = make(activities, num)
				} else {
					start, _ := strconv.Atoi(strings.Split(string(a), " ")[0])
					end, _ := strconv.Atoi(strings.Split(string(a), " ")[1])
					acts[i-1] = &activity{start, end}
				}
			}
		}
		return num, acts
	}
}

func writeToFile(num int) {
	outputFile := "../testfiles/output_activityArrange.txt"
	fout, err := os.Create(outputFile)
	defer fout.Close()
	if err != nil {
		panic(err)
	} else {
		fout.WriteString(strconv.Itoa(num))
	}
}
