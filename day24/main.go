/*
Day 24 ended up being a problem that cannot be solved be using brute force.
The search space is just too large (14^9 values need to be searched).
After looking in reddit for hints and other people's solution, this is the
solution used here:

1.    Analyse the problem input. Similar to what other people had found out. The
    input is divided into 14 similar sections. Each section only had 3
    differences as show below:

    inp w
    mul x 0
    add x z
    mod x 26
    div z <<zdiv>> - This is either "div z 1" or "div z 26"
    add x <<xadd>> - xadd is a random integer value
    eql x w
    eql x 0
    mul y 0
    add y 25
    mul y x
    add y 1
    mul z y
    mul y 0
    add y w
    add y <<yadd>> - yadd is a random integer value
    mul y x
    add z y

    This ends up with following 14 set of values:

    1:  {zdiv:1,  xadd:11,  yadd:1}
    2:  {zdiv:1,  xadd:10,  yadd:10}
    3:  {zdiv:1,  xadd:13,  yadd:2}
    4:  {zdiv:26, xadd:-10, yadd:5}
    5:  {zdiv:1,  xadd:11,  yadd:6}
    6:  {zdiv:1,  xadd:11,  yadd:0}
    7:  {zdiv:1,  xadd:12,  yadd:16}
    8:  {zdiv:26, xadd:-11, yadd:12}
    9:  {zdiv:26, xadd:-7,  yadd:15}
    10: {zdiv:1,  xadd:13,  yadd:7}
    11: {zdiv:26, xadd:-13, yadd:6}
    12: {zdiv:26, xadd:0,   yadd:5}
    13: {zdiv:26, xadd:-11, yadd:6}
    14: {zdiv:26, xadd:0,   yadd:15}

    This is extracted by the extract(string) function.

2.  Next is simplifying the logic in each section. It basically boils down to:

    func section(w, prevz, zdiv, xadd, yadd int) int {
        x1 := prevz%26 + xadd
        var x2 int
        if x1 == w {
            x2 = 0
        } else {
            x2 = 1
        }

        z1 := prevz / zdiv * (25*x2 + 1)

        newz := z1 + x2*(w+yadd)
        return newz
    }

    The key insights to the simplified version of the section is that
    x2 can only have 2 values (0 or 1).

    When x2 is 0, then the calculation is just:

        z1 = prevz / zdiv * (25*0 + 1)
           = prevz / zdiv * 1
           = prevz / zdiv

        newz = z1 + 0*(w+yadd)
             = z1
             = prevz / zdiv

        newz := prevz / zdiv

    When x2 is 1, then the calculation is just:

        z1 = prevz / zdiv * (25*1 + 1)
           = prevz / zdiv * 26

        newz = z1 + 1*(w+yadd)
             = z1 + (w+yadd)
             = z1 + w + yadd

        newz := prevz / zdiv * 26 + w + yadd

    Looking at the input data when zdiv is 1 then x2 is always 1 because
    xadd will increase the value to beyond 0:

        x1 := prevz%26 + xadd
              |------|         - this will end with a value betwen 0 and 25
                         |--|  - the input data always have a value more than 9
                                 this will make x1 always more than 9

        var x2 int
        if x1 == w {           - w is between 0 and 9 and given x2 is always
            x2 = 0               more than 9, the condition will be false
        } else {               - hence, x2 will always be 1 when zdiv is 1
            x2 = 1
        }

    So when zdiv is 1:

        newz = prevz / zdiv * 26 + w + yadd
             = prevz / 1 * 26 + w + yadd
             = prevz * 26 + w + yadd

        newz := prevz * 26 + w + yadd

    Looking at the alternative scenario of zdiv being 26 (the input only has
    1 or 26 in the input data). There is no way to simplify the logic like
    above. However, there is hint from people's solution in reddit: z is used
    like a stack of base 26 numbers. Each time it's multiples the previous z
    with 26 it is pushing the previous z value down and putting the new value
    on the top of the stack:

        newz := prevz * 26 + w + yadd <-- when zdiv = 1
                |--------|              - pushing the previous value down the
                                          stack
                             |------|   - putting the new value on the top of
                                          the stack

    When the zdiv is 26, this is popping from the stack:

        newz := prevz / zdiv <-- when zdiv = 26 and x2 = 0

    Looking at the problem input, there is seven zdiv = 1 and seven zdiv = 26.
    If we can find a value that pop and push seven times, we end up with the
    original value for z which is zero (prevz is 0 for the first digit). And,
    this is what we need for the a valid model number.

    Step Summary:
        PUSH:    newz := prevz * 26 + w + yadd <-- when zdiv = 1
        POP:    newz := prevz / zdiv          <-- when zdiv = 26 and x2 = 0

        To make x2 = 0: prevz%26 + xadd == w

3.  The next step is analyse each section with the input data to get a sequence
    of new z values. During analysis, we pair each pop with a push and then
    try and find the constraint that will make x2 = 0.

    i1 .. i14 - represents the each of the 14 digits in the model number


    Input    New Z                            Constraint
    ----------------------------------------------------------------------------
    1 PUSH    Workout:
            newz = prevz * 26 + w + yadd
                 = 0 * 26 + i1 + 1
                 = i1 + 1

            New Z:
            i1 + 1

    2 PUSH    _Worked out like above._

            26 * prevz + i2 + 10

    3 PUSH    26 * prevz + i3 + 2

    4 POP    prevz / 26                        Workout:
                                            prevz % 26 + xadd = w
                                            prevz % 26 + -10  = i4
                                            i3 + 2 - 10       = i4
                                            |----|    This is a key insight here
                                                    prevz%26 is just PUSH
                                                    without the 26 * prevz
                                            i3 - 8            = i4

                                            This tells the third and fourth
                                            digit are related in this way:
                                            i3 - 8 = i4

    5 PUSH    26 * prevz + i5 + 6

    6 PUSH    26 * prevz + i6

    7 PUSH    26 * prevz + i7 + 16

    8 POP    prevz / 26                        This pairs with input #7.
                                            Worked like like above:
                                            i7 + 5 = i8

    9 POP    prevz / 26                        Pair input: #6
                                            i6 - 7 = i9

    10 PUSH 26 * prevz + i10 +7

    11 POP    prevz / 26                        Pair input: #10
                                            i10 - 6 = i11

    12 POP    prevz / 26                        Pair input: #5
                                            i5 + 6 = i12

    13 POP    prevz / 26                        Pair input: #2
                                            i2 - 1 = i13

    14 POP    prevz / 26                        Pair input: #1
                                            i1 + 1 = i14

    Step Summary (constraints):
        i3 - 8 = i4
        i7 + 5 = i8
        i6 - 7 = i9
        i10 - 6 = i11
        i5 + 6 = i12
        i2 - 1 = i13
        i1 + 1 = i14

    Note that my constraints may be difference from everyone else. I suspect
    each person may be getting a different problem input. My suspicous is
    confined when I worked a solution that is accepted using the above
    constraints.

4.    The final step is working out what is maximum value that I for each digit
    starting from the first digit because we can to find the largest number.

        i1 + 1 = i14

    The maximum value for i1 is 8 because any more will make i14 exceed 9.

        i1 = 8
        i14 = 9

    Filling the table:

    |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |  9 | 10 | 11 | 12 | 13 | 14 |
    -----------------------------------------------------------------------
    |  8 |    |    |    |    |    |    |    |    |    |    |    |    |  9 |

    Filling the rest of the table:

    |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |  9 | 10 | 11 | 12 | 13 | 14 |
    -----------------------------------------------------------------------
    |  8 |  9 |  9 |  1 |  3 |  9 |  4 |  9 |  2 |  9 |  3 |  9 |  8 |  9 |

    For part 2, we need to find the minmum value. Using the similar logic as
    above:

        i1 + 1 = i14 --> i1 = 1  i14 = 2

    Filling the table:

    |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |  9 | 10 | 11 | 12 | 13 | 14 |
    -----------------------------------------------------------------------
    |  1 |  2 |  9 |  1 |  1 |  8 |  1 |  6 |  1 |  7 |  1 |  7 |  1 |  2 |

    Step summary:
        Maximum:    89913949293989
        Minimum:    12911816171712
*/
package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	return find(input, true)
}

