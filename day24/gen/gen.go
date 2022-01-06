package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	inputFile  *string = flag.String("inputFile", "", "Input file with containing instructions")
	outputFile *string = flag.String("outputFile", "", "Generated output file")
)

func main() {
	flag.Parse()

	if strings.TrimSpace(*inputFile) == "" {
		log.Fatalf("Input file is required.")
	}
	if strings.TrimSpace(*outputFile) == "" {
		log.Fatalf("Output file is required.")
	}

	lines := readInputFile()
	writeOutputFile(lines)
	formatOutputFile()

	log.Printf("File generated: %s", *outputFile)
}

func readInputFile() []string {
	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Unable to open file: %s: %v", *inputFile, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Reading file error: %v", err)
	}
	return lines
}

func writeOutputFile(lines []string) {
	f, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Unable to create file: %s: %v", *outputFile, err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Generated on: %s\n\n", time.Now().Format(time.RFC1123))
	fmt.Fprintln(f, "package validator")
	fmt.Fprintf(f, "func Validate(in int) bool {\n")

	inputIndex := 14
	fmt.Fprintf(f, "var w, x, y, z int\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lookAhead := ""
		if i+1 < len(lines) {
			lookAhead = lines[i+1]
		}

		splits := strings.Split(line, " ")
		lookAheadSplits := strings.Split(lookAhead, " ")

		switch splits[0] {
		case "inp":
			inputIndex--
			if inputIndex < 0 {
				panic("Too many inp")
			}
			fmt.Fprintf(f, "\n// %d\n", inputIndex)
			fmt.Fprintf(f, "%s = in / %d %% 10\n", splits[1], int(math.Pow10(inputIndex)))
			fmt.Fprintf(f, "if %s == 0 { return false }\n", splits[1])
		case "add":
			fmt.Fprintf(f, "%s += %s\n", splits[1], splits[2])
		case "mul":
			if splits[2] == "0" {
				if lookAheadSplits[0] == "add" && lookAheadSplits[1] == splits[1] {
					fmt.Fprintf(f, "%s = %s\n", lookAheadSplits[1], lookAheadSplits[2])
					i++ // Skip next line
				} else {
					fmt.Fprintf(f, "%s = 0\n", splits[1])
				}
			} else {
				fmt.Fprintf(f, "%s *= %s\n", splits[1], splits[2])
			}
		case "div":
			fmt.Fprintf(f, "%s /= %s\n", splits[1], splits[2])
		case "mod":
			fmt.Fprintf(f, "%s %%= %s\n", splits[1], splits[2])
		case "eql":
			if splits[1] == lookAheadSplits[1] && lookAheadSplits[2] == "0" {
				fmt.Fprintf(f, "if %s != %s { %[1]s = 1 } else { %[1]s = 0 }\n", splits[1], splits[2])
				i++ // Skip next line
			} else {
				fmt.Fprintf(f, "if %s == %s { %[1]s = 1 } else { %[1]s = 0 }\n", splits[1], splits[2])
			}
		default:
			log.Fatalf("Invalid command: %s", splits[0])
		}
	}
	fmt.Fprintln(f, "\nreturn z == 0")
	fmt.Fprintln(f, "}")
}

func formatOutputFile() {
	cmd := exec.Command("gofmt", "-w", *outputFile)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error formatting file %s: %s", *outputFile, err)
	}
}
