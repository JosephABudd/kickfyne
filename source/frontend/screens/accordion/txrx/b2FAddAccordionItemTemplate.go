package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FAddAccordionItemTemplateData struct {
	PackageName       string
	ImportPrefix      string
	LocalPanelNames   []string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	B2FAddAccordionItemTemplate = `{{ $DOT := . -}}
package txrx

import (
	_accordionItems_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/accordionitems"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} is received when the back-end tells the front-end to spawn a new {{ .PackageName }} accordion accordionItem.
// receive{{ .MessageStructName }} opens the new accordionItem and updates the back-end.
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

	switch msg.AccordionItem {
{{- range $panelName := .LocalPanelNames }}
	case _message_.{{ $panelName }}AccordionItem:
		// Spawn a new {{ $panelName }} accordionItem.
		if msg.AccordionItemMessengerID, err = _accordionItems_.Open{{ $panelName }}AccordionItem(messenger.screen, msg.AccordionItemLabel); err != nil {
			return
		}
{{- end }}
	}
}
`
)
