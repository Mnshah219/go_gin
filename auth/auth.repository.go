package auth

import (
	"context"

	"github.com/mnshah219/go_gin/auth/dto"
	"github.com/mnshah219/go_gin/auth/schema"
	"github.com/mnshah219/go_gin/common"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	client := common.GetClient()
	return client.Database(common.DB_NAME).Collection(schema.USER_TABLE)
}
func createUser(user dto.SignupRequestDto) (*mongo.InsertOneResult, error) {
	coll := getCollection()
	result, err := coll.InsertOne(context.TODO(), user)
	return result, err
}

func findOneUser(filter interface{}) schema.User {
	coll := getCollection()
	result := coll.FindOne(context.TODO(), filter)
	user := schema.User{}
	result.Decode(&user)
	return user
}
