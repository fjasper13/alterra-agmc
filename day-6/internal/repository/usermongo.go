package repository

import (
	"context"
	"day-6/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB_NAME       = "agmc-day2"
	DB_COLLECTION = "users"
)

type UserMongo interface {
	FindAllMongo(ctx context.Context) (*[]models.User, error)
	CreateMongo(ctx context.Context, data *models.User) (*models.User, error)
}

type userMongo struct {
	client *mongo.Client
	// db     *mongo.Database
}

func NewUserMongo(client *mongo.Client) *userMongo {
	return &userMongo{
		client,
	}
}

func (r *userMongo) FindAllMongo(ctx context.Context) (*[]models.User, error) {
	collection := r.client.Database(DB_NAME).Collection(DB_COLLECTION)

	result := []models.User{}

	getCursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for getCursor.Next((context.TODO())) {
		var elem models.User
		if err = getCursor.Decode(&elem); err == nil {
			result = append(result, elem)
		}
	}

	return &result, nil
}

func (r *userMongo) CreateMongo(ctx context.Context, user *models.User) (*models.User, error) {
	collection := r.client.Database(DB_NAME).Collection(DB_COLLECTION)

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	// insert := mongoResult.InsertedID
	// insertID := utils.ObjectIDHex(insert)
	// intID, err := strconv.ParseUint(insertID, 10, 32)
	// if err != nil {
	// 	return nil, err
	// }
	// user.ID = uint(intID)

	return user, nil
}
