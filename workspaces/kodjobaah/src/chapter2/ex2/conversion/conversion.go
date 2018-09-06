package conversion

import (
	"bufio"
	"chapter2/ex1/tempconv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConvertUnits(f io.Reader) {
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
	var (
		err     error
		convNum int
		unit    string
	)

	if len(text) < 2 {
		err = errors.New("you need to provide a number with units .i.e 2K")
	} else {
		unit = text[len(text)-1:]
		num := strings.TrimSpace(string(text[0 : len(text)-1]))
		convNum, err = strconv.Atoi(num)
	}
	return err, convNum, unit
}

func processUnit(unit string, num int) {
	switch strings.ToLower(unit) {
	case "k":
		value := tempconv.Kelvin(num)
		fmt.Printf("Fahrenheit=%v Celsuis=%v\n", tempconv.KToF(value), tempconv.KToC(value))

	case "f":
		value := tempconv.Fahrenheit(num)
		fmt.Printf("Kelvin=%v Celsius=%v\n", tempconv.FToK(value), tempconv.FToC(value))

	case "c":
		value := tempconv.Celsius(num)
		fmt.Printf("Kelvin=%v Fahrenheit=%v\n", tempconv.CToK(value), tempconv.CToF(value))

	default:
		fmt.Println("not supported")

	}
}
