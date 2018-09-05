// Test for conversion.go
package conversion

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestKelvinConversion(t *testing.T) {
	callFunction("Fahrenheit=-456.67°F Celsuis=-271.15°C", "2k\n", t)
}

func TestFahrenheitConversion(t *testing.T) {
	callFunction("Kelvin=256.48333333333335°K Celsius=-16.666666666666668°C", "2f\n", t)
}

func TestCelcuisConversion(t *testing.T) {
	callFunction("Kelvin=275.15°K Fahrenheit=35.6°F", "2c\n", t)

}

func TestForNotSupportedUnit(t *testing.T) {
	callFunction("not supported", "2u\n", t)
}

func TestForInvalidInput(t *testing.T) {
	callFunctionWithError("unitConversion: strconv.Atoi: parsing \"adf\": invalid syntax", "adfk", t)
}

func callFunctionWithError(expected string, text string, t *testing.T) {
	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	ConvertUnits(strings.NewReader(text))
	w.Close()
	os.Stderr = rescueStderr
	out, _ := ioutil.ReadAll(r)
	if strings.TrimSuffix(string(out), "\n") != expected {
		t.Errorf("Output was incorrect expected \"%s\" but got \"%s\"", expected, out)

	}
}

func callFunction(expected string, text string, t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ConvertUnits(strings.NewReader(text))
	w.Close()
	os.Stdout = rescueStdout
	out, _ := ioutil.ReadAll(r)
	if strings.TrimSuffix(string(out), "\n") != expected {
		t.Errorf("Output was incorrect expected \"%s\" but got \"%s\"", expected, out)

	}
}
