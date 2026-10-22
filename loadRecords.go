package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func loadRecords(path string) ([]Record, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Record
	for _, row := range rows {
		//need to handle type conversion from string to int
		year, err := strconv.Atoi(strings.TrimSpace(row[4]))
		if err != nil {
			continue //passing the header
		}
		salary, err := strconv.Atoi(strings.TrimSpace(row[6]))
		if err != nil {
			return nil, err
		}
		usd, err := strconv.Atoi(strings.TrimSpace(row[8]))
		if err != nil {
			return nil, err
		}
		records = append(records, Record{
			JobTitle:        strings.TrimSpace(row[0]),
			ExperienceLevel: strings.TrimSpace(row[1]),
			EmploymentType:  strings.TrimSpace(row[2]),
			WorkModel:       strings.TrimSpace(row[3]),
			WorkYear:        year,
			Residence:       strings.TrimSpace(row[5]),
			Salary:          salary,
			Currency:        strings.TrimSpace(row[7]),
			SalaryUSD:       usd,
			CompanyLocation: strings.TrimSpace(row[9]),
			CompanySize:     strings.TrimSpace(row[10]),
		})
	}
	return records, nil
}
