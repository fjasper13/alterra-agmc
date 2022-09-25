package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDHex(ID interface{}) string {
	var objectID string
	switch v := ID.(type) {
	case string:
		objectID = v
	default:
		objectID = v.(primitive.ObjectID).Hex()
	}
	return objectID
}
