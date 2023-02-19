package parser

import (
	"log"
	"example.com/interfaces"
)

//////
// Package to fetch and parse data from football-data.co.uk
/////

// FetchData requests csv from URL and saves / returns it as ?
func fetchData(base_url string, country_id string, year_code string, div string) ([][]string, error) {

	url := base_url + "/" + country_id + "/" + year_code + "/" + div + ".csv"
	data, err := http.ReadCSVFromUrl(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// stringInSlice checks if a single string is in a list of strings
func stringInSlice(a string, list []string) bool {
    for _, b := range list {

        if b == a {
            return true
        }
    }
    return false
}

// filterColumnsByName filters an array of array by only keeping the columns 
// in which the name in the first row matches one of the col_names_to_keep input
// It retuns the data as map, in which the key is the column name
func filterColumnsByName(col_names_to_keep []string, data[][]string) (map[string][]string) {

	// Get index of columns
	idx_map := make(map[string]int)
	for index, element := range data[0] {
		if stringInSlice(element, col_names_to_keep) {
			idx_map[element] = index
		}
	}
	if len(col_names_to_keep) != len(idx_map) {
		log.Fatalln("One of column names not in first array of data")
	}

	// Filter data
	filtered_data_map := make(map[string][]string)
	for i := 1; i < len(data); i++ {
		for key, element := range idx_map {
			filtered_data_map[key] = append(filtered_data_map[key], data[i][element])
		}
	}

	return filtered_data_map
}

func GetData(base_url string, country_id string, year_code string, div string, col_names_to_keep []string) map[string][]string {
	// Get Data from URL as slcie of slices, with first array being the header
	data, err := fetchData(base_url, country_id, year_code, div)
	if err != nil {
		log.Fatal(err)
	}

	filtered_data_map := filterColumnsByName(col_names_to_keep, data)
	
	return filtered_data_map
}
