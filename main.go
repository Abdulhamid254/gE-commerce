package main

import (
	"context"
	"fmt"

	"net/http"

	"github.com/Abdulhamid254/gggcommerce/api"
	"github.com/Abdulhamid254/gggcommerce/store"
	"github.com/anthdm/weavebox"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//weavebox is like a framework
//go get github.com/anthdm/weaveboxclear

//custom error

func handleAPIError(ctx *weavebox.Context, err error) {
	fmt.Println("API error:", err)
	ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}


func main(){
	app := weavebox.New()
	app.ErrorHandler = handleAPIError
	// box here is like a super route
   // protecting the adminRoute wwith middleware
	adminMW := &api.AdminAuthMiddleware{}
	adminRoute := app.Box("/admin")
    adminRoute.Use(adminMW.Authenticate)



// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		fmt.Printf("%T %+v", err, err)
// 		panic(err)
// 	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

// client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
//    if err != nil {
// 	log.Fatal(err)
//    }
//    ctx := context.Background()
//    err = client.Connect(ctx)
//    if err != nil {
// 	log.Fatal(err)
//    }
//    defer client.Disconnect(ctx)

	// productHandler := &api.ProductHandler{}

	
	productStore := store.NewMongoProductStore(client.Database("gggcommerce"))
	productHandler := api.NewProductHandler(productStore)

	// app.Get("/product", func(*weavebox.Context) error {return nil})
	//adminproduct
	    adminProductRoute := adminRoute.Box("/product")
		adminProductRoute.Get("/:id", productHandler.HandleGetProductById)
		adminProductRoute.Get("/", productHandler.HandleGetProducts)
		adminProductRoute.Post("/", productHandler.HandlePostProduct)

	app.Serve(3001)
	
}