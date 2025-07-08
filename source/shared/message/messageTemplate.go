package message

type messageTemplateData struct {
	MessageStructName string
}

var messageTemplate = `package message

var {{ .MessageStructName }}ID = "{{ .MessageStructName }}"

type {{ .MessageStructName }} struct {
	MessengerID string // Set by the sender.

	/* KICKFYNE TODO:
	Complete this {{ .MessageStructName }} struct definition.
	Add the members that you want this message to have.
	*/

	Error        bool   // Set by the receiver in reply.
	Fatal        bool   // Set by the receiver in reply.
	ErrorMessage string // Set by the receiver in reply.
}

// New{{ .MessageStructName }} returns a *{{ .MessageStructName }} message.
func New{{ .MessageStructName }}(messengerID string) (msg *{{ .MessageStructName }}) {
	msg = &{{ .MessageStructName }}{
		MessengerID: messengerID,

		/* KICKFYNE TODO:
		Complete New{{ .MessageStructName }} as needed.
		*/
	}
	return
}

// {{ .MessageStructName }} implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *{{ .MessageStructName }}) ID() (id string) {
	id = {{ .MessageStructName }}ID
	return
}

// Name returns the message's Name.
func (msg *{{ .MessageStructName }}) Name() (name string) {
	name = {{ .MessageStructName }}ID
	return
}

// AsInterface returns the message as an any.
func (msg *{{ .MessageStructName }}) AsInterface() (m any) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *{{ .MessageStructName }}) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	messengerID = msg.MessengerID
	return
}

`
