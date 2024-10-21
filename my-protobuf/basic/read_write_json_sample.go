package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
)

func WriteProtoToJson(msg proto.Message, fName string) {
	res, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalln("failed marshal proto")
	}

	if err := ioutil.WriteFile(fName, res, 0644); err != nil {
		log.Fatalln("failed write file proto")
	}
}

func ReadJsonToProto(fName string, dest proto.Message) {
	file, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("file not found")
	}

	if err := protojson.Unmarshal(file, dest); err != nil {
		log.Fatalln("failed marshal proto")
	}

}

func WriteToJsonSample() {
	uDummy := DummyExam()
	WriteProtoToJson(&uDummy, "user_contoh.json")
}

func ReadJsonSample() {
	dest := basic.User{}
	fName := "user_contoh.json"

	ReadJsonToProto(fName, &dest)

	log.Println(dest)
}
