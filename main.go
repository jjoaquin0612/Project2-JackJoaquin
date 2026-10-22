package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/fatih/color"
)

const DATAFILE = "salaries.txt"

const (
	entryLevel  = "Entry-level"
	midLevel    = "Mid-level"
	remoteModel = "Remote"
	onsiteModel = "On-site"
)

type Record struct {
	JobTitle        string
	ExperienceLevel string
	EmploymentType  string
	WorkModel       string
	WorkYear        int
	Residence       string
	Salary          int
	Currency        string
	SalaryUSD       int
	CompanyLocation string
	CompanySize     string
}

type titleStat struct {
	Title  string
	Median int
	Count  int
}

var (
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	blue   = color.New(color.FgBlue)
	header = color.New(color.FgHiMagenta)
)

func main() {
	records, err := loadRecords(DATAFILE)
	if err != nil {
		fmt.Println("Could not load records:", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %d salary records from %s\n", len(records), DATAFILE)

	part1(records)
	part2(records)
	part3(records)
	part4(records)
	part5(records)
}

// helper func to take list of records and return just the salaries
func salariesFor(records []Record, year int, level string, workModel string) []int {
	var out []int //empty slice to collect results
	for _, r := range records {
		//each if-statement checks current record (r) is == to the requested parameter
		if r.WorkYear != year { //checks if r is requested year
			continue
		}
		if level != "" && r.ExperienceLevel != level { //checks if we requested level and if r is requested level
			continue
		}
		if workModel != "" && r.WorkModel != workModel { //checks if we requested workModel and if r is requested workModel
			continue
		}
		out = append(out, r.SalaryUSD) //append record (r) salary that passed all 3 tests
	}
	return out
}

func average(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum / len(vals)
}

func median(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	sorted := slices.Clone(vals) //clone initial slice to not rearrange original data
	slices.Sort(sorted)
	mid := len(sorted) / 2  //calculates index of mid
	if len(sorted)%2 == 1 { //checks if len is odd
		return sorted[mid] //index of mid is true mid
	}
	return (sorted[mid-1] + sorted[mid]) / 2 //if even, take average of higher mid and lower mid
}
