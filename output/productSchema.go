// Code generated by jsonschema. DO NOT EDIT.
package main

// A product from Acme's catalog
// ID: http://example.com/product.schema.json
type ProductSchema struct {
	WarehouseLocation struct {
		Latitude  *float64 `json:"omitempty,latitude", validate:"gte=-90,required,lte=90"`
		Longitude *float64 `json:"longitude,omitempty", validate:"required,lte=180,gte=-180"`
	}

	// The unique identifier for a product
	ProductId *int32 `json:"productId,omitempty", validate:"required"`
	// Name of the product
	ProductName string `json:"productName,omitempty", validate:"required"`
	// The price of the product
	Price *float64 `json:"price,omitempty", validate:"required,gt=0"`
	// Tags for the product
	Tags       []string `json:"tags"`
	Dimensions struct {
		Width  *float64 `json:"width,omitempty", validate:"required"`
		Height *float64 `json:"omitempty,height", validate:"required"`
		// length
		Length *float64 `json:"length,omitempty", validate:"required"`
	}
}