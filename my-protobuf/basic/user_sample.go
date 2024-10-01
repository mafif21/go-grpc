package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"
)

func BasicUser() {
	a := basic.Address{
		Street:  "Jln Kyai Mojo No: 11",
		City:    "Jepara",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:   87649939.0844,
			Longtitude: 477474883.980,
		},
	}

	//commChannel := randomCommunication()

	u := basic.User{
		Id:       99,
		Username: "apiipp_____",
		IsActive: true,
		Password: []byte("password"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"lelejepara@gmail.com", "afif21jepara@gmail.com"},
		Address:  &a,
		//CommunicationChannel: &commChannel,
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

func randomCommunication() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{Email: "lele@gmail.com"}
	socialMedia := basic.SocialMedia{SocialMediaUsername: "apiipp_____"}
	instantMessaging := basic.InstantMessaging{Phone: "0839774993"}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalingAnyKnown() {
	socialMedia := basic.SocialMedia{SocialMediaUsername: "apiipp_____"}

	var a anypb.Any

	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalingAnyUnknown() {
	randomCom := randomCommunication()

	var unmarshaled protoreflect.ProtoMessage

	unmarshalNew, err := randomCom.UnmarshalNew()
	if err != nil {
		return
	}

	unmarshaled = unmarshalNew

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())
}

func BasicUnmarshalAnyIs() {
	a := randomCommunication()
	mail := basic.PaperMail{}

	if a.MessageIs(&mail) {
		if err := a.UnmarshalTo(&mail); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&mail)
		log.Println(string(jsonBytes))
	} else {
		log.Println(a.TypeUrl)
	}
}

func BasicOneOf() {
	//socialMedia := basic.SocialMedia{SocialMediaUsername: "apiipp_____"}
	//comm := basic.User_SocialMedia{&socialMedia}

	instantMessaging := &basic.InstantMessaging{Phone: "0839774993"}
	comm := basic.User_InstantMessaging{instantMessaging}

	u := basic.User{
		Id:       44,
		Username: "lele",
		IsActive: true,
		Password: []byte("jdbdbbbdd"),
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "Jln Kyai Mojo",
			City:    "Jepara",
			Country: "Indonesia",
		},
		CommunicationChannel: &comm,
		Emails:               []string{"muhammadafifjpr@gmail.com"},
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
