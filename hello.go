// Alex Cheung
// Cloudflare take home test
package main 

import "fmt"

// some other imports
import "net/http"
import "io/ioutil"
// import "strconv"
import "encoding/json"


func main() {

	// print statement
	// fmt.Println("Hello, World!")
	// for loop 
	// i := 1
	// for i <= 3 {
	// 	fmt.Print(i)
	// 	i = i + 1
	// }
	// fmt.Println()

	// Cloudflare SF, wfo Forecast office ID: MTR San Francisco Bay Area, CA (Monterey)
	latitude := 37.780231
	longitude := -122.390472

	forecast := get_grid_forecast(latitude, longitude)
	// fmt.Println(forecast)
	if forecast == nil {
		fmt.Println("ERROR: Failed to retrieve forecast")
		return 
	}

	// initialize dictionary
	temperatureByStartTime := make(map[string] int)

	for _, period := range forecast.Properties.Periods {
		// fmt.Print("Temp:", period.Temperature)
		// fmt.Print("Start:", period.StartTime)
		// fmt.Print()
		temperatureByStartTime[period.StartTime] = period.Temperature
	}

	// print dictionary
	fmt.Println(temperatureByStartTime)

	// fmt.Println(forecast)

	// // Create Request
	// url := "https://api.weather.gov"
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	fmt.Println("ERROR:", err.Error())
	// 	return
	// }
	// fmt.Println("REQUEST", req)
	
	// // Authentication
	// // User-Agent: (myweatherapp.com, contact@myweatherapp.com)
	// req.Header.Set("User-Agent", "alexcheung880@gmail.com")

	// // Send request
	// // client := &http.Client{}
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("ERROR:", err.Error())
	// 	return 
	// }
	// fmt.Println("RESPONSE:", resp)

	// // Apparently we need to close the response body
	// defer resp.Body.Close()

	// // Parse Response
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("ERROR:", err.Error())
	// 	return 
	// }
	// fmt.Println("BODY:", body)

	// done with program
	return

}

type GridForecastResponse struct {
	Properties struct {
		Periods []struct {
			Number int `json:"number"`
			Name string `json:"name"`
			StartTime string `json:"startTime"`
			EndTime string `json:"endTime"`
			Temperature int `json:"temperature"`
		}
	} `json:"properties"`
		
}


// returns gridpoints X, Y used to get forecasts 
func get_grid_forecast(latitude float64, longitude float64) (*GridForecastResponse) {
	// https://stackoverflow.com/questions/53312828/how-to-convert-float-to-string
	// lat_long := strconv.FormatFloat(latitude, 'f', -1, 64) + "," + strconv.FormatFloat(longitude, 'f', -1, 64)

	// Request
	// url := "https://api.weather.gov/points/" + lat_long
	url := "https://api.weather.gov/gridpoints/MTR/86,105/forecast/hourly"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return nil
	}
	// fmt.Println("REQUEST", req)
	
	// Authentication
	// User-Agent: (myweatherapp.com, contact@myweatherapp.com)
	req.Header.Set("User-Agent", "alexcheung880@gmail.com")

	// Send request
	// client := &http.Client{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return nil
	}
	// fmt.Println("RESPONSE:", resp)

	// Apparently we need to close the response body
	defer resp.Body.Close()

	// Parse Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return nil
	}
	// fmt.Println("BODY:", string(body))

	var gridForecast GridForecastResponse
	err = json.Unmarshal(body, &gridForecast)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return nil
	}

	// fmt.Println("Grid object", gridForecast)

	return &gridForecast

}
