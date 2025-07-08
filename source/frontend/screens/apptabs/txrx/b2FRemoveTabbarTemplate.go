package txrx

type B2FRemoveTabbarTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
}

const (
	B2FRemoveTabbarTemplate = `{{ $DOT := . -}}
package txrx

import (
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
)

// receive{{ .MessageStructName }} receives the {{ .MessageStructName }} message from the back-end.
// The message tells the front-end to remove this tab-bar.
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

	// Remove (unbind) this tabbar from it's consumer.
	tabBarConsumer := messenger.screen.Layout.TabbarConsumer()
	tabBarConsumer.UnBind()
}
`
)
