package main

import(
	// "github.com/onecthree/owlchild/horizon"
	// . "github.com/onecthree/owlchild/facetype"
	"github.com/onecthree/owlchild/database/odm/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "github.com/kr/pretty"	

	// "fmt"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type UsersInterface struct {
	ID     		primitive.ObjectID 	`bson:"_id,omitempty"`
	Username 	string             	`bson:"username,omitempty"`
	PhoneNumber string				`bson:"phone_number,omitempty`
}

func main() {
	
	uri := "mongodb://simon:SimonProject135$!@localhost:27017";
	models := model.SetCollection(uri, "DbTest", "CollectionTest");
	models.Create(UsersInterface {
		Username:		"Sultan Ilham",
		PhoneNumber:	"081210981006",
	});

}