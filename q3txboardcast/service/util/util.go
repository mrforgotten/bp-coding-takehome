package util

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const MAX_ALLOW_STRING_NUMBER = "340282366920938463463374607431768211455"
const MAX_LEN_NUMBER = len(MAX_ALLOW_STRING_NUMBER)

var ErrorTooBigInput = fmt.Errorf("input must not bigger than `%v`", MAX_ALLOW_STRING_NUMBER)
var ErrorRequireInput = fmt.Errorf("input is required")
var ErrorInvalidInput = fmt.Errorf("invalid input, must be possitive Integer only")

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func ValidateLongIntString(str string) (bool, error) {

	if str == "" {
		return false, ErrorRequireInput
	}
	if len(str) > MAX_LEN_NUMBER {
		return false, ErrorTooBigInput
	} else if len(str) == MAX_LEN_NUMBER {
		for i := range str {
			if str[i] > MAX_ALLOW_STRING_NUMBER[i] {

				return false, ErrorTooBigInput
			}
		}
	} else {
		for _, v := range str {
			_, err := strconv.Atoi(string(v))

			if err != nil {

				return false, ErrorInvalidInput
			}
		}
	}

	return true, nil
}
