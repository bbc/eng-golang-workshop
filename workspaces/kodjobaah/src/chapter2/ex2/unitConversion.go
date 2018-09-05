//General Purpose unit conversion application
package main

import (
	"bufio"
	"chapter2/ex1/tempconv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	files := os.Args[1:]
	if len(files) == 0 {
		convertUnits(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unitConversion: %v\n", err)
				continue
			}
			convertUnits(f)
			f.Close()
		}

	}
}

func convertUnits(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		val := input.Text()
		err, num, unit := parseInput(strings.TrimSpace(val))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitConversion: %v\n", err)
			continue
		}
		processUnit(unit, num)
	}
}

func parseInput(text string) (error, int, string) {
	if len(text) < 2 {
		return errors.New("unitConversion: you need to provide a number with units .i.e 2K"), 0, ""
	}

	unit := text[len(text)-1:]
	num := string(text[0 : len(text)-1])

	convNum, err := strconv.Atoi(num)
	return err, convNum, unit
}

func processUnit(unit string, num int) {
	switch strings.ToLower(unit) {
	case "k":
		value := tempconv.Kelvin(float64(num))
		fmt.Printf("Fahrenheit=%v Celsuis=%v\n", tempconv.KToF(value), tempconv.KToC(value))

	case "f":
		value := tempconv.Fahrenheit(float64(num))
		fmt.Printf("Kelvin=%v Celsius=%v\n", tempconv.FToK(value), tempconv.FToC(value))

	case "c":
		value := tempconv.Celsius(float64(num))
		fmt.Printf("Kelvin=%v Fahrenheit=%v\n", tempconv.CToK(value), tempconv.CToF(value))

	default:
		fmt.Println("not supported")

	}
}
