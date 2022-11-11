package horizon

import (
	. "github.com/onecthree/owlchild/facetype"
	
	"net"
	"encoding/binary"

	"log"
	"fmt"
)

type Config struct {
	Host			string
	route 			MapInterface
}

func ( this *Config ) Route( routelist MapInterface ) {
	this.route = routelist;
}

func ( this *Config ) Listen() {
	instance, err := net.Listen("tcp4", this.Host)

	if err != nil {
			fmt.Println("Error caught:", err);
			return
	}
	
	for {
		connection, err := instance.Accept()

		if err != nil {
			log.Println(err)
			return
		}
		
		go func() {
			var data []byte;
			data = this.messageParse(connection);
			serialize_data := string(data);
			this.route[serialize_data].(func())();
		}();
	}
}	

func ( this *Config ) messageParse(connection net.Conn) []byte {
	var result []byte

	for {   
			// Receive the length of actual message
			var size int32
			err := binary.Read(connection, binary.LittleEndian, &size)

			if err != nil {
					fmt.Println("err:", err)
			}

			// Send message callback for delivered lenght message
			connection.Write([]byte("OK 200"))

			// Receive blank message for EOF.
			bufferBytes := make([]byte, 0);

			len, err := connection.Read(bufferBytes)

			if err != nil {
					panic(err); break
			};
			
			result = bufferBytes[:len] 

			// Receive actually message for End-of-connection.
			bufferBytesf := make([]byte, size)

			lenf, err := connection.Read(bufferBytesf)
				
			if err != nil {
					panic(err); break
			}

			// passing the byte of message to function
			result = bufferBytesf[:lenf]
			
			break
	}
	
	return result
}
