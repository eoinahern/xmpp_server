package authentication

import "net"

func initialiseTLS(c *net.Conn) {

	//read conn.
	b := make([]byte, 80)
	int, err := c.Read(b)

	if err != nil {
		c.Close()
		println("connection closed")
	}

	//unmarshal and check version
	checkVersion(1)

}

func checkVersion(verno int) bool {

	if verno >= 1 {
		return true
	}

	return false

}
