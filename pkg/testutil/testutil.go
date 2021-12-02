package testutil

import "testing"

type TestCase struct {
	Desc     string
	PartFunc func(string) int
	Expected int
}

func RunTests(t *testing.T, testData string, tests []TestCase) {
	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			output := test.PartFunc(testData)
			if output != test.Expected {
				t.Errorf("Expected output: %d Got: %d", test.Expected, output)
			}
		})
	}
}
