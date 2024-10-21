package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
)

func DummyExam() basic.User {
	address := basic.Address{
		Street:  "Jalan Kyai Mojo",
		City:    "Jepara",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:   83768493,
			Longtitude: -3984,
		},
	}

	skills := map[string]uint32{
		"football": 2,
		"design":   1,
		"coding":   5,
	}

	return basic.User{
		Id:        1,
		Username:  "apiipp",
		IsActive:  true,
		Password:  []byte("password"),
		Gender:    basic.Gender_GENDER_MALE,
		Address:   &address,
		UserSkill: skills,
		Emails:    []string{"lelejepara@gmail.com", "afif21jepara@gmail.com"},
	}
}

func WriteProtoToFile(msg proto.Message, fName string) {
	res, err := proto.Marshal(msg)

	if err != nil {
		log.Fatalln("failed marshal proto")
	}

	if err := ioutil.WriteFile(fName, res, 0644); err != nil {
		log.Fatalln("failed write file proto")
	}
}

func ReadFileToProto(fName string, dest proto.Message) {
	file, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("file not found")
	}

	if err := proto.Unmarshal(file, dest); err != nil {
		log.Fatalln("failed marshal proto")
	}

}

func WriteToFileSample() {
	uDummy := DummyExam()
	WriteProtoToFile(&uDummy, "user_contoh.bin")
}

func ReadFileSample() {
	dest := basic.User{}
	fName := "user_contoh.bin"

	ReadFileToProto(fName, &dest)

	marshal, _ := protojson.Marshal(&dest)
	log.Println(string(marshal))
}
