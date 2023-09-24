package types

import (
	"fmt"
	"strings"
)

const minProductNameLen = 3


type Product struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
	// Slug string `bson:"slug" json: "slug.omitempty"`
}


type CreateProductRequest struct {
    SKU  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductFromRequest(req *CreateProductRequest) (*Product, error) {
	if err := validateCreateProductRequest(req); err != nil {
		return nil, err
	}

	parts := strings.Split(strings.ToLower(req.Name), " ")
	slug := strings.Join(parts, "-")
	//slug := append(parts,"-")


	return &Product{
		SKU:  req.SKU,
		Name: req.Name,
		Slug: slug,
	}, nil
}


// veryfyingthe payload
func validateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.SKU) < 3 {
		return fmt.Errorf("the SKU of the product is to short")
	}
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product is to short")
	}
	return nil
}