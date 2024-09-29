package util_test

import (
	"strconv"
	"testing"
	"txboardcast/service/util"
)

type expectation struct {
	value bool
	err   *error
}

type validateLongIntStringTestCase struct {
	input  string
	expect expectation
}

// util.MAX_ALLOW_STRING_NUMBER should match this or vice-versa
const MAX_ALLOW_STRING_NUMBER = "340282366920938463463374607431768211455"

func TestValidateLongIntString(t *testing.T) {
	testCases := []validateLongIntStringTestCase{
		{
			input: "5",
			expect: expectation{
				value: true,
			},
		},
		{
			input: "123456",
			expect: expectation{
				value: true,
			},
		},
		{
			input: MAX_ALLOW_STRING_NUMBER,
			expect: expectation{
				value: true,
			},
		},
		{
			input: "3402823669209384634633746074317682114551",
			expect: expectation{
				err: &util.ErrorTooBigInput,
			},
		},
		{
			input: func() string {
				var max = MAX_ALLOW_STRING_NUMBER

				newMax := ""

				// loop from last digit
				i := len(max) - 1
				for ; i > 0; i-- {
					currInt, _ := strconv.Atoi(string(max[i]))
					currInt++
					newMax = strconv.Itoa(currInt%10) + newMax // append to front
					if currInt < 10 {
						newMax = max[:i] + newMax
						break
					}
				}

				return newMax
			}(),
			expect: expectation{
				err: &util.ErrorTooBigInput,
			},
		},
		{
			input: "",
			expect: expectation{
				err: &util.ErrorRequireInput,
			},
		},
		{
			input: "a",
			expect: expectation{
				err: &util.ErrorInvalidInput,
			},
		},
		{
			input: "11111A",
			expect: expectation{
				err: &util.ErrorInvalidInput,
			},
		},
		{
			input: "-11111",
			expect: expectation{
				err: &util.ErrorInvalidInput,
			},
		},
	}

	for _, testCase := range testCases {
		isFail := false
		result, err := util.ValidateLongIntString(testCase.input)

		if result != testCase.expect.value {
			isFail = true
			t.Errorf("Expecting value: %v\ngot: %v\n", testCase.expect.value, result)
		}
		if testCase.expect.err != nil && err != *testCase.expect.err {
			isFail = true
			t.Errorf("Expecting error: %v\ngot: %v\n", (*testCase.expect.err).Error(), err)
		}
		if isFail {
			t.Errorf("input: %v\n", testCase.input)
		}
	}
}
