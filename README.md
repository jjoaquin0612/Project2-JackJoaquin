# Project 2: COMP 415 - Golang (Jack Joaquin)

A Go program that reads `salaries.txt` and prints color-coded salary analysis.
All salaries use the **USD** column (`salary_in_usd`) so they are comparable.

## How to read the output
Each section is headed by a magenta title:
1. **Average entry level salary** for 2023 and 2024, with the difference.
2. **Median entry level salary** for 2023 and 2024, the change between them, and how far the
   average sits above the median in each year (a big gap means a few very high salaries).
3. **Entry to mid level raise** using 2024 medians, in dollars and percent.
4. **Remote vs on-site** 2024 medians, for entry level jobs and for all levels.
5. **Top paying entry level jobs in 2024** by median salary, excluding titles containing "Research".

## What the colors mean
- **Red**: a bad/lower outcome: 2024 average lower than 2023 (section 1), 2023 median higher than
  2024 (section 2), or on-site median higher than remote (section 4).
- **Green**: the opposite outcome in sections 1, 2 and 4; also the #1 highest paying job in section 5.
- **Yellow**: the entry-to-mid raise (section 3).
- **Blue**: average-vs-median comparison (section 2) and jobs ranked #2-#5 (section 5).
- **Magenta**: section headings.
