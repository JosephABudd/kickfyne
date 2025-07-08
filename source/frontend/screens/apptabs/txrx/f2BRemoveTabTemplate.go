package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BRemoveTabTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

func F2BRemoveTabFileName(packageName string) (fileName string) {
	messageNameFileName := _utils_.AppTabMessageNameFileName(packageName)
	fileName = messageNameFileName[_utils_.AppTabMessageF2BRemoveTab]
	return
}

const (
	F2BRemoveTabTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tells the back-end that the front-end has unspawned a tab.
func (messenger *Messenger) Send{{ .MessageStructName }}(
	tabMessengerID string,
) {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
		tabMessengerID,
	)
	// Send the message back with the PanelID.
	_txrxchans_.Send(msg)
}
`
)
