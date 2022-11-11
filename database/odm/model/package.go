package model

import (
	"github.com/onecthree/owlchild/database/odm/collectionModel"
	"github.com/onecthree/owlchild/database/odm/databaseModel"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"context"

	"fmt"
	"log"
)

type Init struct {
	Client			*mongo.Client;
}

func setClient( uri string ) *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri));

	if err != nil {
		fmt.Println(err)
	}
	
	return client;
}

func Default(uri string) *Init {
	set := Init{
		Client:		setClient(uri),
	};

	return &set;
}

func SetDatabase( uri string, databaseName string ) *databaseModel.Init {
	dbInit := databaseModel.Init {
		Client:			setClient(uri),
		DatabaseName: 	databaseName,
	};

	return &dbInit;
}

func SetCollection( uri string, databaseName string, collectionName string ) *collectionModel.Init {
	collInit := collectionModel.Init {
		Client:			setClient(uri),
		DatabaseName: 	databaseName,
		CollectionName: collectionName,
		QueryColumn:	bson.M{},
		FindType:		"many",
		AssignModel:	false,
	};

	return &collInit;
}









func (this *Init) Database( databaseName string ) *databaseModel.Init {
	dbInit := databaseModel.Init {
		Client:			this.Client,
		DatabaseName: 	databaseName,
	};

	return &dbInit;
}

func ( this *Init ) Get() []string {
	dbs, err := (this.Client).ListDatabaseNames(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	return dbs;
}