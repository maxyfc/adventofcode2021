package types

import "fmt"

type Range struct {
	v1, v2 int
}

func NewRange(v1, v2 int) Range {
	if v1 > v2 {
		v1, v2 = v2, v1 // Ensures that v1 is the smallest value
	}
	return Range{v1, v2}
}

func (r1 Range) Intersect(r2 Range) bool {
	if r1.v1 > r2.v1 {
		r1, r2 = r2, r1 // Ensure that r1 contains the smallest v1
	}
	return r1.v2 >= r2.v1
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.v1, r.v2)
}
