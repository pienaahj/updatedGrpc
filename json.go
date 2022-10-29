package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Encode  message to JSON
func toJSON(pb proto.Message) string {
	// create a marshal option
	option := protojson.MarshalOptions{
		Multiline: true, // pretty print json
	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(out)
}

// Decode message from JSON
func fromJSON(in string, pb proto.Message) {
	// create a unmarshal option
	option := protojson.UnmarshalOptions{ // will discard the field if the field does not exist
		DiscardUnknown: true,
	}
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Couldn't unmarshal from JSON", err)
	}
}
