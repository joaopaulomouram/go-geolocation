package countries

import (
	services "api-geolocation/pkg/services"
	struct_country "api-geolocation/pkg/structures/country"
	struct_state "api-geolocation/pkg/structures/state"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCountries(c echo.Context) error {
	countryList := services.GetCountriesList()
	var response []struct_country.Country
	for _, v := range countryList {
		v.States = nil
		response = append(response, v)
	}		
	return c.JSON(http.StatusOK, response)
}


func GetCountryByCode(c echo.Context) error {
	countryList := services.GetCountriesList()
	countryCode := c.Param("countryCode")
	countryObj := services.FindCountryByCode(countryCode, countryList)
	countryObj.States = nil
	return c.JSON(http.StatusOK, countryObj)
}

func GetCountryStateList(c echo.Context) error {
	countryList := services.GetCountriesList()
	countryCode := c.Param("countryCode")
	countryObj := services.FindCountryByCode(countryCode, countryList)

	var response []struct_state.State
	for _, v := range countryObj.States {
		v.Cities = nil
		response = append(response, v)
	}		
	return c.JSON(http.StatusOK, response)
}

func GetCountryState(c echo.Context) error {
	countryList := services.GetCountriesList()
	countryCode := c.Param("countryCode")
	stateCode := c.Param("stateCode")
	countryObj := services.FindCountryByCode(countryCode, countryList)
	state := services.FindStateByCode(stateCode, countryObj.States)
	state.Cities = nil
	return c.JSON(http.StatusOK, state)
}

func GetCountryStateCityList(c echo.Context) error {
	countryList := services.GetCountriesList()
	countryCode := c.Param("countryCode")
	stateCode := c.Param("stateCode")
	countryObj := services.FindCountryByCode(countryCode, countryList)
	state := services.FindStateByCode(stateCode, countryObj.States)
	return c.JSON(http.StatusOK, state.Cities)
}