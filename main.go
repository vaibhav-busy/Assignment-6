package main

import (
	"encoding/json"
	"fmt"
	"regexp"

	// "internal/abi"
	"io"
	"net/http"
)

type City struct {
	Name    string   `json:"name"`
	Weather string   `json:"weather"`
	Status  []string `json:"status"`
}

type ApiResponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []City `json:"data"`
}

func extractNumber(s string) string {
	re := regexp.MustCompile(`\d+`)
	num := re.FindString(s)
	return num
}

func search(url, name string) ([]City, error) {

	pageNo := 1
	var res []City

	for {
		url = fmt.Sprintf("%v%v&page=%v", url, name, pageNo)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while fetching data:", err)
			return nil, err
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var apiRes ApiResponse
		err = json.Unmarshal(body, &apiRes)
		if err != nil {
			return nil, err
		}

		res = append(res, apiRes.Data...)
		if pageNo >= apiRes.TotalPages {
			break
		}
		pageNo++

	}

	return res, nil

}

func main() {

	var name string
	fmt.Println("Enter the name to search in Api: ")
	fmt.Scanln(&name)

	url := "https://jsonmock.hackerrank.com/api/weather/search?name="

	result, err := search(url, name)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// fmt.Println(result)
	// println("Name\t\t", "temp\t", "Wind\t", "Humidity" )
	fmt.Printf("%-15s%-7s%-7s%-7s\n", "Name", "Temp", "Wind", "Humidity")
	for _, val := range result {
		// fmt.Println(val)

		weather := extractNumber(val.Weather)
		wind := extractNumber(val.Status[0])
		humidity := extractNumber(val.Status[1])

		fmt.Printf("%-15s%-7s%-7s%-7s\n", val.Name, weather, wind, humidity)

	}

}
