package main

import (
	"encoding/json"
	"log"
	"os"
	"server/database"
	"server/database/cache"
)

type CountriesJSON struct {
	Countries CountriesJSONObj `json:"countries"`
}

type CountriesJSONObj struct {
	Country []CountryJSONObj `json:"country"`
}

type CountryJSONObj struct {
	CountryCode    string `json:"countryCode"`
	CountryName    string `json:"countryName"`
	CurrencyCode   string `json:"currencyCode"`
	Population     string `json:"population"`
	Capital        string `json:"capital"`
	CountinentName string `json:"continentName"`
}

func main() {
	_, err := database.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile("./scripting/resources/countries.json")
	if err != nil {
		log.Fatal(err)
	}

	countries := CountriesJSON{}
	err = json.Unmarshal(file, &countries)
	if err != nil {
		log.Fatal(err)
	}

	dto := cache.ReferenceItemDTO{
		DTO: database.DTO{
			CollectionName: "References",
		},
		Table: cache.ReferenceTable{
			TableCode: "COUNTRIES",
			Text1:     "Countries",
			Disabled:  false,
		},
		Values: make([]cache.ReferenceValue, len(countries.Countries.Country)),
	}

	for i := 0; i < len(countries.Countries.Country); i++ {
		country := countries.Countries.Country[i]
		dto.Values[i] = cache.ReferenceValue{
			Code:  country.CountryCode,
			Text1: country.CountryName,
			Text2: country.Capital,
			Text3: country.CurrencyCode,
			Text4: country.CountinentName,
			Text5: country.Population,
		}
	}

	dto.Save(dto)
}
