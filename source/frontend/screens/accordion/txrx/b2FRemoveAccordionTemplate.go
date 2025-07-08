package txrx

type B2FRemoveAccordionTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
}

const (
	B2FRemoveAccordionTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} receives the {{ .MessageStructName }} message from the back-end.
// The message tells the front-end to remove this accordionItem-bar.
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

	// Remove (unbind) this accordion from it's consumer.
	accordionConsumer := messenger.screen.Layout.AccordionConsumer()
	accordionConsumer.UnBind()
}
`
)
