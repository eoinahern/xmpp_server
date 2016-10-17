package core

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"xmpp_server/messages"
	"xmpp_server/utils"
)

/**
*
* This will be initialization process
* going to need a DB for roster,
* connected users etc.
**/

func HandleConn(conn net.Conn) {

	by := make([]byte, 128)
	i, err := conn.Read(by)

	if err != nil {
		println("error reading message")
	}

	if utils.GetType(by[0:i]) != "stream" {
		println("not stream type?")
		os.Exit(1)
	}

	stream := utils.GetStreamType(by[0:i])
	bytereturnstr := CreateReturnStream(&stream)
	conn.Write(bytereturnstr)

	defer conn.Close()

}

func CreateReturnStream(stream *messages.Stream) []byte {

	stream.CreateKey()
	to := stream.Inner.To

	stream.Inner.To = stream.Inner.From
	stream.Inner.From = to

	streamarr, err := json.Marshal(stream)
	if err != nil {
		fmt.Println("error creating stream to return")
	}

	return streamarr

}

func readMessage() {

}
