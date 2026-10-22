package main

import "fmt"

func part2(records []Record) {
	header.Println("\nPart 2.) MEDIAN entry level salary (USD)")
	s24 := salariesFor(records, 2024, entryLevel, "")
	s23 := salariesFor(records, 2023, entryLevel, "")
	med24, med23 := median(s24), median(s23)

	fmt.Printf("2023 median entry-level salary: $%d\n", med23)
	fmt.Printf("2024 median entry-level salary: $%d\n", med24)
	if med23 > med24 {
		red.Printf("The median DROPPED by $%d from 2023 to 2024\n", med23-med24)
	} else {
		green.Printf("The median ROSE by $%d from 2023 to 2024\n", med24-med23)
	}
}
