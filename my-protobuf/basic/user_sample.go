package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	a := basic.Address{
		Street:  "Jln Kyai Mojo No: 11",
		City:    "Jepara",
		Country: "Indonesia",
	}

	u := basic.User{
		Id:       99,
		Username: "apiipp_____",
		IsActive: true,
		Password: []byte("password"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"lelejepara@gmail.com", "afif21jepara@gmail.com"},
		Address:  &a,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	u2 := basic.User{
		Id:       99,
		Username: "apiipp_____",
		IsActive: true,
		Password: []byte("password"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"lelejepara@gmail.com", "afif21jepara@gmail.com"},
	}

	jsonBytes, _ := protojson.Marshal(&u2)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	exampleJson := `{
		"id": 1234,
		"username": "funny_monkey_420",
		"is_active": true,
		"password": "c3VwZXJfc2VjcmV0X2JhbmFuYQ==",
		"emails": [
			"monkey@banana.com",
			"chimpanzee.lol@jungle.org"
		],
		"gender": "GENDER_MALE"
	}`

	var p basic.User
	err := protojson.Unmarshal([]byte(exampleJson), &p)

	if err != nil {
		log.Fatalln("Got Error : ", err)
	}

	log.Println(p)
}
