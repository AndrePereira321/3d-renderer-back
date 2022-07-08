package main

import (
	"encoding/json"
	"log"
	"os"
	"server/database"
	"server/database/cache"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	database.Init()

	//insertCountries()
	//insertLanguages()
	//insertComparationOperators()
	//insertLogicalOperators()
}

func insertComparationOperators() {
	dto := cache.ReferenceItemDTO{
		DTO: database.DTO{
			CollectionName: "References",
		},
		Table: cache.ReferenceTable{
			TableCode: "COMP_OPS",
			Text1:     "Comparation Operators",
		},
		Values: []cache.ReferenceValue{
			{
				Code:  "$eq",
				Text1: "Equals",
			},
			{
				Code:  "$ne",
				Text1: "Different",
			},
			{
				Code:  "$lt",
				Text1: "Lesser Than",
			},
			{
				Code:  "$lte",
				Text1: "Lesser/Equals Than",
			},
			{
				Code:  "$gt",
				Text1: "Greater Than",
			},
			{
				Code:  "$gte",
				Text1: "Greater/Equals Than",
			},
			{
				Code:  "$in",
				Text1: "In",
			},
			{
				Code:  "$nin",
				Text1: "Not In",
			},
			{
				Code:  "$regex",
				Text1: "Contains",
			},
		},
	}
	dto.Save(&dto)
}

func insertLogicalOperators() {
	dto := cache.ReferenceItemDTO{
		DTO: database.DTO{
			CollectionName: "References",
		},
		Table: cache.ReferenceTable{
			TableCode: "LOG_OPS",
			Text1:     "Logical Operators",
		},
		Values: []cache.ReferenceValue{
			{
				Code:  "$and",
				Text1: "And",
			},
			{
				Code:  "$or",
				Text1: "Or",
			},
			{
				Code:  "$not",
				Text1: "Not",
			},
		},
	}
	dto.Save(&dto)
}

func insertLanguages() {
	dto := cache.ReferenceItemDTO{
		DTO: database.DTO{
			CollectionName: "References",
		},
		Table: cache.ReferenceTable{
			TableCode: "LANG",
			Text1:     "Languages",
		},
		Values: []cache.ReferenceValue{
			{
				Code:  "EN",
				Text1: "English",
			},
			{
				Code:  "FR",
				Text1: "French",
			},
			{
				Code:  "PT",
				Text1: "Portuguese",
			},
			{
				Code:  "ES",
				Text1: "Spanish",
			},
		},
	}
	dto.Save(&dto)
}

func insertCountries() {
	type CountryJSONObj struct {
		CountryCode    string `json:"countryCode"`
		CountryName    string `json:"countryName"`
		CurrencyCode   string `json:"currencyCode"`
		Population     string `json:"population"`
		Capital        string `json:"capital"`
		CountinentName string `json:"continentName"`
	}
	type CountriesJSONObj struct {
		Country []CountryJSONObj `json:"country"`
	}

	type CountriesJSON struct {
		Countries CountriesJSONObj `json:"countries"`
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
			CreatedDate:    primitive.Timestamp{T: uint32(time.Now().Unix())},
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
