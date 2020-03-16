// Code generated by jsonschema. DO NOT EDIT.
package main

// A product from Acme's catalog
// ID: http://example.com/product.schema.json
type Product struct {
	// Tags for the product
	Tags       []string `json:"tags"`
	Dimensions struct {
		// length
		Length *float64 `json:"omitempty,length", validate:"required"`
		Width  *float64 `json:"width,omitempty", validate:"required"`
		Height *float64 `json:"omitempty,height", validate:"required"`
	} `json:"Dimensions"`

	GeographicalLocation struct {
		Latitude  *float64 `json:"omitempty,latitude", validate:"required,lte=90,gte=-90"`
		Longitude *float64 `json:"longitude,omitempty", validate:"required,lte=180,gte=-180"`
	} `json:"GeographicalLocation"`

	// The unique identifier for a product
	ProductId *int32 `json:"productId,omitempty", validate:"required"`
	// Name of the product
	ProductName string `json:"productName,omitempty", validate:"required"`
	// The price of the product
	Price *float64 `json:"omitempty,price", validate:"required,gt=0"`
}
