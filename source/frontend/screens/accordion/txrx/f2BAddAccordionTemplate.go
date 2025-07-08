package txrx

type F2BAddAccordionTemplateData struct {
	PackageName       string
	AllPanelNames     []string
	ImportPrefix      string
	MessageStructName string
}

const (
	F2BAddAccordionTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tell the back-end that the front-end has added a new accordion.
func (messenger *Messenger) Send{{ .MessageStructName }}() {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
	)
	_txrxchans_.Send(msg)
}
`
)
