package main

import (
	util "adventofcode2021/pkg/testutil"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 26},
		{Desc: "Part 2", PartFunc: part2, Expected: 61229},
	})
}

func TestDecode(t *testing.T) {
	expected := map[string]int{
		"acedgfb": 8,
		"cdfbe":   5,
		"gcdfa":   2,
		"fbcad":   3,
		"dab":     7,
		"cefabd":  9,
		"cdfgeb":  6,
		"eafb":    4,
		"cagedb":  0,
		"ab":      1,
	}

	var signals []string
	for s := range expected {
		signals = append(signals, s)
	}

	output := decode(signals)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected: %v Got: %v", expected, output)
	}
}
