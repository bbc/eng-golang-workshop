// Cf converts its numeric argument to Celsius, Fahrenheit, Feet, Metres, Pounds and Kilograms.
package main

import (
    "fmt"
    "os"
    "strconv"

    "garymacindoe.co.uk/ch2/lengthconv"
    "garymacindoe.co.uk/ch2/tempconv"
    "garymacindoe.co.uk/ch2/weightconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        ft := lengthconv.Foot(t)
        m := lengthconv.Meter(t)
        p := weightconv.Pound(t)
        kg := weightconv.Kilogram(t)

        fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
        fmt.Printf("%s = %s, %s = %s\n", ft, lengthconv.FToM(ft), m, lengthconv.MToF(m))
        fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), kg, weightconv.KToP(kg))
    }
}
