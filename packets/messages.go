package messages

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
