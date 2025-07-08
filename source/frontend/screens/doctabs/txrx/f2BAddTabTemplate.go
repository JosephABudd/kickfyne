package txrx

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

type F2BAddTabTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	F2BAddTabTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tell the back-end that the front-end has added a new tabbar.
func (messenger *Messenger) Send{{ .MessageStructName }}(
	tabMessengerID string,
) {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
		tabMessengerID,
	)
	_txrxchans_.Send(msg)
}
`
)
