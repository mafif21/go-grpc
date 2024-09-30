package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func SayHello()  {
	h := basic.Person{
		Name: "Lele Maliki",
	}

	log.Println(&h)
}