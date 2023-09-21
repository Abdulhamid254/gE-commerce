package main

import (
	"context"

	"github.com/Abdulhamid254/gggcommerce/api"
	"github.com/Abdulhamid254/gggcommerce/store"
	"github.com/anthdm/weavebox"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//weavebox is like a framework
//go get github.com/anthdm/weaveboxclear


func main(){
	app := weavebox.New()
	// box here is like a super route
	adminRoute := app.Box("/admin")



	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
	    panic(err)
      }

	// productHandler := &api.ProductHandler{}

	
	productStore := store.NewMongoProductStore(client.Database(("gggcommerce")))
	productHandler := api.NewProductHandler(productStore)

	// app.Get("/product", func(*weavebox.Context) error {return nil})
		adminRoute.Get("/product/:id", productHandler.HandleGetProductById)
		adminRoute.Post("/product", productHandler.HandlePostProduct)

	app.Serve(3001)
	
}