package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BRemoveAccordionItemTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

func F2BRemoveAccordionItemFileName(packageName string) (fileName string) {
	messageNameFileName := _utils_.AccordionMessageNameFileName(packageName)
	fileName = messageNameFileName[_utils_.AccordionMessageF2BRemoveAccordionItem]
	return
}

const (
	F2BRemoveAccordionItemTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// Send{{ .MessageStructName }} tells the back-end that the front-end has unspawned a accordionItem.
func (messenger *Messenger) Send{{ .MessageStructName }}(
	accordionItemMessengerID string,
) {
	msg := _message_.New{{ .MessageStructName }}(
		messenger.ID(),
		accordionItemMessengerID,
	)
	// Send the message back with the PanelID.
	_txrxchans_.Send(msg)
}
`
)
