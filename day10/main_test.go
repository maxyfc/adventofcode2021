package main

import (
	util "adventofcode2021/pkg/testutil"
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 26397},
		{Desc: "Part 2", PartFunc: part2, Expected: 288957},
	})
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		line         string
		isValid      bool
		expectedChar rune
		gotChar      rune
		complete     string
	}{
		{"", true, 0, 0, ""},
		{"(", true, 0, 0, ")"},
		{"[", true, 0, 0, "]"},
		{"{", true, 0, 0, "}"},
		{"<", true, 0, 0, ">"},
		{"<<", true, 0, 0, ">>"},
		{"([{<", true, 0, 0, ">}])"},
		{"()", true, 0, 0, ""},
		{"[]", true, 0, 0, ""},
		{"{}", true, 0, 0, ""},
		{"<>", true, 0, 0, ""},
		{"(()", true, 0, 0, ")"},
		{"(<>", true, 0, 0, ")"},
		{"(<>{}", true, 0, 0, ")"},
		{"((>", false, ')', '>', ""},
		{"{([(<{}[<>[]}>{[]{[(<()>", false, ']', '}', ""},
		{"[[<[([]))<([[{}[[()]]]", false, ']', ')', ""},
		{"[{[{({}]{}}([{[{{{}}([]", false, ')', ']', ""},
		{"[<(<(<(<{}))><([]([]()", false, '>', ')', ""},
		{"<{([([[(<>()){}]>(<<{{", false, ']', '>', ""},
		{"[({(<(())[]>[[{[]{<()<>>", true, 0, 0, "}}]])})]"},
		{"[(()[<>])]({[<{<<[]>>(", true, 0, 0, ")}>]})"},
		{"(((({<>}<{<{<>}{[]{[]{}", true, 0, 0, "}}>}>))))"},
		{"{<[[]]>}<{[{[{[]{()[[[]", true, 0, 0, "]]}}]}]}>"},
		{"<{([{{}}[<[[[<>{}]]]>[]]", true, 0, 0, "])}>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Line:%s", test.line), func(t *testing.T) {
			isValid, expectedChar, gotChar, complete := parseLine(test.line)
			if isValid != test.isValid {
				t.Errorf("Expected isValid: %v Got: %v", test.isValid, isValid)
			}
			if expectedChar != test.expectedChar {
				t.Errorf("Expected expectedChar: %v Got: %v", test.expectedChar, expectedChar)
			}
			if gotChar != test.gotChar {
				t.Errorf("Expected gotChar: %v Got: %v", test.gotChar, gotChar)
			}
			if complete != test.complete {
				t.Errorf("Expected complete: %v Got: %v", test.complete, complete)
			}
		})
	}
}

func TestScoreComplete(t *testing.T) {
	tests := []struct {
		complete string
		score    int
	}{
		// {"}}}]])})]", 288957},
		// {")}>]})", 5566},
		// {"}}>}>))))", 1480781},
		// {"]]}}]}]}>", 995444},
		{"])}>", 294},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Complete:%s", test.complete), func(t *testing.T) {
			score := scoreComplete(test.complete)
			if score != test.score {
				t.Errorf("Expected isValid: %v Got: %v", test.score, score)
			}
		})
	}
}
