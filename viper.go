package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("File to read? (ex: test-rest-get.yaml)")

	// var then variable name then variable type
	var file string

	// Taking input from user
	fmt.Scanln(&file)

	vi := viper.New()
	vi.SetConfigFile(file)
	vi.ReadInConfig()

	// Using https://github.com/spf13/viper#getting-values-from-viper

	if validFormat(vi) {
		// Extract YAML first level parameters
		name := vi.Get("name")
		id := vi.Get("id")

		// Any Implementation
		fmt.Println("Name:", name)
		fmt.Println("Id:", id)

		api := vi.Get("inputs.request.api")
		if api == "REST" {
			// Extract request parameters
			method := vi.Get("inputs.request.method").(string)
			url := vi.Get("inputs.request.url").(string)
			headers := vi.Get("inputs.request.headers").(map[string]interface{})
			var body map[string]interface{}
			if method != "GET" && method != "DELETE" {
				body = vi.Get("inputs.request.body").(map[string]interface{})
			}

			// Any Implementation
			fmt.Println("Method:", method)
			fmt.Println("Url:", url)
			fmt.Println("Headers:", headers)
			fmt.Println("Body:", body)

			// Test method
			consumeAPI(method, url, headers, body)
		}
		if api == "gRPC" {
			// Extract request parameters
			serverAddress := vi.Get("inputs.request.server-address")
			service := vi.Get("inputs.request.service")
			fields := vi.Get("inputs.request.fields")

			// Any Implementation
			fmt.Println("Server Address:", serverAddress)
			fmt.Println("Service:", service)
			fmt.Println("Fields:", fields)
		}
	} else {
		fmt.Println("Please, check the YAML format on the repository README.md file!")
	}
}

func validFormat(vi *viper.Viper) bool {
	// Check YAML first level parameters
	if !vi.IsSet("name") {
		fmt.Println("ERROR yaml format: Name")
		return false
	}
	if !vi.IsSet("id") {
		fmt.Println("ERROR yaml format: Id")
		return false
	}
	if !vi.IsSet("inputs") {
		fmt.Println("ERROR yaml format: Inputs")
		return false
	}
	// Check YAML request parameter
	if !vi.IsSet("inputs.request") {
		fmt.Println("ERROR yaml format: Request")
		return false
	}
	api := vi.Get("inputs.request.api")
	if api == "REST" {
		// Check REST request parameters
		if !vi.IsSet("inputs.request.method") {
			fmt.Println("ERROR request format: Method")
			return false
		}
		if !vi.IsSet("inputs.request.url") {
			fmt.Println("ERROR request format: Url")
			return false
		}
		if !vi.IsSet("inputs.request.headers") {
			fmt.Println("ERROR request format: Headers")
			return false
		}
		method := vi.Get("inputs.request.method")
		if method != "GET" && method != "DELETE" {
			if !vi.IsSet("inputs.request.body") {
				fmt.Println("ERROR request format: Body")
				return false
			}
		}
	}

	if api == "gRPC" {
		// Check gRPC request parameters
		if !vi.IsSet("inputs.request.server-address") {
			fmt.Println("ERROR request format: Server Address")
			return false
		}
		if !vi.IsSet("inputs.request.service") {
			fmt.Println("ERROR request format: Service")
			return false
		}
		if !vi.IsSet("inputs.request.fields") {
			fmt.Println("ERROR request format: Fields")
			return false
		}
	}

	return true
}

func consumeAPI(method string, url string, headers map[string]interface{}, body map[string]interface{}) {
	client := &http.Client{}

	var jsonStr []byte
	if len(body) != 0 {
		j, err := json.Marshal(body)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		jsonStr = j
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Print("\nError Request:", err.Error())
	}

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("\nError Call:", err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("\nError Read:", err.Error())
	}

	fmt.Println("\nStatusCode:", resp.StatusCode)
	fmt.Println("Response:", string(bodyBytes))
}
