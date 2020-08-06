package main


/*
Setup



- Install protobuf libraries (github.com/golang/protobuf)
- Install plugin in IDE for editing .proto files
- Install package with protobuf compiler (go install google.golang.org/protobuf/cmd/protoc-gen-go)


TUTORIAL

- Form ur .proto file (message)
- compile it into a go file
	- $ protoc --go_out=. *.proto


*/


import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "kafkaProtobuf/pb"
	"log"
)



func main() {

	//Set up emails
	var emails []*pb.Person_Email
	personalEmail := &pb.Person_Email{
		Primary: "personal@gmail.com",
	}
	workEmail := &pb.Person_Email{
		Primary: "work@gmail.com",
	}
	emails = append(emails, personalEmail, workEmail)


	shane := &pb.Person{
		Name: "Shane",
		Age:  21,
		Emails: emails,
	}

	data, err := proto.Marshal(shane)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// unmarhsal byte array into an object we can modify and use
	unmarshalShane := &pb.Person{}
	err = proto.Unmarshal(data, unmarshalShane)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}


	fmt.Println(unmarshalShane.GetName())
	fmt.Println(unmarshalShane.GetEmails())

}