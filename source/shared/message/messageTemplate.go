package message

type messageTemplateData struct {
	MessageName string
}

var messageTemplate = `package message

var {{ .MessageName }}ID = NextID()

type {{ .MessageName }} struct {
	id          uint64
	name        string
	MessengerID string // Set by the sender.

	/* KICKFYNE TODO:
	Complete this {{ .MessageName }} struct definition.
	Add the members that you want this message to have.
	*/

	Error        bool   // Set by the receiver in reply.
	Fatal        bool   // Set by the receiver in reply.
	ErrorMessage string // Set by the receiver in reply.
}

// New{{ .MessageName }} returns a *{{ .MessageName }} message.
func New{{ .MessageName }}(messengerID string) (msg *{{ .MessageName }}) {
	msg = &{{ .MessageName }}{
		id:          {{ .MessageName }}ID,
		name:        "{{ .MessageName }}",
		MessengerID: messengerID,

		/* KICKFYNE TODO:
		Complete New{{ .MessageName }} as needed.
		*/
	}
	return
}

// {{ .MessageName }} implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *{{ .MessageName }}) ID() (id uint64) {
	id = msg.id
	return
}

// Name returns the message's Name.
func (msg *{{ .MessageName }}) Name() (name string) {
	name = msg.name
	return
}

// AsInterface returns the message as an any.
func (msg *{{ .MessageName }}) AsInterface() (m any) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *{{ .MessageName }}) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	messengerID = msg.MessengerID
	return
}

`
