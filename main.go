package main

import(
	"github.com/onecthree/owlchild/horizon"
	. "github.com/onecthree/owlchild/facetype"

	"fmt"
)

func main() {

	AppModels := horizon.Config {
		Host:	"127.0.0.1:8007",
	}

	AppModels.Route( MapInterface {
		"UsersModel": func() {
			fmt.Println("Hello World?");
		},
	})

	AppModels.Listen();

}