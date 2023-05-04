package routes

import (
	controller "api-geolocation/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func GetRoutes(e *echo.Echo) {
	e.GET("/countries", controller.GetCountries)
	e.GET("/countries/:countryCode", controller.GetCountryByCode)
	e.GET("/countries/:countryCode/states", controller.GetCountryStateList)
	e.GET("/countries/:countryCode/states/:stateCode", controller.GetCountryState)
	e.GET("/countries/:countryCode/states/:stateCode/cities", controller.GetCountryStateCityList)
}