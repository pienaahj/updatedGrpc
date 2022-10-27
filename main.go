package main

import (
	"fmt"
	"reflect"

	pb "github.com/pienaahj/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

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

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{EyeColor: pb.EyeColor_EYE_COLOR_GREEN}
}

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

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId":  {Id: 42},
			"myId2": {Id: 43},
			"myId3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
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
	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	jsonString = doToJSON(doComplex())
	message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)
}
