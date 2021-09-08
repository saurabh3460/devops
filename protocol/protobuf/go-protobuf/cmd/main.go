package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "gitlab.com/saurabh3460/devops/protocol/protobuf/go-protobuf/dest"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{
		People: []*pb.Person{&p, &p},
	}

	// write to the file
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Fail to encode address book:", err)
	}
	if err := ioutil.WriteFile("test", out, 0644); err != nil {
		log.Fatalln("Faild to write address book:", err)
	}

	// read from the file
	in, err := ioutil.ReadFile("test")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	book = &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Fail to pars ")
	}
	fmt.Println(book)
}

// https://developers.google.com/protocol-buffers/docs/gotutorial
