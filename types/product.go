package types

import "fmt"

const minProductNameLen = 3


type Product struct {
	ID string `bson:"_id,omitempty" json: "id"`
	SKU string `bson:"sku" json: "SKU"`
	Name string `bson:"name" json: "name"`
	Slug string `bson:"slug" json: "slug"`
	// Slug string `bson:"slug" json: "slug.omitempty"`
}


type CreateProductRequest struct {
    SKU string `json: "SKU"`
	Name string `json: "name"`
}

func NewProductFromRequest(req *CreateProductRequest) (*Product, error){
	if err := validateCreateProductRequest(req); err != nil {
		return nil,err
	}
	return &Product{
        SKU: req.SKU,
		Name: req.Name,
	},nil
}

func validateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product id to short!!")
	}
	return nil
}