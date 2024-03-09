tp.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	fmt.Println(response.Body)
	body, err := io.ReadAll(response.Body)
	fmt.Println(body)

	var weather_Info WeatherInfo
	err = json.Unmarshal(body, &weather_Info)
	if err != nil {
		return 
	}

	fmt.Println(weather_Info)