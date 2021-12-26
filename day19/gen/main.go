package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
	"time"

	"gonum.org/v1/gonum/mat"
)

var (
	pkgName  *string = flag.String("pkgName", "main", "Name of the package of the Point type")
	typeName *string = flag.String("typeName", "Point", "Name of the point type that have the X, Y, Z public fields")
	output   *string = flag.String("output", "rotate.go", "Output file name")
)

func main() {
	flag.Parse()
	writeFile()
	formatFile()
}

func writeFile() {
	out, err := os.Create(*output)
	if err != nil {
		log.Fatalf("Error creating file %s: %s", *output, err)
	}
	defer out.Close()

	rot := genRotations()
	sort.Slice(rot, func(i, j int) bool {
		return compareSlices(rot[i][:], rot[j][:]) > 0
	})

	fmt.Fprintf(out, "// Generated on: %s\n\n", time.Now().Format(time.RFC1123))
	fmt.Fprintf(out, "package %s\n\n", *pkgName)
	fmt.Fprint(out, "import \"fmt\"\n\n")
	fmt.Fprintf(out, "func (p %[1]s) Rotate(id int) %[1]s {\n", *typeName)
	fmt.Fprintf(out, "switch id {\n")
	for i, m := range rot {
		fmt.Fprintf(out, "case %d:\n", i)

		mx := m[0:3]
		my := m[3:6]
		mz := m[6:9]
		printMatrix(out, mx, my, mz)

		fmt.Fprint(out, "return Point{")
		for c, row := range [][]int{mx, my, mz} {
			if c > 0 {
				fmt.Fprint(out, ",")
			}

			s := ""
			if row[0] == -1 {
				s = "-p.X"
			} else if row[0] == 1 {
				s = "p.X"
			} else if row[1] == -1 {
				s = "-p.Y"
			} else if row[1] == 1 {
				s = "p.Y"
			} else if row[2] == -1 {
				s = "-p.Z"
			} else if row[2] == 1 {
				s = "p.Z"
			}
			fmt.Fprint(out, s)
		}
		fmt.Fprintf(out, "}\n")
	}
	fmt.Fprintf(out, "default:\npanic(fmt.Sprintf(\"Invalid rotation ID: %%d\", id))\n}\n}\n")
}

func formatFile() {
	cmd := exec.Command("gofmt", "-w", *output)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error formatting file %s: %s", *output, err)
	}
}

func printMatrix(w io.Writer, x, y, z []int) {
	fmt.Fprintf(w, "// [ %2d %2d %2d ] [ X ]\n", x[0], x[1], x[2])
	fmt.Fprintf(w, "// [ %2d %2d %2d ] [ Y ]\n", y[0], y[1], y[2])
	fmt.Fprintf(w, "// [ %2d %2d %2d ] [ Z ]\n", z[0], z[1], z[2])
}

func compareSlices(a, b []int) int {
	for {
		if len(a) == 0 && len(b) == 0 {
			return 0
		}
		if len(a) == 0 && len(b) > 0 {
			return -1
		}
		if len(b) == 0 && len(a) > 0 {
			return 1
		}
		if a[0] == b[0] {
			a = a[1:]
			b = b[1:]
			continue
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	}
}

func genRotations() [][9]int {
	ms := make(map[[9]int]struct{})

	// For each axis we compute it's 90 degree rotations
	// It doesn't matter if the we rotated X axis first then Y first
	// as the rotations are commutative
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				mx := rotateX(math.Pi / 2 * float64(x))
				my := rotateY(math.Pi / 2 * float64(y))
				mz := rotateZ(math.Pi / 2 * float64(z))

				var m mat.Dense
				m.Mul(mx, my)
				m.Mul(&m, mz)

				// Round values that are very close to zero to zero
				nr, nc := m.Dims()
				for r := 0; r < nr; r++ {
					for c := 0; c < nc; c++ {
						val := m.At(r, c)
						if math.Abs(val) < 1e-6 {
							m.Set(r, c, 0)
						}
					}
				}

				// Convert the matrix into a slice so that we can distinct them
				// Some combination of rotations end of with the same matrix
				var ma [9]int
				for i, val := range m.RawMatrix().Data {
					ma[i] = int(val)
				}

				ms[ma] = struct{}{}
			}
		}
	}

	result := make([][9]int, 0, len(ms))
	for m := range ms {
		result = append(result, m)
	}
	return result
}

// Basic roation matrix from wikipedia
// https://en.wikipedia.org/wiki/Rotation_matrix#Basic_rotations
func rotateX(rad float64) mat.Matrix {
	return mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, math.Cos(rad), -math.Sin(rad),
		0, math.Sin(rad), math.Cos(rad),
	})
}

func rotateY(rad float64) mat.Matrix {
	return mat.NewDense(3, 3, []float64{
		math.Cos(rad), 0, math.Sin(rad),
		0, 1, 0,
		-math.Sin(rad), 0, math.Cos(rad),
	})
}

func rotateZ(rad float64) mat.Matrix {
	return mat.NewDense(3, 3, []float64{
		math.Cos(rad), -math.Sin(rad), 0,
		math.Sin(rad), math.Cos(rad), 0,
		0, 0, 1,
	})
}
