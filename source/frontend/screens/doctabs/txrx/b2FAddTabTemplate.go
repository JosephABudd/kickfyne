package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FAddTabTemplateData struct {
	PackageName       string
	ImportPrefix      string
	LocalPanelNames   []string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	B2FAddTabTemplate = `{{ $DOT := . -}}
package txrx

import (
	_tabs_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/tabs"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} is received when the back-end tells the front-end to spawn a new {{ .PackageName }} tab-bar tab.
// receive{{ .MessageStructName }} opens the new tab and updates the back-end.
func (messenger *Messenger) receive{{ .MessageStructName }}(msg *_message_.{{ .MessageStructName }}) {
	if msg.TabbarMessengerID != messenger.ID() {
		// This message is for another tabbar.
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

	switch msg.Tab {
{{- range $panelName := .LocalPanelNames }}
	case _message_.{{ $panelName }}Tab:
		// Spawn a new {{ $panelName }} tab.
		if msg.TabMessengerID, err = _tabs_.Open{{ $panelName }}Tab(messenger.screen, msg.TabIcon, msg.TabLabel); err != nil {
			return
		}
{{- end }}
	}
}
`
)
