// Generated on: Fri, 07 Jan 2022 00:38:36 AEDT

package validator

func Validate(in int) bool {
	var w, x, y, z int

	// 13
	w = in / 10000000000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 1
	y *= x
	z += y

	// 12
	w = in / 1000000000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 10
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 10
	y *= x
	z += y

	// 11
	w = in / 100000000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 13
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 2
	y *= x
	z += y

	// 10
	w = in / 10000000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += -10
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 5
	y *= x
	z += y

	// 9
	w = in / 1000000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 6
	y *= x
	z += y

	// 8
	w = in / 100000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 0
	y *= x
	z += y

	// 7
	w = in / 10000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 12
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 16
	y *= x
	z += y

	// 6
	w = in / 1000000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += -11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 12
	y *= x
	z += y

	// 5
	w = in / 100000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += -7
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 15
	y *= x
	z += y

	// 4
	w = in / 10000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 1
	x += 13
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 7
	y *= x
	z += y

	// 3
	w = in / 1000 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += -13
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 6
	y *= x
	z += y

	// 2
	w = in / 100 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += 0
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 5
	y *= x
	z += y

	// 1
	w = in / 10 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += -11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 6
	y *= x
	z += y

	// 0
	w = in / 1 % 10
	if w == 0 {
		return false
	}
	x = z
	x %= 26
	z /= 26
	x += 0
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25
	y *= x
	y += 1
	z *= y
	y = w
	y += 15
	y *= x
	z += y

	return z == 0
}
