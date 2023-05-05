package countries

import (
	struct_country "api-geolocation/pkg/structures/country"
	struct_state "api-geolocation/pkg/structures/state"
	"encoding/json"
	"log"
	"os"
)

var countryList []struct_country.Country

func GetCountriesList() []struct_country.Country{

	if len(countryList) > 0 {
		return countryList
	}

	file, err := os.ReadFile("internal/data.json")
	if err != nil {
		log.Fatal("Failed to read file: ", err)
	}

	err = json.Unmarshal(file, &countryList)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON data: ", err)
	}

	return countryList
}

func FindCountryByCode(countryCode string, countryList []struct_country.Country) struct_country.Country {
	for _, v := range countryList {
		if v.Iso3 == countryCode {
			return v
		}
	}

	return struct_country.Country {}
}

func FindStateByCode(stateCode string, stateList []struct_state.State) struct_state.State {
	for _, v := range stateList {
		if v.State_code == stateCode {
			return v
		}
	}

	return struct_state.State {}
}