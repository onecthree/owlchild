package collectionModel

import (
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"

	// "encoding/json"
	"context"

	"fmt"
)

type Init struct {
	Client				*mongo.Client
	DatabaseName		string
	CollectionName		string
	QueryColumn			bson.M
	FindType			string

	AssignModel			bool
	assignType			interface{}	

	result				bson.M
}

func ( this *Init ) Get() interface{} {
	client := this.Client;
	db := client.Database(this.DatabaseName).Collection(this.CollectionName);
	
	if( this.FindType == "many" ) {
		var result []bson.M;

		cursor, err := db.Find(context.TODO(), this.QueryColumn);

		if err = cursor.All(context.TODO(), &result); err != nil {
			fmt.Println(err);
		}

		// this.result = []byte(result);
		
		return result;
	}

	if( this.FindType == "one") {
		var result bson.M

		if err := db.FindOne(context.TODO(), this.QueryColumn).Decode(&result); err != nil {
			fmt.Println(err);
		}

		this.result = result;

		return result;
	}

	return nil;
}

func ( this *Init ) Create( dataModel bson.M ) {
	client := this.Client;
	collection := client.Database(this.DatabaseName).Collection(this.CollectionName)
	res, _ := collection.InsertOne(context.TODO(), dataModel );

	fmt.Println(res);
}

func ( this *Init ) Find( queryModel bson.M ) *Init {
	this.QueryColumn = queryModel;
	this.FindType = "many";

	return this;
}


func ( this *Init ) FindOne( queryModel bson.M ) *Init {
	this.QueryColumn = queryModel;
	this.FindType = "one";

	return this;
}

func ( this *Init ) Assign( modelInterface interface{} ) {
	if( this.FindType == "many") {

	}

	if( this.FindType == "one" ) {
		var err = bson.Unmarshal(this.result, &modelInterface);
		if err != nil {
			fmt.Println(err.Error());
			return;
		}
	}
}