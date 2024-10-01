package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUserGroup() {
	lele := basic.User{
		Id:       1,
		Username: "lelegoreng",
		IsActive: true,
		Password: []byte("lelegoreng"),
		Emails:   []string{"lelegoreng@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "Jalan Kyai Mojo",
			City:    "Jepara",
			Country: "Indonesia",
		},
	}

	dino := basic.User{
		Id:       2,
		Username: "aldinodino",
		IsActive: true,
		Password: []byte("aldd"),
		Emails:   []string{"aldino@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	ug := basic.UserGroup{
		GroupId:     999,
		GroupName:   "AKMJ",
		Roles:       []string{"Admin", "User"},
		Users:       []*basic.User{&lele, &dino},
		Description: "ini adalah grub akmj",
	}

	res, _ := protojson.Marshal(&ug)
	log.Println(string(res))
}
