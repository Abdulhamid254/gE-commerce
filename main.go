package main

import (
	"github.com/Abdulhamid254/gggcommerce/api"
	"github.com/anthdm/weavebox"
)

//weavebox is like a framework
//go get github.com/anthdm/weaveboxclear


func main(){
	app := weavebox.New()

	productHandler := &api.ProductHandler{}

	// app.Get("/product", func(*weavebox.Context) error {return nil})
		app.Get("/product", productHandler.HandleGetProduct)

	app.Serve(3001)
	
}