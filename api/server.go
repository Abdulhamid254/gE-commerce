package api

import (
	"encoding/json"
	"net/http"

	"github.com/Abdulhamid254/gggcommerce/store"
	"github.com/Abdulhamid254/gggcommerce/types"

	"github.com/anthdm/weavebox"
)



type ProductStorer interface {
	Insert(*types.Product) error
	GetById(string) (*types.Product, error)
}

type ProductHandler struct {
  store store.ProductStorer
}


func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler {
		store: pStore,
	}

}

func (h *ProductHandler) HandlePostProduct(c *weavebox.Context) error {
	productReq := &types.CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}
	// doing some validations
	product, err := types.NewProductFromRequest(productReq)
	if err != nil {
		return err
	}
	
	if err := h.store.Insert(c.Context, product); err != nil {
		return err
	}
	return c.JSON(http.StatusOK,product)
}

func (h *ProductHandler) HandleGetProductById(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU:"SHOE-1111"})
}