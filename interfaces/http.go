package http

import (
	//"io/ioutil"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

// https://www.football-data.co.uk/mmz4281/2223/D1.csv

// FetchData requests csv from URL and saves / returns it as ?
func ReadCSVFromUrl(url string) ([][]string, error) {

	fmt.Println("Requesting URL \n" + url + "\n")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}