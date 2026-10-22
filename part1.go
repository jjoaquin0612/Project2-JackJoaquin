package main

import (
	"fmt"
)

func part1(records []Record) {
	header.Println("\nPart 1.) AVERAGE entry level salary (USD)")
	s24 := salariesFor(records, 2024, entryLevel, "")
	s23 := salariesFor(records, 2023, entryLevel, "")
	avg24, avg23 := average(s24), average(s23)

	fmt.Printf("2023: %d entry-level jobs, average salary $%d\n", len(s23), avg23)
	fmt.Printf("2024: %d entry-level jobs, average salary $%d\n", len(s24), avg24)

	if avg24 < avg23 {
		red.Printf("The 2024 average is LOWER than 2023 by $%d\n", avg23-avg24)
	} else {
		green.Printf("The 2024 average is higher than (or equal to) 2023 by $%d\n", avg24-avg23)
	}
}
