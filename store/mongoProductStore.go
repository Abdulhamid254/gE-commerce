package store

import (
	"context"

	"github.com/Abdulhamid254/gggcommerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)




type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

// ​Btw, when a function uses a context, the function that uses the context should also have it in the function arguments. This way you can easily chain them.


func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		coll: "products",
	}
}
// The problem is that you can set anything as ID in mongodb.Mongo driver has no idea what you set 
// so it returns an interrface.The autogenerated ID whic is why we use of type primitive.ObjectID so that you can 
// safely cast without checks, and use Hex()


func (s *MongoProductStore) Insert(ctx context.Context, p *types.Product) error {
	res, err := s.db.Collection(s.coll).InsertOne(ctx, p)
	if err != nil {
		return err
	}
	// LIKE CONSOLE LOG FMT
	//fmt.Printf("+>>>>>%+v\n",res)
	// p.ID = res.InsertedID.(string)
	p.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (s *MongoProductStore) GetById(ctx context.Context, id string) (*types.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.db.Collection(s.coll).FindOne(ctx, bson.M{"_id": objID})
		p        = &types.Product{}
		err      = res.Decode(p)
	)
	return p, err
}