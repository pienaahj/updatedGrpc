package main

import (
	"fmt"
	"strings"

	pb "github.com/pienaahj/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// check this page for detail: https://developers.google.com/protocol-buffers/docs/reference/go-generated

// readPerson reads a person from the addressbook on disk
func readPerson(fname string, pb proto.Message) {
	readFromFile(fname, pb)
}

// writeAddressBook writes the addressbook to disk
func writeAddressBook(fname string, book *pb.AddressBook) {
	writeToFile(fname, book)
}

// doPhoneNumber creates the phoneNumber
func doPhoneNumber(pNumber string, pType *pb.Person_PhoneType) *pb.Person_PhoneNumber {
	return &pb.Person_PhoneNumber{
		Number: pNumber,
		Type:   *pType,
	}
}

// doPhones created the phones slice
func doPhones(phones map[string]*pb.Person_PhoneType) []*pb.Person_PhoneNumber {
	newPhones := make([]*pb.Person_PhoneNumber, len(phones))
	for _, v := range phones {
		newPhones = append(newPhones, doPhoneNumber(v.String(), v.Enum()))
	}
	return newPhones
}

//	func doPhones(phones map[string]*pb.Person_PhoneType) []*pb.Person_PhoneNumber {
//		newPhones := make([]*pb.Person_PhoneNumber, len(phones))
//		for _, v := range phones {
//			newPhones = append(newPhones, doPhoneNumber(v.String(), v.Enum()))
//		}
//		return newPhones
//	}
//
// doPerson returns a person
func doPerson(name string, id int32, email string, phones []*pb.Person_PhoneNumber, lastUpdated *timestamppb.Timestamp) *pb.Person {
	return &pb.Person{
		Name:        name,
		Id:          id,
		Email:       email,
		Phones:      phones,
		LastUpdated: lastUpdated,
	}
}

// doAddressbook creates the addressbook
func doAddressBook() *pb.AddressBook {
	return &pb.AddressBook{}
}

// addToAddressBook adds a person to the addressbook
func addToAddressBook(person *pb.Person, people *pb.AddressBook) *pb.AddressBook {
	// make a copy of addressbook people
	newPeople := people.People
	newPeople = append(newPeople, person)
	//  create new addressBook
	newAddressbook := &pb.AddressBook{
		People: newPeople,
	}
	return (newAddressbook)
}

// printAddressBook prints the addresses in addressbook
func printAddressBook(book *pb.AddressBook) {
	fmt.Println(strings.Repeat("*", 52))
	for i, person := range book.People {
		fmt.Printf("Person(%d): %s, Id: %d, Email: %s\n", i+1, person.Name, person.Id, person.Email)
		fmt.Println("Phones: ")
		for p, phone := range person.Phones {
			fmt.Printf("Phone(%d): Number : %s, Type: %v\n", p+1, phone.Number, phone.Type)
		}
		fmt.Printf("Last updated: %v\n", person.LastUpdated.AsTime())
		fmt.Println(strings.Repeat("*", 20), "end person", strings.Repeat("*", 20))
	}
	fmt.Println()
}

/*
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` // Unique ID number for this person.
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones      []*Person_PhoneNumber  `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

Our address book file is just one of these.
type AddressBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

*/
