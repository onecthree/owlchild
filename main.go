package main

import(
	// "github.com/onecthree/owlchild/horizon"
	// . "github.com/onecthree/owlchild/facetype"
	"github.com/onecthree/owlchild/database/odm/model"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/kr/pretty"

	// "fmt"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

func main() {
	
	uri := "mongodb://simon:SimonProject135$!@localhost:27017";

	Users := model.SetCollection(uri, "DbTest", "CollectionTest");
	Users.FindOne( bson.M {
		"age" : "23",
	}).Get();

	var UsersSet UsersInterface;
	Users.Assign(&var);

}