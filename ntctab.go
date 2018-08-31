package main

import (
	"flag"
	"fmt"
	"os"
	"math"
)

func getNtcTemp(r, r0, b, t0 float64) float64 {
	t := 1.0 / (math.Log(r / r0) / b + 1.0 / (t0 + 273.0))
	t -= 273.0
	return t
} 

func printNtcTable(name string, r0 float64, b float64, pullup float64, adcMax int, scale float64) {
	fmt.Printf("static int16_t %s[] = {\n", name)
	for i := 0; i < adcMax+1; i += 1 {
		p := i
		if p == 0 {
			p = 1
		}
		
		u := float64(p)
		r := u / (float64(adcMax) - u) * pullup;
		t := getNtcTemp(float64(r), r0, b, 25.0)
		if i % 16 == 0 {
			fmt.Printf("\t")
		}
		fmt.Printf("%d", int(t*scale))
		if i != adcMax {
			fmt.Printf(",\t")
		}
		if i % 16 == 15 {
			fmt.Println()
		}
	}
	fmt.Println("};")
}

func main() {
	name := flag.String("name", "ntc_temp_table", "generated array var name")
	resistance := flag.Float64("resistance", 2.2, "nominal resistance (in 25°C) of ntc sensor [kOhm]")
	pullup := flag.Float64("pullup", 6.8, "sensor pullup circuit resistance [kOhm]")
	beta := flag.Float64("beta", 3750.0, "contant beta value")
	adc := flag.Int("adc", 255, "maximum adc value")
	scale := flag.Float64("scale", 1, "temperature scale")
	flag.Parse()
	printNtcTable(*name, *resistance, *beta, *pullup, *adc, *scale);
}

