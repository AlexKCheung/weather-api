// Alex Cheung
// Cloudflare take home test
package main 

import "fmt"

// some other imports
import "net/http"
import "io/ioutil"
import "strconv"

func main() {

	// print statement
	fmt.Println("Hello, World!")
	// for loop 
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Cloudflare SF, wfo Forecast office ID: MTR San Francisco Bay Area, CA (Monterey)
	latitude := 37.780231
	longitude := -122.390472

	forecast := get_grid_forecast(latitude, longitude)
	fmt.Println(forecast)

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

// returns gridpoints X, Y used to get forecasts 
func get_grid_forecast(latitude float64, longitude float64) (string) {
	// https://stackoverflow.com/questions/53312828/how-to-convert-float-to-string
	lat_long := strconv.FormatFloat(latitude, 'f', -1, 64) + "," + strconv.FormatFloat(longitude, 'f', -1, 64)

	// Request
	url := "https://api.weather.gov/points/" + lat_long
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "Error in request"
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
		return "Error in sending request"
	}
	// fmt.Println("RESPONSE:", resp)

	// Apparently we need to close the response body
	defer resp.Body.Close()

	// Parse Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "Error in parsing response"
	}
	// fmt.Println("BODY:", body)

	return string(body)

}
