package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FAddAccordionItemTemplateData struct {
	PackageName       string
	MessageStructName string
	ImportPrefix      string
	Funcs             _utils_.Funcs
}

const (
	B2FAddAccordionItemTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

var {{ .MessageStructName }}ID = "{{ .MessageStructName }}"

// This goes back to front to back.
// It declares that the back-end needs the front-end to add a new accordionItem in the {{ .PackageName }} Accordion.
type {{ .MessageStructName }} struct {
	// These are sent from the back to the front.
	AccordionMessengerID string
	AccordionItem        AccordionItem
	AccordionItemLabel   string
	PanelsState          PanelsState

	// These are sent from the front to the back.
	AccordionItemMessengerID string
	Error        bool
	Fatal        bool
	ErrorMessage string
}

// New{{ .MessageStructName }} constructs a *{{ .MessageStructName }} message for the back-end to send to the front-end.
func New{{ .MessageStructName }}(
	accordionMessengerID string,
	accordionItem AccordionItem,
	accordionItemLabel string,
	panelsState PanelsState,
) (msg *{{ .MessageStructName }}) {
	msg = &{{ .MessageStructName }}{
		AccordionMessengerID: accordionMessengerID,
		AccordionItem:        accordionItem,
		AccordionItemLabel:   accordionItemLabel,
		PanelsState:          panelsState,
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
