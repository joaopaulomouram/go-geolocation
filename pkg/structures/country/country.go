package struct_country

import struct_state "api-geolocation/pkg/structures/state"

type Country struct {
	Name string `json:"name"`
	Iso3 string `json:"iso3"`
	Currency_symbol string `json:"currency_symbol"`
	States []struct_state.State `json:"states,omitempty"`
}