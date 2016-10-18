package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"xmpp_server/packets"
)

type InnerMessage struct {
	Name       string `json:"name"`
	Second     string `json:"second"`
	Messagestr string `json:"message"`
}

type Message struct {
	MyMessage InnerMessage `json:"messagetype"`
}

/**
*  get type of message
*
***/

func GetType(message []byte) string {

	var holder interface{}
	err := json.Unmarshal(message, &holder)

	if err != nil {
		fmt.Println("error converting")
	}
	return determineTypeHelper(holder.(map[string]interface{}))
}

/**
*  checks tag on Inner message
* to decide what type of message has been sent
***/

func determineTypeHelper(obj map[string]interface{}) string {

	if _, ok := obj["messagetype"]; ok {
		return "message"
	} else if _, ok := obj["stream"]; ok {
		return "stream"
	} else if _, ok := obj["featurestream"]; ok {
		return "featurestream"
	} else {
		return "unknown"
	}

}

/**
*  unmarshals and returns Inner message
*
***/

func GetMessage(b []byte) InnerMessage {
	var mess Message
	err := json.Unmarshal(b, &mess)

	if err != nil {
		log.Fatal(err)
	}

	return mess.MyMessage
}

func GetStreamType(b []byte) messages.Stream {
	var str messages.Stream

	err := json.Unmarshal(b, &str)
	if err != nil {
		log.Fatal(err)
	}

	return str
}
