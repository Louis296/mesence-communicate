package mongodb

import (
	"context"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DB *mongo.Database

func InitClient(url, database string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://"+url))
	if err != nil {
		return err
	}
	DB = client.Database(database)
	return nil
}

func SaveMessage(msg *pb.Msg) error {
	message := Message{
		Content:  msg.Data.Content,
		From:     msg.Data.From,
		To:       msg.Data.To,
		SendTime: util.TimeParse(msg.Data.SendTime),
	}
	_, err := DB.Collection("message").InsertOne(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}

func ListMessage(userA, userB string, offset, limit int64, startTime, endTime time.Time) ([]Message, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"From": userA, "To": userB},
			bson.M{"From": userB, "To": userA},
		},
		"SendTime": bson.M{"$gte": startTime, "$lte": endTime},
	}
	cur, err := DB.Collection("message").Find(context.Background(), filter, options.Find().SetSkip((offset-1)*limit).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	var res []Message
	if err := cur.All(context.Background(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
