package dto

import (
	"encoding/json"
	"io"
)

type CreateProductRequest struct {
	Name string `json:"name"`
	Price float32 `json:"price"`
	Description string `json:"description"`
}

// Converts JSON body request to a CreateProductRequest struct
func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {

	// Create a CreateProductRequest struct
	createProductRequest := CreateProductRequest{}

	// Decode JSON body into a CreateProductRequest struct
	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	// Return the CreateProductRequest struct
	return &createProductRequest, nil
}