func part2(input string) int {
	return find(input, false)
}

type constraint struct {
	digit1 int
	digit2 int
	inc    int // digit1 = digit2 + inc
}

func find(input string, maximum bool) int {
	// Find the push and pop pairs
	var stack []data
	var constraints []constraint
	for _, d := range extract(inputData) {
		switch d.zdiv {
		case 1: // PUSH
			stack = append(stack, d)
		case 26: // POP
			var pop data
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			constraints = append(constraints, constraint{
				digit1: d.digit,
				digit2: pop.digit,
				inc:    d.xadd + pop.yadd,
			})
		default:
			panic(fmt.Sprintf("Unexpected zdiv: %d", d.zdiv))
		}
	}

	// Resolve constraints
	var digits [14]int
	for _, c := range constraints {
		if maximum { // Find the maximum
			if c.inc < 0 {
				digits[c.digit2] = 9
				digits[c.digit1] = 9 + c.inc
			} else {
				digits[c.digit2] = 9 - c.inc
				digits[c.digit1] = 9
			}
		} else { // Find the minimum
			if c.inc < 0 {
				digits[c.digit2] = 1 + -c.inc
				digits[c.digit1] = 1
			} else {
				digits[c.digit2] = 1
				digits[c.digit1] = 1 + c.inc
			}
		}
	}

	// Convert array of digits to int
	value := 0
	for _, d := range digits {
		value = value*10 + d
	}
	return value
}

var extractRE *regexp.Regexp = regexp.MustCompile(`inp w\s+mul x 0\s+add x z\s+mod x 26\s+div z (-?\d+)\s+add x (-?\d+)\s+eql x w\s+eql x 0\s+mul y 0\s+add y 25\s+mul y x\s+add y 1\s+mul z y\s+mul y 0\s+add y w\s+add y (-?\d+)\s+mul y x\s+add z y`)

type data struct {
	digit int
	zdiv  int
	xadd  int
	yadd  int
}

// extract pulls out the zdiv, xadd and yadd values
func extract(input string) []data {
	var d []data
	for i, m := range extractRE.FindAllStringSubmatch(input, -1) {
		d = append(d, data{
			digit: i,
			zdiv:  strutil.MustAtoi(m[1]),
			xadd:  strutil.MustAtoi(m[2]),
			yadd:  strutil.MustAtoi(m[3]),
		})
	}
	return d
}
