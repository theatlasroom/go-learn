package main

import (
	"flag"
	"fmt"
)

/*
* Calculates beer calories
* ABV * Factor (2.5) * Volume in Oz
 */

const FACTOR = 2.5

func mlToOz(ml float64) float64 {
	ozfactor := 0.033814
	return ml * ozfactor
}

func calc(abv, oz float64) float64 {
	return FACTOR * abv * oz
}

func main() {
	volume := flag.Float64("v", 0, "Volume of beer")
	abv := flag.Float64("abv", 0, "ABV of beer")

	flag.Parse()
	fmt.Printf("\nvolume %fml", *volume)
	fmt.Printf("\nabv %f", *abv)
	fmt.Println("\nignored flags", flag.Args())

	oz := mlToOz(*volume)
	cals := calc(*abv, oz)
	fmt.Printf("\nCalories %f from %foz\n", cals, oz)
}
