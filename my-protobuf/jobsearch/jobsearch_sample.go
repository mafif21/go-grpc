package jobsearch

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/first"
	"my-protobuf/protogen/jobsearch"
	"my-protobuf/protogen/second"
)

func JobSearchGeneral() {
	jg := jobsearch.JobCandidate{
		JobCandidateId: 83,
		Application: &first.Application{
			Name:   "Ratu Andri Yuliani",
			Age:    22,
			Skills: []string{"Excel", "Data Analysis", "SPSS"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&jg)
	log.Println(string(jsonBytes))
}

func JobSeacrhSoftwareEngineer() {
	jse := jobsearch.SoftwareEngineerJob{
		SoftwareEngineerId: 54,
		Application: &second.Application{
			Name:                "Muhammad Afif",
			Age:                 22,
			Skills:              []string{"Load Testing", "Rest API", "Auth"},
			ProgrammingLanguage: []string{"Node JS", "Python", "Golang"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&jse)
	log.Println(string(jsonBytes))
}
