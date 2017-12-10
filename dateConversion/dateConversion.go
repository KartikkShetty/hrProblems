package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	date, _ := reader.ReadString('\n')
	ampm := date[8:len(date)]
	ampm = strings.Trim(ampm, "\n")

	hour := date[0:2]
	remaining := date[3:8]
	if ampm == "AM" {
		if hour == "12" {
			hour = "00"
		}
		fmt.Printf("%v:%v\n", hour, remaining)
	} else {
		h, _ := strconv.ParseInt(hour, 10, 32)
		if h != 12 {
			h = h + 12
		}

		fmt.Printf("%v:%v\n", h, remaining)
	}

}
