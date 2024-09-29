package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const GoodBoy = "Good boy"
const BadBoy = "Bad boy"

// condition no first R, no end S, contain R
func BabyBossIs(shotEvent string) string {

	if string(shotEvent[0]) == "R" || string(shotEvent[len(shotEvent)-1]) == "S" {

		return BadBoy
	}

	shotCount := 0
	for _, event := range shotEvent {
		if event == 'S' {
			shotCount++
		}
		if event == 'R' && shotCount > 0 {
			shotCount--
		}
	}

	if shotCount > 0 {
		return BadBoy
	}

	return GoodBoy
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	eventShot := readLine(reader)

	fmt.Println(BabyBossIs(eventShot))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
