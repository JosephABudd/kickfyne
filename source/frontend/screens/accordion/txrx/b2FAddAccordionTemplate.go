package txrx

type B2FAddAccordionTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
}

const (
	B2FAddAccordionTemplate = `{{ $DOT := . -}}
package txrx

import (
	"fmt"

	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} is not used and returns an error to the back-end.
// receive{{ .MessageStructName }} receives the {{ .MessageStructName }} message from the back-end.
// The message tells the front-end to spawn a new {{ .PackageName }} accordion accordionItem.
// receive{{ .MessageStructName }} would if used, open a new accordion and update the back-end.
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

	err = fmt.Errorf("The {{ .MessageStructName }} message is not currently used.")
}
`
)
