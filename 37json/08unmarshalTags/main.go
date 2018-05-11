package main

import (
	"encoding/json"
	"fmt"
)

type city struct {
	PostalType    string  `json:"Postal"`
	GeoLatitude   float64 `json:"Latitude"`
	GeoLongitude  float64 `json:"Longitude"`
	AdressStreet  string  `json:"Address"`
	AdressCity    string  `json:"City"`
	AdressState   string  `json:"State"`
	AdressZip     string  `json:"Zip"`
	AdressCountry string  `json:"Country"`
}

type cities []city

func main() {

	var data cities

	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(data[1].AdressCountry)
}
