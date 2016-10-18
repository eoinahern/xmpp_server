package messages

// message steam

type Stream struct {
	Inner InnerStream `json:"stream"`
}

type InnerStream struct {
	From    string `json:"from"` //myid prob email
	To      string `json:"to"`
	Id      string `json:"id"` //This attribute is a unique identifier created by the receiving entity to function as a session key for the initiating entity's streams with
	Version string `json:"version"`
	Lang    string `json:"lang"`
}

func NewStream() *Stream {
	stream := new(Stream)
	return stream
}

func (s *Stream) CreateKey() {
	s.Inner.Id = "12345"
}

/*
* features stream. sent after initial presence message.
*
**/

func NewFeatures() *FeaturesStream {
	fstream := new(FeaturesStream)
	return fstream
}

type FeaturesStream struct {
	Inner InnerFStream `json:"featurestream"`
}

type InnerFStream struct {
	Saslinit    bool     `json:"sasl"`
	Tlsinit     bool     `json:"tls"`
	Encryptypes []string `json:"encrypttypes"`
}
