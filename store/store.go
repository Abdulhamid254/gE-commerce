// where we have our interfaces

package store

import (
	"context"

	"github.com/Abdulhamid254/gggcommerce/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetById(context.Context,string) ( *types.Product, error)
}