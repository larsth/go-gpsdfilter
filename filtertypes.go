package gpsdfilter

//Type descripes what to do with a gpsd JSON document: Unknown, Log, Use, Ignore
type Type byte

const (
	//TypeUnknown means that this is an unknown gpsd JSON document
	TypeUnknown Type = 0
	//TypeLog means that this gpsd JSON document should only be logged
	TypeLog Type = 1
	//TypeParse means that this gpsd JSON document should be parsed, and
	//the parsed content should be used
	TypeParse Type = 2
	//TypeIgnore means that this gpsd JSON document should be ignored
	TypeIgnore Type = 3
)
