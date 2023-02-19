package main

import (
	"fdc/parser"
	"fdc/csv_writer"
)

func main() {
	col_names_to_keep := []string{
		"Div",
		"Date",
		"HomeTeam",
		"AwayTeam",
		"FTHG",
		"FTAG",
		"FTR",
	}
	data_map := parser.GetData("https://www.football-data.co.uk", "mmz4281", "2223", "D1", col_names_to_keep)
	csv_writer.MapToCsv(data_map, "data.csv")
}