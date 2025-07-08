package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BRemoveTabbarTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	F2BRemoveTabbarTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tells the back-end that the front-end has closed a tabbar.
func (messenger *Messenger) Send{{ .MessageStructName }}() {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
	)
	// Send the message back with the PanelID.
	_txrxchans_.Send(msg)
}
`
)
