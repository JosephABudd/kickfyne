package message

const (
	chansFileName = "txrxchans.go"
)

var chansTemplate = `package message

type MSGer interface {
	ID() (id string)
	Name() (name string)
	FatalError() (fatal bool, message, messengerID string)
	AsInterface() (msg interface{})
}

var FrontEndToBackEnd = make(chan MSGer, 255)
var BackEndToFrontEnd = make(chan MSGer, 255)

func IsValidID(id string) (isvalid bool) {
	isvalid = len(id) > 0
	return
}

`
