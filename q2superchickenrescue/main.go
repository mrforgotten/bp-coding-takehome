package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ChickenSave(chickenRange, roofRange int, chickenPosition []int) int {
	maxSave := 1
	for i := range chickenPosition {
		currentSave := 1
		totalRange := 0
		for j := i + 1; j < chickenRange; j++ {
			if (totalRange + (chickenPosition[j] - chickenPosition[j-1])) >= roofRange {
				break
			}
			currentSave = currentSave + 1
			totalRange += (chickenPosition[j] - chickenPosition[j-1])
		}

		if currentSave > maxSave {
			maxSave = currentSave
		}
	}

	return maxSave
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	newStr := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(newStr[0])
	k, _ := strconv.Atoi(newStr[1])

	arr := strings.Split(readLine(reader), " ")

	posArr := []int{}

	for _, v := range arr {
		newV, _ := strconv.Atoi(v)
		posArr = append(posArr, newV)
	}

	fmt.Println(ChickenSave(n, k, posArr))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
