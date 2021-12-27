// Generated on: Tue, 28 Dec 2021 10:40:05 AEDT

package point

import "fmt"

func (p Point) Rotate(id int) Point {
	switch id {
	case 0:
		// [  1  0  0 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{p.X, p.Y, p.Z}
	case 1:
		// [  1  0  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{p.X, p.Z, -p.Y}
	case -1:
		// [  1  0  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{p.X, -p.Z, p.Y}
	case 2:
		// [  1  0  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{p.X, -p.Z, p.Y}
	case -2:
		// [  1  0  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{p.X, p.Z, -p.Y}
	case 3:
		// [  1  0  0 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{p.X, -p.Y, -p.Z}
	case -3:
		// [  1  0  0 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{p.X, -p.Y, -p.Z}
	case 4:
		// [  0  1  0 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{p.Y, p.X, -p.Z}
	case -4:
		// [  0  1  0 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{p.Y, p.X, -p.Z}
	case 5:
		// [  0  1  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{p.Y, p.Z, p.X}
	case -5:
		// [  0  0  1 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{p.Z, p.X, p.Y}
	case 6:
		// [  0  1  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{p.Y, -p.Z, -p.X}
	case -6:
		// [  0  0 -1 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{-p.Z, p.X, -p.Y}
	case 7:
		// [  0  1  0 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{p.Y, -p.X, p.Z}
	case -7:
		// [  0 -1  0 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{-p.Y, p.X, p.Z}
	case 8:
		// [  0  0  1 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{p.Z, p.X, p.Y}
	case -8:
		// [  0  1  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{p.Y, p.Z, p.X}
	case 9:
		// [  0  0  1 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{p.Z, p.Y, -p.X}
	case -9:
		// [  0  0 -1 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{-p.Z, p.Y, p.X}
	case 10:
		// [  0  0  1 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{p.Z, -p.Y, p.X}
	case -10:
		// [  0  0  1 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{p.Z, -p.Y, p.X}
	case 11:
		// [  0  0  1 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{p.Z, -p.X, -p.Y}
	case -11:
		// [  0 -1  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{-p.Y, -p.Z, p.X}
	case 12:
		// [  0  0 -1 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{-p.Z, p.X, -p.Y}
	case -12:
		// [  0  1  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{p.Y, -p.Z, -p.X}
	case 13:
		// [  0  0 -1 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{-p.Z, p.Y, p.X}
	case -13:
		// [  0  0  1 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{p.Z, p.Y, -p.X}
	case 14:
		// [  0  0 -1 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{-p.Z, -p.Y, -p.X}
	case -14:
		// [  0  0 -1 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{-p.Z, -p.Y, -p.X}
	case 15:
		// [  0  0 -1 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{-p.Z, -p.X, p.Y}
	case -15:
		// [  0 -1  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{-p.Y, p.Z, -p.X}
	case 16:
		// [  0 -1  0 ] [ X ]
		// [  1  0  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{-p.Y, p.X, p.Z}
	case -16:
		// [  0  1  0 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{p.Y, -p.X, p.Z}
	case 17:
		// [  0 -1  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [ -1  0  0 ] [ Z ]
		return Point{-p.Y, p.Z, -p.X}
	case -17:
		// [  0  0 -1 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{-p.Z, -p.X, p.Y}
	case 18:
		// [  0 -1  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  1  0  0 ] [ Z ]
		return Point{-p.Y, -p.Z, p.X}
	case -18:
		// [  0  0  1 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{p.Z, -p.X, -p.Y}
	case 19:
		// [  0 -1  0 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{-p.Y, -p.X, -p.Z}
	case -19:
		// [  0 -1  0 ] [ X ]
		// [ -1  0  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{-p.Y, -p.X, -p.Z}
	case 20:
		// [ -1  0  0 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{-p.X, p.Y, -p.Z}
	case -20:
		// [ -1  0  0 ] [ X ]
		// [  0  1  0 ] [ Y ]
		// [  0  0 -1 ] [ Z ]
		return Point{-p.X, p.Y, -p.Z}
	case 21:
		// [ -1  0  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{-p.X, p.Z, p.Y}
	case -21:
		// [ -1  0  0 ] [ X ]
		// [  0  0  1 ] [ Y ]
		// [  0  1  0 ] [ Z ]
		return Point{-p.X, p.Z, p.Y}
	case 22:
		// [ -1  0  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{-p.X, -p.Z, -p.Y}
	case -22:
		// [ -1  0  0 ] [ X ]
		// [  0  0 -1 ] [ Y ]
		// [  0 -1  0 ] [ Z ]
		return Point{-p.X, -p.Z, -p.Y}
	case 23:
		// [ -1  0  0 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{-p.X, -p.Y, p.Z}
	case -23:
		// [ -1  0  0 ] [ X ]
		// [  0 -1  0 ] [ Y ]
		// [  0  0  1 ] [ Z ]
		return Point{-p.X, -p.Y, p.Z}
	default:
		panic(fmt.Sprintf("Invalid rotation ID: %d", id))
	}
}
