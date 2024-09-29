package main

import "testing"

type testCaseChickenSave struct {
	expect int
	input  ChickenSaveInput
}

type ChickenSaveInput struct {
	N   int
	K   int
	Pos []int
}

func TestChickenSave(t *testing.T) {
	testCases := []testCaseChickenSave{
		// test case from question
		{
			expect: 2,
			input: ChickenSaveInput{
				N:   5,
				K:   5,
				Pos: []int{2, 5, 10, 12, 15},
			},
		},
		// test case from question
		{
			expect: 4,
			input: ChickenSaveInput{
				N:   6,
				K:   10,
				Pos: []int{1, 11, 30, 34, 35, 37},
			},
		},
		{
			expect: 4,
			input: ChickenSaveInput{
				N:   5,
				K:   10,
				Pos: []int{1, 3, 6, 10, 14},
			},
		},
		{
			expect: 3,
			input: ChickenSaveInput{
				N:   5,
				K:   9,                      // 123456789XABCD
				Pos: []int{1, 3, 6, 10, 14}, // x_x__x___x___x
			},
		},
		{
			expect: 1,
			input: ChickenSaveInput{
				N:   1,
				K:   5,
				Pos: []int{2},
			},
		},
		{
			expect: 1,
			input: ChickenSaveInput{
				N:   3,
				K:   2,              // 123456789
				Pos: []int{3, 5, 8}, // __x_x__x_
			},
		},
		{
			expect: 4,
			input: ChickenSaveInput{
				N:   4,
				K:   20,
				Pos: []int{1, 3, 6, 10},
			},
		},
		{
			expect: 4,
			input: ChickenSaveInput{
				N:   4,
				K:   100,
				Pos: []int{1, 2, 3, 4},
			},
		},
		{
			expect: 1,
			input: ChickenSaveInput{
				N:   4,
				K:   5,
				Pos: []int{1, 10, 20, 30},
			},
		},
	}

	for _, testCase := range testCases {
		caseResult := ChickenSave(testCase.input.N, testCase.input.K, testCase.input.Pos)

		if testCase.expect != caseResult {
			t.Errorf("\ninput: %v\nexpected: %v \ngot: %v", testCase.input, testCase.expect, caseResult)
		}
	}
}
