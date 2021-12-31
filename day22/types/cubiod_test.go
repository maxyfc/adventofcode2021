package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCubiodIntersect(t *testing.T) {
	tests := []struct {
		c1, c2   *Cubiod
		expected bool
	}{
		{
			NewCubiod(0, 0, 0, 0, 0, 0),
			NewCubiod(1, 1, 1, 1, 1, 1),
			false,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(0, 1, 0, 1, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(1, 2, 0, 1, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(0, 1, 1, 2, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(0, 1, 0, 1, 1, 2),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(-1, 0, 0, 1, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(0, 1, -1, 0, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(0, 1, 0, 1, -1, 0),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(1, 2, 1, 2, 0, 1),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(1, 2, 0, 1, 1, 2),
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1),
			NewCubiod(1, 2, 1, 2, 1, 2),
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output := test.c1.Intersect(test.c2)
			if test.expected != output {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestCubiodSplit(t *testing.T) {
	tests := []struct {
		c1, c2   *Cubiod
		expected []*Cubiod
		split    bool
	}{
		{
			NewCubiod(0, 0, 0, 0, 0, 0), NewCubiod(1, 1, 1, 1, 1, 1),
			[]*Cubiod{
				NewCubiod(0, 0, 0, 0, 0, 0),
			},
			false,
		},
		{
			NewCubiod(0, 2, 0, 2, 0, 2), NewCubiod(1, 1, 1, 1, 1, 1),
			[]*Cubiod{
				NewCubiod(0, 0, 0, 2, 0, 2),
				NewCubiod(2, 2, 0, 2, 0, 2),
				NewCubiod(1, 1, 0, 0, 0, 2),
				NewCubiod(1, 1, 2, 2, 0, 2),
				NewCubiod(1, 1, 1, 1, 0, 0),
				NewCubiod(1, 1, 1, 1, 1, 1),
				NewCubiod(1, 1, 1, 1, 2, 2),
			},
			true,
		},
		{
			NewCubiod(0, 1, 0, 1, 0, 1), NewCubiod(1, 2, 1, 2, 1, 2),
			[]*Cubiod{
				NewCubiod(0, 0, 0, 1, 0, 1),
				NewCubiod(1, 1, 0, 0, 0, 1),
				NewCubiod(1, 1, 1, 1, 0, 0),
				NewCubiod(1, 1, 1, 1, 1, 1),
			},
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output, split := test.c1.Split(test.c2)
			if test.split != split {
				t.Errorf("c1: %s c2: %s Expected: %v Got: %v",
					test.c1, test.c2, test.split, split)
			}
			if len(test.expected) != len(output) {
				t.Errorf("c1: %s c2: %s Expected length: %v Got: %v",
					test.c1, test.c2, len(test.expected), len(output))
			}
			assertContains(t, test.expected, output)
		})
	}
}

func assertContains(t *testing.T, c1, c2 []*Cubiod) {
	t.Helper()
	for _, c := range c1 {
		found := false
		for _, d := range c2 {
			if reflect.DeepEqual(c, d) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Missing: %s Expected: %s, Got: %s", c, c1, c2)
			return
		}
	}
}

func TestCubiodSplitByRange(t *testing.T) {
	tests := []struct {
		c        *Cubiod
		a        Axis
		r        Range
		expected []*Cubiod
		split    bool
	}{
		{
			NewCubiod(0, 2, 0, 2, 0, 2),
			AxisX,
			NewRange(3, 3),
			[]*Cubiod{
				NewCubiod(0, 2, 0, 2, 0, 2),
			},
			false,
		},
		{
			NewCubiod(0, 2, 0, 2, 0, 2),
			AxisX,
			NewRange(1, 1),
			[]*Cubiod{
				NewCubiod(0, 0, 0, 2, 0, 2),
				NewCubiod(1, 1, 0, 2, 0, 2),
				NewCubiod(2, 2, 0, 2, 0, 2),
			},
			true,
		},
		{
			NewCubiod(0, 2, 0, 2, 0, 2),
			AxisY,
			NewRange(1, 1),
			[]*Cubiod{
				NewCubiod(0, 2, 0, 0, 0, 2),
				NewCubiod(0, 2, 1, 1, 0, 2),
				NewCubiod(0, 2, 2, 2, 0, 2),
			},
			true,
		},
		{
			NewCubiod(0, 2, 0, 2, 0, 2),
			AxisZ,
			NewRange(1, 1),
			[]*Cubiod{
				NewCubiod(0, 2, 0, 2, 0, 0),
				NewCubiod(0, 2, 0, 2, 1, 1),
				NewCubiod(0, 2, 0, 2, 2, 2),
			},
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output, split := test.c.SplitByRange(test.a, test.r)
			if test.split != split {
				t.Errorf("c: %s a: %s r: %s Expected: %v Got: %v",
					test.c, test.a, test.r, test.split, split)
			}
			if !reflect.DeepEqual(test.expected, output) {
				t.Errorf("c: %s a: %s r: %s Expected: %v Got: %v",
					test.c, test.a, test.r, test.expected, output)
			}
		})
	}
}
