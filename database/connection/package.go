package connection

import(
	"fmt"
)

type Init struct {
	Uri			string
}

func ( *Init ) Ping() {
	fmt.Println("Pong!");
}
