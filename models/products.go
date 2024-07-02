package models

import (
	"context"
	"log"
	"time"

	"github.com/anhhuy1010/cms-menu/database"
	"github.com/anhhuy1010/cms-menu/helpers/util"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson"

	"github.com/anhhuy1010/cms-menu/constant"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Products struct {
	Uuid        string    `json:"uuid,omitempty" bson:"uuid"`
	Price       int       `json:"price,omitempty" bson:"price"`
	Image       string    `json:"image" bson:"image"`
	Name        string    `json:"name,omitempty" bson:"name"`
	Sequence    int       `json:"sequence" bson:"sequence"`
	Quantity    int       `json:"quantity" bson:"quantity"`
	Description string    `json:"description" bson:"description"`
	Gallery     []string  `json:"gallery" bson:"gallery"`
	IsActive    int       `json:"is_active" bson:"is_active"`
	IsDelete    int       `json:"is_delete" bson:"is_delete"`
	StartDate   time.Time `json:"start_date" bson:"start_date"`
	EndDate     time.Time `json:"end_date" bson:"end_date  "`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *Products) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("products")
}

func (u *Products) Find(conditions map[string]interface{}, opts ...*options.FindOptions) ([]*Products, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	cursor, err := coll.Find(context.TODO(), conditions, opts...)
	if err != nil {
		return nil, err
	}

	var product []*Products
	for cursor.Next(context.TODO()) {
		var elem Products
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		product = append(product, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return product, nil
}

func (u *Products) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*Products, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var product []*Products
	for cursor.Next(context.TODO()) {
		var elem Products
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("[Decode] PopularCuisine:", err)
			log.Println("-> #", elem.Uuid)
			continue
		}

		product = append(product, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return product, nil
}

func (u *Products) Distinct(conditions map[string]interface{}, fieldName string, opts ...*options.DistinctOptions) ([]interface{}, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	values, err := coll.Distinct(context.TODO(), fieldName, conditions, opts...)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (u *Products) FindOne(conditions map[string]interface{}) (*Products, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *Products) Insert() (interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (u *Products) InsertMany(menu []interface{}) ([]interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertMany(context.TODO(), menu)
	if err != nil {
		return nil, err
	}

	return resp.InsertedIDs, nil
}

func (u *Products) Update() (int64, error) {
	coll := u.Model()

	condition := make(map[string]interface{})
	condition["uuid"] = u.Uuid

	u.UpdatedAt = util.GetNowUTC()
	updateStr := make(map[string]interface{})
	updateStr["$set"] = u

	resp, err := coll.UpdateOne(context.TODO(), condition, updateStr)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Products) UpdateByCondition(condition map[string]interface{}, data map[string]interface{}) (int64, error) {
	coll := u.Model()

	resp, err := coll.UpdateOne(context.TODO(), condition, data)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Products) UpdateMany(conditions map[string]interface{}, updateData map[string]interface{}) (int64, error) {
	coll := u.Model()
	resp, err := coll.UpdateMany(context.TODO(), conditions, updateData)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Products) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
