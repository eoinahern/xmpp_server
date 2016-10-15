package main

/**
* xmpp server prototype using sqlite
* and json as the messaging format.
* written by eoin ahern.
* need TLS and sasl
*
**/

import (
	"log"
	"net"
	"os"
	"xmpp_server/core"

)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("incorrect no. args  entered: %s", os.Args[0])
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5222")
	checkErr(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkErr(err)

	for {

		conn, err := listener.AcceptTCP()

		if err != nil {
			continue
		}

		go core.HandleConn(conn)
		defer conn.Close()

		//may need to set timeout on connection

	}

}

func checkErr(e error) {

	if e != nil {
		log.Print(e)
		os.Exit(1)
	}

}
