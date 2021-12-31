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

func (r1 Range) Split(r2 Range) ([]Range, bool) {
	if !r1.Intersect(r2) {
		return []Range{r1}, false
	}

	var s []Range
	if r1.v1 < r2.v1 {
		s = append(s, NewRange(r1.v1, r2.v1-1), NewRange(r2.v1, min(r1.v2, r2.v2)))
		if r1.v2 > r2.v2 {
			s = append(s, NewRange(r2.v2+1, r1.v2))
		}
	} else {
		s = append(s, NewRange(r1.v1, min(r1.v2, r2.v2)))
		if r1.v2 > r2.v2 {
			s = append(s, NewRange(r2.v2+1, r1.v2))
		}
	}
	return s, true
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.v1, r.v2)
}
