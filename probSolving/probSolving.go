package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type prob struct {
// 	value int64
// }
// type node struct {
// 	this  prob
// 	nodes []prob
// }

func main() {

	input := readInput()

	for i := 0; i <= len(input)-1; i++ {

		line1 := strings.Split(input[i], " ")
		_, _ = strconv.ParseInt(trimNewLine(line1[0]), 10, 65)
		diff, _ := strconv.ParseInt(trimNewLine(line1[1]), 10, 64)

		line2 := strings.Split(input[i+1], " ")

		arr1 := convertToIntArray(line2)
		arr2 := constructDAG(arr1, diff)
		//fmt.Println(arr2)
		//fmt.Println("----------------output-------------------")
		var path []int64
		var pathes [][]int64
		var pathOfPathes [][][]int64
		for i := 0; i <= len(arr2)-1; i++ {

			_, op2 := makeBranch(arr2, i, path, pathes)
			pathOfPathes = append(pathOfPathes, op2)

		}
		days := 0
		for i := 0; i <= len(arr1)-1; i++ {
			if arr1[i] == 0 {
				continue
			}
			path := getLongestBranch(pathOfPathes, arr1[i])
			temp := make([]int64, len(path))
			copy(temp, path)
			pathOfPathes = removePath(pathOfPathes, temp)
			arr1 = addZero(arr1, path)
			days++
		}
		fmt.Println(days)

		// //fmt.Println(pathOfPathes)
		// paths := getLongestPath(pathOfPathes)

		// sortedpaths := sort2DArray(paths)
		// fmt.Println(calucateDays(sortedpaths, arr1))
		i = i + 1
	}
}

func calucateDays(arr [][]int64, input []int64) int64 {
	var days int64
	days = 0
	for i := 0; i <= len(arr)-1; i++ {
		if checkElementExists(input, arr[i][0]) {

			days++
			fmt.Printf("%v\n", arr[i])
			input = removeElements(input, arr[i])
		}
	}
	return days
}

func checkElementExists(arr []int64, elem int64) bool {
	for i := 0; i <= len(arr)-1; i++ {
		if elem == arr[i] {
			return true
		}
	}
	return false
}

func removeElements(arr1 []int64, arr2 []int64) []int64 {
	for i := 0; i <= len(arr1)-1; i++ {
		for j := 0; j <= len(arr2)-1; j++ {
			if arr1[i] == arr2[j] {
				arr1[i] = 0
				break
			}
		}
	}
	return arr1
}

func sort2DArray(arr [][]int64) [][]int64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-2; j++ {
			if len(arr[j]) < len(arr[j+1]) {
				var tmp []int64
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func getLongestBranch(pathOfPathes [][][]int64, startWith int64) []int64 {
	var maxPath []int64
	maxStartIndex := 0
	maxPathIndex := 0
	for i := 0; i <= len(pathOfPathes)-1; i++ {
		if len(pathOfPathes[i]) > 0 {
			if pathOfPathes[i][0][0] == startWith {
				maxPathSize := len(pathOfPathes[i][0])
				for j := 1; j <= len(pathOfPathes[i])-1; j++ {
					if maxPathSize < len(pathOfPathes[i][j]) {
						maxPathSize = len(pathOfPathes[i][j])
						maxPathIndex = j
						maxStartIndex = i

					}
				}
			}
		}
	}
	maxPath = pathOfPathes[maxStartIndex][maxPathIndex]
	return maxPath
}

func removePath(pathOfPathes [][][]int64, path []int64) [][][]int64 {

	for l := 0; l <= len(path)-1; l++ {
		for i := 0; i <= len(pathOfPathes)-1; i++ {
			if len(pathOfPathes[i]) > 0 {
				for j := 0; j <= len(pathOfPathes[i])-1; j++ {
					for k := 0; k <= len(pathOfPathes[i][j])-1; k++ {
						if pathOfPathes[i][j][k] == path[l] {
							prev := k - 1
							if prev < 0 {
								prev = 0
							}
							next := k + 1
							if next >= len(pathOfPathes[i][j]) {
								next = len(pathOfPathes[i][j]) - 1
							}
							if prev == next {
								pathOfPathes[i][j] = []int64{0}
								break
							} else {
								pathOfPathes[i][j] = append(pathOfPathes[i][j][0:prev], pathOfPathes[i][j][next:]...)
							}
							k = 0
						}

					}
				}
			}
		}
	}

	return pathOfPathes
}

func getLongestPath(pathOfPathes [][][]int64) [][]int64 {
	var maxPath [][]int64
	for i := 0; i <= len(pathOfPathes)-1; i++ {
		if len(pathOfPathes[i]) > 0 {
			maxPathIndex := 0
			maxPathSize := len(pathOfPathes[i][0])
			for j := 1; j <= len(pathOfPathes[i])-1; j++ {
				if maxPathSize < len(pathOfPathes[i][j]) {
					maxPathSize = len(pathOfPathes[i][j])
					maxPathIndex = j
				}
			}
			maxPath = append(maxPath, pathOfPathes[i][maxPathIndex])
		}
	}
	return maxPath
}

func makeBranch(arr [][]int64, index int, path []int64, pathes [][]int64) ([]int64, [][]int64) {
	path = append(path, arr[index][0])
	if len(arr[index]) > 1 {
		for j := 1; j <= len(arr[index])-1; j++ {
			nextIndex := getIndex(arr, arr[index][j], index+1)
			if nextIndex != -1 {
				path, pathes = makeBranch(arr, nextIndex, path, pathes)
				// temp := make([]int64, len(path))
				// copy(temp, path)
				// pathes = append(pathes, temp)
				// path = path[0 : len(path)-1]
			} else {
				temp := make([]int64, len(path))
				copy(temp, path)
				pathes = append(pathes, temp)
				path = path[0 : len(path)-1]

				return path, pathes
			}

		}

		temp := make([]int64, len(path))
		copy(temp, path)
		pathes = append(pathes, temp)
		path = path[0 : len(path)-1]

		return path, pathes
	}
	temp := make([]int64, len(path))
	copy(temp, path)
	pathes = append(pathes, temp)
	path = path[0 : len(path)-1]

	return path, pathes
}

func getIndex(arr [][]int64, elem int64, index int) int {

	for j := index; j <= len(arr)-1; j++ {
		if arr[j][0] == elem {
			return j
		}
	}
	return -1
}

// func makeDAG(arr []int64, diff int64) [][]int64 {

// 	var lines [][]int64

// 	for i := 0; i <= len(arr)-1; i++ {
// 		var line []int64
// 		line = append(line, arr[i])
// 		for j := i + 1; j <= len(arr)-1; j++ {
// 			if arr[j]-arr[i] >= diff || arr[j]-arr[i] <= (diff*-1) {
// 				line = append(line, arr[j])
// 			}
// 		}
// 		lines = append(lines, line)
// 	}
// 	return lines
// }

// func addNode(n node) {
// 	n.nodes = append(n.nodes, n)
// }

func computeDays(arr [][]int64, path []int64, days []int64, index int) {
	for j := 0; j <= len(arr)-1; j++ {
		path = append(path, arr[j][0])
		for k := 1; k <= len(arr[j])-1; k++ {
			path = append(path, arr[j][k])
			for l := 0; l <= len(arr)-1; l++ {
				if arr[j][l] == arr[j][0] {

				}
			}
		}
	}
}

// func computePath(arr [][]int64, path []int64, days []int64) {
// 		for i:=0; i<=
// }

func constructDAG(arr []int64, diff int64) [][]int64 {
	var data [][]int64
	for i := 0; i <= len(arr)-1; i++ {
		var line []int64
		line = append(line, arr[i])
		for j := i; j <= len(arr)-1; j++ {
			if (arr[i]-arr[j]) >= diff || arr[i]-arr[j] <= (-1*diff) {
				line = append(line, arr[j])
			}
		}
		data = append(data, line)
	}
	return data
}

func readInput() []string {
	file, _ := os.Open("./input.txt")
	var i int64
	reader := bufio.NewReader(file)
	line1, err := reader.ReadString('\n')
	line1 = strings.Trim(line1, "\n")
	noOfInouts, err := strconv.ParseInt(line1, 10, 32)
	input := make([]string, noOfInouts*2)
	for i = 0; i <= (noOfInouts*2)-1; i++ {
		input[i], err = reader.ReadString('\n')
		if err != nil {
			break
		}
	}
	return input
}
func trimNewLine(s string) string {
	return strings.Trim(s, "\n")
}

func convertToIntArray(data []string) []int64 {
	var newData []int64

	for _, v := range data {
		i, err := strconv.ParseInt(trimNewLine(v), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		newData = append(newData, i)
	}
	return newData
}

func addZero(data []int64, path []int64) []int64 {
	for i := 0; i <= len(data)-1; i++ {
		for j := 0; j <= len(path)-1; j++ {
			if data[i] == path[j] {
				data[i] = 0
			}
		}
	}
	return data
}

func removeZeroEntries(data []string) []string {
	var newData []string

	for _, v := range data {
		if v != "0" {
			newData = append(newData, v)
		}
	}
	return newData
}
