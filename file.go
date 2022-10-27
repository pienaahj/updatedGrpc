package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't write to bytes", err)
		return
	}

	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}

	fmt.Println("Data written!")

}

func readFromFile(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read from file", err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal", err)
		return
	}

}
