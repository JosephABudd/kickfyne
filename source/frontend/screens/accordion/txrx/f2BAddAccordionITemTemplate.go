package txrx

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

type F2BAddAccordionItemTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	F2BAddAccordionItemTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tell the back-end that the front-end has added a new accordion.
func (messenger *Messenger) Send{{ .MessageStructName }}(
	accordionItemMessengerID string,
) {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
		accordionItemMessengerID,
	)
	_txrxchans_.Send(msg)
}
`
)
