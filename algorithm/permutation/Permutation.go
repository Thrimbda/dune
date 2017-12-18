/*
* @Author: 64509
* @Date:   2016-09-29 03:54:15
* @Last Modified by:   64509
* @Last Modified time: 2016-09-29 05:00:40
 */

package permutation

import (
	"fmt"
	"os"
)

func run() {
	// list := readFromFile()
	list := []byte{'a', 'b', 'c', 'd'}
	// result := make([]byte,0 , 1024)
	var result []byte
	result = Permutation(list, result, 0, len(list)-1)
	writeToFile(result)
}

func Permutation(list, result []byte, low, high int) []byte {
	if low == high {
		result = append(result, list...)
		return append(result, '\n')
	} else {
		for i := low; i <= high; i++ {
			Swap(list, low, i)
			result = Permutation(list, result, low+1, high)
			Swap(list, low, i)
		}
		return result
	}
}

func Swap(list []byte, i, j int) {
	list[i], list[j] = list[j], list[i]
}

func readFromFile() []byte {
	inputFile := "../testfiles/input_permutation.txt"
	fin, err := os.Open(inputFile)
	defer fin.Close()
	if err != nil {
		panic(err)
	} else {
		buf := make([]byte, 32)
		list, _ := fin.Read(buf)
		fmt.Println(buf[:list])
		return buf[:list]
	}
}

func writeToFile(result []byte) {
	outputFile := "../testfiles/output_permutation.txt"
	fout, err := os.Create(outputFile)
	defer fout.Close()
	if err != nil {
		panic(err)
	} else {
		fout.WriteString(string(result))
	}
}
