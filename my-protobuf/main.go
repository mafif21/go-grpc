package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct{}

func (writer *logWriter) Write(byte []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(byte)))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.WriteToFileSample()
	//basic.ReadFileSample()
	//basic.UserWithSkill()
	// basic.SayHello()
	//basic.BasicUser()
	//basic.BasicUserGroup()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser()
	//jobsearch.JobSearchGeneral()
	//jobsearch.JobSeacrhSoftwareEngineer()
	//basic.BasicUnmarshalingAnyKnown()
	//basic.BasicUnmarshalingAnyUnknown()
	//basic.BasicUnmarshalAnyIs()
	//basic.BasicOneOf()
	//basic.WriteToJsonSample()
	//basic.ReadJsonSample()
	//basic.BasicWriteUserContentV1()
	//basic.BasicReadUserContentV1()
	//basic.BasicWriteUserContentV2()
	//basic.BasicReadUserContentV2()
	//basic.BasicWriteUserContentV3()
	//basic.BasicReadUserContentV3()
	//basic.BasicWriteUserContentV4()
	//basic.BasicReadUserContentV4()
	basic.BasicReadUserPayment()
}
