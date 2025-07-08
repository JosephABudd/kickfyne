package message

const (
	initFileName = "init.go"
)

var initTemplate = `package message

var InitID = "Init"

type Init struct {
	name         string
	Message      string // to front
	Error        bool   // to front
	Fatal        bool   // to front
	ErrorMessage string // to front
}

// NewInit constructs a new NewInit message.
func NewInit() (msg *Init) {
	msg = &Init{}
	return
}

// Init implements the MSGer interface with ID and AsInterface.

// ID returns the message's id.
func (msg *Init) ID() (id string) {
	id = InitID
	return
}

// Name returns the message's name.
func (msg *Init) Name() (name string) {
	name = InitID
	return
}

// FatalError returns if there was a fatal error and it's message.
func (msg *Init) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	// Init has no messengerID because its sent by the front-end not a screen.
	return
}

// AsInterface returns msg as an interface{}.
func (msg *Init) AsInterface() (m interface{}) {
	m = msg
	return
}

`
