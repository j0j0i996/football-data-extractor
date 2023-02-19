package csv_writer

import (
	"log"
	"encoding/csv"
	"os"
)

// arrayMapToArrayOfArrays transforms a map with arrays as elements to an array of arrays
func arrayMapToArrayOfArrays(data_map map[string][]string) [][]string {

	// get length of arrays in map
	var len_array int
	for _, element := range data_map {
		len_array = len(element)
		break
	}
	
	data := make([][]string, len_array + 1)

	// transform to list of lists
	for key, element := range data_map {
		data[0] = append(data[0], key)
		for i := 0; i < len(element); i++ {
			data[i+1] = append(data[i+1], element[i])
		} 
	}

	return data
}

// MapToCsv Takes a map with arrays as elements and writes it to a string, with the keys being the headers
func MapToCsv(data_map map[string][]string, filename string) {
	data := arrayMapToArrayOfArrays(data_map)

	f, err := os.Create(filename)
    defer f.Close()

    if err != nil {
        log.Fatalln("failed to open file", err)
    }
    w := csv.NewWriter(f)
	defer w.Flush()

	w.WriteAll(data)
} 