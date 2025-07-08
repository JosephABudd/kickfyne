package txrx

type B2FRemoveAccordionItemTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
}

const (
	B2FRemoveAccordionItemTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} receives the {{ .MessageStructName }} message from the back-end.
// The message tells the front-end to remove this accordionItem.
func (messenger *Messenger) receive{{ .MessageStructName }}(msg *_message_.{{ .MessageStructName }}) {
	if msg.AccordionMessengerID != messenger.ID() {
		// This message is for another accordion.
		return
	}

	var err error
	defer func() {
		if err != nil {
			msg.Error = true
		}
		// Send the message back to the back-end.
		_txrxchans_.Send(msg)
	}()

	// Remove the accordionItem from the accordion.
	messenger.screen.Layout.RemoveID(msg.AccordionItemMessengerID)
}
`
)
