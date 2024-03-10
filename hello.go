// Alex Cheung
// Cloudflare take home test
package main 

import "fmt"

// some other imports
import "net/http"
import "io/ioutil"

func main() {

	// print statement
	fmt.Println("Hello, World!")
	// for loop 
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Create Request
	url := "https://api.weather.gov"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	fmt.Println("REQUEST", req)
	
	// Authentication
	// User-Agent: (myweatherapp.com, contact@myweatherapp.com)
	req.Header.Set("User-Agent", "alexcheung880@gmail.com")

	// Send request
	// client := &http.Client{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return 
	}
	fmt.Println("RESPONSE:", resp)

	// Apparently we need to close the response body
	defer resp.Body.Close()

	// Parse Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return 
	}
	fmt.Println("BODY:", body)
	









}
