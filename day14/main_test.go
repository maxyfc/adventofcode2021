package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	tests := []struct {
		partFunc func(string) uint64
		expected uint64
	}{
		{part1, 1588},
		{part2, 2188189693529},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(testData)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		template map[[2]byte]uint64
		rules    map[[2]byte]byte
		count    map[byte]uint64
		expected map[[2]byte]uint64
		newCount map[byte]uint64
	}{
		{
			// NN
			map[[2]byte]uint64{
				{'N', 'N'}: 1,
			},
			map[[2]byte]byte{
				{'N', 'N'}: 'C',
			},
			map[byte]uint64{
				'N': 2,
			},
			// NCN
			map[[2]byte]uint64{
				{'N', 'C'}: 1,
				{'C', 'N'}: 1,
			},
			map[byte]uint64{
				'N': 2,
				'C': 1,
			},
		},
		{
			// NCN
			map[[2]byte]uint64{
				{'N', 'N'}: 1,
				{'N', 'C'}: 1,
			},
			map[[2]byte]byte{
				{'N', 'N'}: 'C',
				{'N', 'C'}: 'B',
			},
			map[byte]uint64{
				'N': 2,
				'C': 1,
			},
			// NCNBC
			map[[2]byte]uint64{
				{'N', 'C'}: 1,
				{'C', 'N'}: 1,
				{'N', 'B'}: 1,
				{'B', 'C'}: 1,
			},
			map[byte]uint64{
				'N': 2,
				'C': 2,
				'B': 1,
			},
		},
		{
			// NN
			map[[2]byte]uint64{
				{'N', 'N'}: 1,
			},
			map[[2]byte]byte{
				{'N', 'N'}: 'N',
			},
			map[byte]uint64{
				'N': 2,
			},
			// NNN
			map[[2]byte]uint64{
				{'N', 'N'}: 2,
			},
			map[byte]uint64{
				'N': 3,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output, count := step(test.template, test.rules, test.count)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
			if !reflect.DeepEqual(count, test.newCount) {
				t.Errorf("Expected: %v Got: %v", test.newCount, count)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		template map[[2]byte]uint64
		rules    map[[2]byte]byte
		count    map[byte]uint64
	}{
		{
			`NN
			
NN -> C`,
			map[[2]byte]uint64{
				{'N', 'N'}: 1,
			},
			map[[2]byte]byte{
				{'N', 'N'}: 'C',
			},
			map[byte]uint64{
				'N': 2,
			},
		},
		{
			`NNC
			
NN -> C
NC -> B`,
			map[[2]byte]uint64{
				{'N', 'N'}: 1,
				{'N', 'C'}: 1,
			},
			map[[2]byte]byte{
				{'N', 'N'}: 'C',
				{'N', 'C'}: 'B',
			},
			map[byte]uint64{
				'N': 2,
				'C': 1,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			template, rules, count := parse(test.input)
			if !reflect.DeepEqual(template, test.template) {
				t.Errorf("Expected: %v Got: %v", test.template, template)
			}
			if !reflect.DeepEqual(rules, test.rules) {
				t.Errorf("Expected: %v Got: %v", test.rules, rules)
			}
			if !reflect.DeepEqual(count, test.count) {
				t.Errorf("Expected: %v Got: %v", test.count, count)
			}
		})
	}
}
