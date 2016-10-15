package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

/*
*
* create correct type
**/

func TestConverter(t *testing.T) {

	teststream := []byte(`{ "streamtype": { "name": "eoin", "second": "ahern"}}`)
	testmessage := []byte(`{ "messagetype": {"name": "eoin", "second": "ahern", "message": "hellothere!!!"}}`)

	typestr := GetType(teststream)
	if typestr != "stream" {
		t.Log("stream tyep not recognised")
		t.Fail()
	}

	typestr2 := GetType(testmessage)
	if typestr2 != "message" {
		t.Log("message tyep not recognised")
		t.Fail()
	}

}

/*
*
* create obj from byte array
**/

func TestCreateObj(t *testing.T) {

	teststream := []byte(`{ "streamtype": { "name": "eoin", "second": "ahern"}}`)
	testmessage := []byte(`{ "messagetype": {"name": "eoin", "second": "ahern", "message": "h"}}`)

	tt := GetType(testmessage)
	fmt.Println(tt)

	if tt == "message" {
		var messi Message
		err := json.Unmarshal(testmessage, &messi)

		if err != nil {
			log.Fatal(err)
			t.Error("failed type check")
		}

		fmt.Println(messi)
		fmt.Println(messi.MyMessage.Name)
		fmt.Println(messi.MyMessage.Second)

	}

	msg := GetMessage(testmessage)
	stream := GetStreamType(teststream)

	if msg.Name == "" || msg.Second == "" {
		t.Error("test message failed")
		t.Fail()
	}

	if stream.Name == "" || stream.Second == "" {
		t.Error("test stream failed")
		t.Fail()
	}

	fmt.Println(msg)
	fmt.Println(stream)

}
