// Code generated by jsonschema. DO NOT EDIT.
package basic

// ID: https://example.com/basic.json

type Person struct {
	// The person's first name.
	FirstName *string `json:"firstName,omitempty"`
	// The person's last name.
	LastName *string `json:"lastName,omitempty"`
	// Age in years which must be equal to or greater than zero.
	Age *int32 `json:"age,omitempty", validate:"gte=0"`
}
