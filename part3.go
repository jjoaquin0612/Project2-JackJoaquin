package main

import "fmt"

func part3(records []Record) {
	header.Println("\nPart 3.) Moving from entry level to mid level (2024 medians)")
	entry := median(salariesFor(records, 2024, entryLevel, ""))
	mid := median(salariesFor(records, 2024, midLevel, ""))

	fmt.Printf("Median entry-level salary: $%d\n", entry)
	fmt.Printf("Median mid-level salary: $%d\n", mid)
	raise := mid - entry
	if entry > 0 {
		yellow.Printf("Moving from entry to mid-level could mean a raise of about $%d\n", raise)
	} else {
		yellow.Printf("Raise: $%d (no entry level data to compute a percentage)\n", raise)
	}
}
