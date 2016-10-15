package messages

/**
* file for various message types
*
**/

//stanza

type Message struct {
	from string
	to   string
	body string
}

type IQ struct {
	from string
	to   string
	body string
}

type Presence struct {
	from string
	to   string
	body string
}

// stream

/*type Stream struct {
	Inner InnerStream `json:"stream"`
}

type InnerStream struct {
	From    string `json:"from"` //myid prob email
	To      string `json:"to"`
	Id      string `json:"id"`  //This attribute is a unique identifier created by the receiving entity to function as a session key for the initiating entity's streams with
	Version string `json:"version"`
	Lang    string `json:"lang"`
}*/

type Feature struct {
}
