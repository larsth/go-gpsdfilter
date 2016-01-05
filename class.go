package gpsdfilter

//Class is a type that contain a gpsd JSON document type, fx. "TPV"
type Class struct {
	Class string `json:"class"`
}
