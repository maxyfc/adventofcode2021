package types

import "fmt"

type Axis int

const (
	AxisX Axis = iota
	AxisY
	AxisZ
)

func (a Axis) String() string {
	switch a {
	case AxisX:
		return "X"
	case AxisY:
		return "Y"
	case AxisZ:
		return "Z"
	default:
		panic(fmt.Sprintf("Invalid axis: %d", a))
	}
}

type Cubiod struct {
	x, y, z Range
}

func NewCubiod(x1, x2, y1, y2, z1, z2 int) *Cubiod {
	return &Cubiod{
		x: NewRange(x1, x2),
		y: NewRange(y1, y2),
		z: NewRange(z1, z2),
	}
}

func (c1 *Cubiod) Intersect(c2 *Cubiod) bool {
	return c1.x.Intersect(c2.x) && c1.y.Intersect(c2.y) && c1.z.Intersect(c2.z)
}

func (c1 *Cubiod) Split(c2 *Cubiod) ([]*Cubiod, bool) {
	cs := []*Cubiod{c1}
	if !c1.Intersect(c2) {
		return cs, false
	}

	for _, axis := range []struct {
		a Axis
		r Range
	}{
		{AxisX, c2.x},
		{AxisY, c2.y},
		{AxisZ, c2.z},
	} {
		var newcs []*Cubiod
		for _, c := range cs {
			if !c.Intersect(c2) {
				newcs = append(newcs, c)
				continue
			}
			output, _ := c.SplitByRange(axis.a, axis.r)
			newcs = append(newcs, output...)
		}
		cs = newcs[:]
	}
	return cs, true
}

func (c *Cubiod) SplitByRange(a Axis, r Range) ([]*Cubiod, bool) {
	var cr Range
	switch a {
	case AxisX:
		cr = c.x
	case AxisY:
		cr = c.y
	case AxisZ:
		cr = c.z
	default:
		panic(fmt.Sprintf("Invalid axis: %d", a))
	}

	rs, split := cr.Split(r)
	if !split {
		return []*Cubiod{c}, false
	}

	var result []*Cubiod
	for _, nr := range rs {
		cc := c.Copy()
		switch a {
		case AxisX:
			cc.x = nr
		case AxisY:
			cc.y = nr
		case AxisZ:
			cc.z = nr
		default:
			panic(fmt.Sprintf("Invalid axis: %d", a))
		}
		result = append(result, cc)
	}
	return result, true
}

func (c *Cubiod) Copy() *Cubiod {
	return &Cubiod{x: c.x, y: c.y, z: c.z}
}

func (c *Cubiod) String() string {
	return fmt.Sprintf("(x=%s y=%s z=%s)", c.x, c.y, c.z)
}
