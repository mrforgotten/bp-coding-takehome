package main

import "testing"

type testCaseBabyBoss struct {
	expect string
	input  string
}

func TestBabyBossIs(t *testing.T) {
	var testCases = []testCaseBabyBoss{
		{
			expect: GoodBoy,
			input:  "SRSSRRR",
		},
		{
			expect: BadBoy,
			input:  "RSSRR",
		},
		{
			expect: BadBoy,
			input:  "SSSRRRRS",
		},
		{
			expect: BadBoy,
			input:  "SRRSSR",
		},
		{
			expect: GoodBoy,
			input:  "SSRSRR",
		},
	}

	for _, v := range testCases {
		caseResult := BabyBossIs(v.input)
		if caseResult != v.expect {
			t.Errorf("\ncase input: %v \nexpected: %v \ngot: %v", v.input, v.expect, caseResult)
		}

	}
}
