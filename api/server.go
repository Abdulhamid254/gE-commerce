package api

import (
	"net/http"

	"github.com/Abdulhamid254/gggcommerce/types"

	"github.com/anthdm/weavebox"
)

type ProductHandler struct {

}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU:"SHOE-1111"})
}