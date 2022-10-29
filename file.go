package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

// writeFile is a helper function to write a proto message to a file
func writeToFile(fname string, pb proto.Message) {
	// marshal the message
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't write to bytes", err)
		return
	}

	// write the message to the file
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}

	fmt.Println("Data written!")

}

// readFromFile is a helper function to read a proto message from a file
func readFromFile(fname string, pb proto.Message) {

	// Read the message from the file
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read from file", err)
		return
	}
	//  unmarshal the message
	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal", err)
		return
	}

}
