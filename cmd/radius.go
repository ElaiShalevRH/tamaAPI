package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	cobra "github.com/spf13/cobra"
)

// Use Authorization header wuth apiAuth to connect
// To the TLV API
const authToken = "tlv_848e1b42-c0c2"
const exportURL = "https://api.busnear.by/external/gtfsrt/export"
const apiURL = "https://gisn.tel-aviv.gov.il/arcgis/rest/services/IView2/MapServer/772/query"

// example for a Query with Query params
// "https://gisn.tel-aviv.gov.il/arcgis/rest/services/IView2/MapServer/772/
// query?where=request_num%3D20221983&outFields=*&f=pjson"

var client *http.Client
var address string

var radiusCmd = &cobra.Command{
	Use:   "radius",
	Short: "Locate all building permits around an address",
	Run: func(cmd *cobra.Command, args []string) {
		if address == "" {
			fmt.Println("Address is required")
			return
		}

		params := map[string]string{
			"outFields": "*",
			"f":         "json",
			"where":     "request_num>20241000",
		}

		resp, err := fetchRadius(params)
		if err != nil {
			fmt.Println("Address is required")
			return
		}

		/*
			for key, value := range resp {
				fmt.Printf("|%s| : |%s|", key, value)
				fmt.Println("")

			}
		*/
		// insert CMD logic
		// logic()
		fmt.Println("Response")

		fmt.Println(resp)
	},
}

func init() {
	radiusCmd.Flags().StringVarP(&address, "addy", "a", "a", "a")
}

func connect() *http.Client {

	req, err := http.NewRequest("GET", exportURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	defer resp.Body.Close()
	return client
}

func fetchRadius(params map[string]string) (map[string]interface{}, error) {

	client = &http.Client{}
	connect()

	baseURL, err := url.Parse(apiURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing URL")
	}

	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}

	baseURL.RawQuery = query.Encode()
	resp, err := client.Get(baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error: could not unmarshal json")
	}

	return data, nil

}
