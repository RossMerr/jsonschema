// Code generated by jsonschema. DO NOT EDIT.
package main

// A geographical coordinate on a planet (most commonly Earth).
// ID: https://example.com/geographical-location.schema.json
type GeographicalLocation struct {
	Latitude  *float64 `json:"latitude,omitempty", validate:"required,lte=90,gte=-90"`
	Longitude *float64 `json:"omitempty,longitude", validate:"gte=-180,required,lte=180"`
}
