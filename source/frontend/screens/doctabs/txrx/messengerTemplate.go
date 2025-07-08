package txrx

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type MessengerTemplateData struct {
	PackageName                  string
	ImportPrefix                 string
	LocalPanelNames              []string
	DocTabMessageB2FAddTabbar    string
	DocTabMessageB2FAddTab       string
	DocTabMessageB2FRemoveTabbar string
	DocTabMessageB2FRemoveTab    string
	DocTabMessageF2BAddTabbar    string
	DocTabMessageF2BAddTab       string
	DocTabMessageF2BRemoveTabbar string
	DocTabMessageF2BRemoveTab    string
	Funcs                        _utils_.Funcs
}

const (
	MessengerFileName = "messenger.go"

	MessengerTemplate = `{{ $DOT := . -}}
package txrx

import (
	"fmt"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"

	// Local panels.
{{- range $panelName := .LocalPanelNames }}
	// _{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
{{- end }}
)

// type Messenger communicates with the backend using 8 messages.
// Messenger implements the _txrxchans_.Receiver interface.
type Messenger struct {
	screen *_misc_.Miscellaneous
}

// NewMessenger constructs this Messenger.
func NewMessenger(screen *_misc_.Miscellaneous) (messenger *Messenger, err error) {
	messenger = &Messenger{
		screen: screen,
	}
	err = messenger.startReceiving()
	return
}

func (messenger *Messenger) StopReceiving() {
	_txrxchans_.UnSpawnReceiver(messenger)
}

// ID returns this messenger's id.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) ID() (id string) {
	id = messenger.screen.ScreenID
	return
}

// startReceiving begins the receiving process.
// It adds this message receiver as a listener to the Spawn{{ .PackageName }}Tab message from the back-end.
func (messenger *Messenger) startReceiving() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Messenger.startReceiving: %w", err)
		}
	}()

	// Listen for the Spawn{{ .PackageName }}Tab message.
	if err = _txrxchans_.AddReceiver(
		messenger,
		_message_.B2FAddTabbar{{ .PackageName }}ID,
		_message_.B2FAddTab{{ .PackageName }}ID,
		_message_.B2FRemoveTabbar{{ .PackageName }}ID,
		_message_.B2FRemoveTab{{ .PackageName }}ID,
	); err != nil {
		return
	}

	return
}

// Receive receives dispatched messages.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) Receive(msg any) {
	switch msg := msg.(type) {
	case *_message_.B2FAddTabbar{{ .PackageName }}:
		messenger.receive{{ .DocTabMessageB2FAddTabbar }}(msg)
	case *_message_.B2FAddTab{{ .PackageName }}:
		messenger.receive{{ .DocTabMessageB2FAddTab }}(msg)
	case *_message_.B2FRemoveTabbar{{ .PackageName }}:
		messenger.receive{{ .DocTabMessageB2FRemoveTabbar }}(msg)
	case *_message_.B2FRemoveTab{{ .PackageName }}:
		messenger.receive{{ .DocTabMessageB2FRemoveTab }}(msg)
	}
}

// LoadStartupData(data any) requests initialization data from the back-end.
// It is an implementation of the _types_.StartupMessenger any.
// If this func is to be used:
// * A message will need to be created and defined.
// * Param data should be publically defined above.
// * This func should be completed to send the message.
// * Another func should be completed to receive the message.
func (messenger *Messenger) LoadStartupData(data any) {}
`
)
