package main

import (
	"fmt"
	"sort"
	"strings"
)

func part5(records []Record) {
	header.Println("\nPart 5.) Highest paying entry level jobs, 2024 " +
		"\n(median USD, 'Research' titles excluded)")

	byTitle := make(map[string][]int) //make a map that takes strings as keys and ints as values
	for _, r := range records {
		if r.WorkYear != 2024 || r.ExperienceLevel != entryLevel { //find all entryLevel jobs in 2024
			continue
		}
		if strings.Contains(strings.ToLower(r.JobTitle), "research") {
			continue
		}
		byTitle[r.JobTitle] = append(byTitle[r.JobTitle], r.SalaryUSD) //all records that pass 2 tests get appended
	}

	var stats []titleStat //add to struct to save both job title and
	for title, vals := range byTitle {
		stats = append(stats, titleStat{title, median(vals), len(vals)})
	}

	//sort.Slice takes in 2 arguments,
	//stats is the struct to sort and func is the function that sorts it
	sort.Slice(stats, func(i, j int) bool {
		//if statement checks if equal, returns true or false depending on if i or j median is bigger
		if stats[i].Median != stats[j].Median {
			return stats[i].Median > stats[j].Median
		}
		return stats[i].Title < stats[j].Title //sorts alphabetically
	})

	if len(stats) == 0 {
		fmt.Println("No results found...")
		return
	}

	green.Printf("#1 %s: median %d (%d jobs)\n", stats[0].Title, stats[0].Median, stats[0].Count)
	for i := 1; i < len(stats) && i < 5; i++ {
		blue.Printf("#%d %s: median %d (%d jobs)\n", i+1, stats[i].Title, stats[i].Median, stats[i].Count)
	}
}
