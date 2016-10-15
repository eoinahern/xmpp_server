package messages

/**
* file for various message types
*
**/

var Message struct {
	from string
	to   string
	body string
}

var IQ struct {
	from string
	to   string
	body string
}

var Presence struct {
	from string
	to   string
	body string
}

var Stream struct {
}
