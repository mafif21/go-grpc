package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)


func BasicUser()  {
	u := basic.User{
		Id: 99,
		Username: "apiipp_____",
		IsActive: true,
		Password: []byte("password"),
	}

	log.Println(&u)
}