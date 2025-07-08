package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FRemoveAccordionItemTemplateData struct {
	PackageName       string
	MessageStructName string
	ImportPrefix      string
	Funcs             _utils_.Funcs
}

const (
	B2FRemoveAccordionItemTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

var {{ .MessageStructName }}ID = "{{ .MessageStructName }}"

// This goes front to back to front.
// This message is telling the front-end that a accordionItem needs to be removed from a accordion.
// The front-end must return the message to indicate success or failure.
type {{ .MessageStructName }} struct {
	// These are sent from the front to the back.
	AccordionMessengerID string
	AccordionItemMessengerID string

	// These are sent from the back to the front
	Error        bool   // Back to front.
	Fatal        bool   // Back to front.
	ErrorMessage string // Back to front.
}

// New{{ .MessageStructName }} constructs a *{{ .MessageStructName }} message for the back-end to send to the front-end.
func New{{ .MessageStructName }}(
	accordionMessengerID string,
	accordionItemMessengerID string,
) (msg *{{ .MessageStructName }}) {
	msg = &{{ .MessageStructName }}{
		AccordionMessengerID: accordionMessengerID,
		AccordionItemMessengerID:    accordionItemMessengerID,
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

// FrontendMessengerID returns the id of the front-end screen.
func (msg *{{ .MessageStructName }}) FrontendMessengerID() (id string) {
	id = msg.AccordionItemMessengerID
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
	messengerID = msg.AccordionItemMessengerID
	return
}

`
)
