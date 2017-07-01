package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFileLineByLine()

	can := strings.Split(lines[1], " ")
	output := 1
	maxValue, _ := strconv.ParseInt(can[0], 10, 64)

	var i int64

	for i = 1; i <= int64(len(can))-1; i++ {
		value, _ := strconv.ParseInt(can[i], 10, 64)
		if maxValue == value {
			output = output + 1
		} else if maxValue < value {
			maxValue, _ = strconv.ParseInt(can[i], 10, 64)
			output = 1
		}

	}
	fmt.Println(output)
}

func readFileLineByLine() [2]string {
	var input [2]string
	var err error

	reader := bufio.NewReader(os.Stdin)
	input[0], err = reader.ReadString('\n')
	for {
		input[1], err = reader.ReadString('\n')
		if err != nil {
			break
		}
	}
	return input
}
