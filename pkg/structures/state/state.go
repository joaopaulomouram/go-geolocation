package struct_state

import struct_city "api-geolocation/pkg/structures/city"

type State struct {
	Name string `json:"name"`
	State_code string `json:"state_code"`
	Cities []struct_city.City `json:"cities,omitempty"`
}