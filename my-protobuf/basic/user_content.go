package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 764,
		Slug:          "/this-is-v1",
		//Title:         "Holiday Gengs",
		//HtmlContent: "<h1>Hello World</h1>",
		//AuthorId:      34,
	}

	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")
	uc := basic.UserContent{}

	ReadFileToProto("user_content_v1.bin", &uc)

	log.Println(uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 764,
		Slug:          "/this-is-v2",
		//Title:         "Holiday Gengs",
		//HtmlContent: "<h1>Hello World</h1>",
		//AuthorId:      34,
		//Category: "reels",
	}

	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")
	uc := basic.UserContent{}

	ReadFileToProto("user_content_v2.bin", &uc)

	log.Println(uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 764,
		Slug:          "/this-is-v3",
		//HtmlContent:   "<h1>Hello World</h1>",
		//Category:      "reels",
	}

	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read V3")
	uc := basic.UserContent{}

	ReadFileToProto("user_content_v3.bin", &uc)

	log.Println(uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV4() {
	uc := basic.UserContent{
		UserContentId: 764,
		Slug:          "/this-is-v4",
		//Rating:        3,
	}

	WriteProtoToFile(&uc, "user_content_v4.bin")
}

func BasicReadUserContentV4() {
	log.Println("Read V4")
	uc := basic.UserContent{}

	ReadFileToProto("user_content_v4.bin", &uc)

	log.Println(uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}
