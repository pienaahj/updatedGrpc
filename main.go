package main

import (
	"fmt"
	"reflect"
	"time"

	pb "github.com/pienaahj/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// create a simple message
func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

// create a complex message
func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "My Name2"},
			{Id: 44, Name: "My Name3"},
			{Id: 45, Name: "My Name4"},
		},
	}
}

// create a enum message
func doEnum() *pb.Enumeration {
	return &pb.Enumeration{EyeColor: pb.EyeColor_EYE_COLOR_GREEN}
}

// create a oneOf message
func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type %v", x)
	}
}

// create a map message
func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId":  {Id: 42},
			"myId2": {Id: 43},
			"myId3": {Id: 44},
		},
	}
}

// evoke write and read to file functions
func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

// evoke the JSON functionality, encode to JSON
func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	return jsonString
}

// evoke the JSON functionality, decode from JSON
func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message) // make a message of the type that is passed in as a parameter
	fromJSON(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// fmt.Println("This should be an Id: ")
	// doOneOf(&pb.Result_Id{Id: 42})
	// fmt.Println("This should be a message: ")
	// doOneOf(&pb.Result_Message{Message: "a message"})
	// fmt.Println(doMap())
	// doFile(doSimple())
	// jsonString := doToJSON(doSimple())
	// message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	// fmt.Println(jsonString)
	// fmt.Println(message)

	// jsonString = doToJSON(doComplex())
	// message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	// fmt.Println(jsonString)
	// fmt.Println(message)

	lastUpdated := timestamppb.New(time.Now())
	newAddressbook := doAddressBook()
	pNumber1 := "0987654321"
	pType1 := pb.Person_PhoneType(2).Enum()
	pNumber2 := "1234567890"
	pType2 := pb.Person_PhoneType(1).Enum()
	phoneNumber1 := doPhoneNumber(pNumber1, pType1)
	phoneNumber2 := doPhoneNumber(pNumber2, pType2)
	phoneX := make([]*pb.Person_PhoneNumber, 0)
	phoneX = append(phoneX, phoneNumber1)
	newPerson := doPerson("name1", 42, "email@domain.com", phoneX, lastUpdated)
	newAddressbook = addToAddressBook(newPerson, newAddressbook)
	printAddressBook(newAddressbook)
	phoneX = append(phoneX, phoneNumber2)
	// phoneX := map[string]*pb.Person_PhoneType{pNumber: pType}
	// phones := doPhones(phoneX)
	newPerson = doPerson("name2", 43, "email1@domain.com", phoneX, lastUpdated)
	newAddressbook = addToAddressBook(newPerson, newAddressbook)
	printAddressBook(newAddressbook)

}
