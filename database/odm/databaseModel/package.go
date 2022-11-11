package databaseModel

import (
	"github.com/onecthree/owlchild/database/odm/collectionModel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// "context"
	
	// "fmt"
)

type Init struct {
	Client				*mongo.Client
	DatabaseName		string
}

func ( this *Init ) Collection( collectionName string ) *collectionModel.Init {
	collInit := collectionModel.Init {
		Client:				this.Client,
		DatabaseName:		this.DatabaseName,
		CollectionName:		collectionName,
		QueryColumn:		bson.M{},
		FindType:			"many",
		AssignModel:		false,
	};
	
	return &collInit;
}

// func ( this *Init ) Get( collectionName string ) *collectionModel.Init {
// 	client := this.Client;
// 	db := client.Database(this.DatabaseName);
// 	cursor, err := db.Find(context.TODO(), bson.D{});
	
// 	var results []bson.M;
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		fmt.Println(err);
// 	}
	
// 	fmt.Println(results);
// }