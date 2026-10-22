package main

import "fmt"

func part4(records []Record) {
	header.Println("\nPart 4.) Remote vs. OnSite (2024 medians)")
	remoteEntry := median(salariesFor(records, 2024, entryLevel, remoteModel))
	onsiteEntry := median(salariesFor(records, 2024, entryLevel, onsiteModel))
	remoteAll := salariesFor(records, 2024, "", remoteModel)
	onsiteAll := salariesFor(records, 2024, "", onsiteModel)
	medRemote, medOnsite := median(remoteAll), median(onsiteAll)

	blue.Printf("2024 remote jobs: %d\n", len(remoteAll))
	blue.Printf("2024 on-site jobs: %d\n", len(onsiteAll))
	fmt.Printf("Median remote entry-level salary: $%d\n", remoteEntry)
	fmt.Printf("Median onsite entry-level salary: $%d\n", onsiteEntry)
	fmt.Printf("Median remote salary: $%d\n", medRemote)
	fmt.Printf("Median on-site salary: $%d\n", medOnsite)
	if remoteEntry > onsiteEntry {
		green.Println("Median remote entry-level salary is greater than on-site")
	} else {
		red.Println("Median on-site entry-level salary is greater than remote")
	}
	if medRemote > medOnsite {
		green.Println("Median remote salary for all levels is greater than on-site")
	} else {
		red.Println("Median on-site salary for all levels is greater than remote")
	}
}
