package misc

type MessengerTemplateData struct {
	PackageName      string
	ImportPrefix     string
	LocalPanelNames  []string
	DefaultPanelName string
}

const (
	MessengerFileName = "messenger.go"

	MessengerTemplate = `{{ $DOT := . -}}
package misc

import (
	"fmt"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_tabs_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/tabs"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	_message_ "{{ .ImportPrefix }}/shared/message"

	// Local panels.
{{- range $panelName := .LocalPanelNames }}
	_{{ $DOT.Funcs.LowerCase $panelName }}panel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ $panelName }}Panel"
{{- end }}
)

// type Messenger communicates with the backend listening for 1 message.
// The message is Spawn{{ .PackageName }}Tab.
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

func (messenger *Messenger) SendOpenedNew{{ .PackageName }}(remoteData _tabbar_model_.{{ .PackageName }}RemoteData) {
	msg := _message_.NewSpawn{{ .PackageName }}(
		messenger.ScreenPackage(),
		remoteData,
	)
	// Send the message back with the PanelID.
	_txrxchans_.Send(msg)
}

func (messenger *Messenger) StopReceiving() {
	_txrxchans_.UnSpawnReceiver(messenger)
}

// ScreenPackage returns this screen's package name.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) ScreenPackage() (name string) {
	name = messenger.screen.ScreenID
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
		_message_.Spawn{{ .PackageName }}TabID,
	); err != nil {
		return
	}

	return
}

// Receive receives dispatched messages.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) Receive(msg any) {
	switch msg := msg.(type) {
	case *_message_.Spawn{{ .PackageName }}Tab:
		messenger.receiveSpawn{{ .PackageName }}Tab(msg)
	case *_message_.Spawned{{ .PackageName }}Tab:
		messenger.receiveSpawned{{ .PackageName }}Tab(msg)
	}
}

// receiveSpawned{{ .PackageName }}Tab handles a received Spawn{{ .PackageName }}Tab message from the back-end.
func (messenger *Messenger) receiveSpawned{{ .PackageName }}Tab(msg *_message_.Spawned{{ .PackageName }}Tab) {
	// Here the back-end is replying to the front-end.
	// If there was an error then unspawn the spawned tab.
	if msg.Error {
		// There was a back-end error so cancel the spawn.
		// Unspawn the tab.
		messenger.screen.Layout.RemoveID(msg.TabMessengerID)
	}
}

// receiveSpawn{{ .PackageName }}Tab handles a received Spawn{{ .PackageName }}Tab message from the back-end.
// The back-end is asking the front-end to spawn a new tab.
func (messenger *Messenger) receiveSpawn{{ .PackageName }}Tab(msg *_message_.Spawn{{ .PackageName }}Tab) {
	var err error
	switch msg.Tab {
{{- range $panelName := .LocalPanelNames }}
	case _tabbar_model_.Tab{{ $panelName }}:
		// Spawn a new {{ $panelName }} Tab with it's {{ $panelName }} panel.
		var panel *_panels_.{{ $panelName }}Panel
		if panel, err = _tabs_.New{{ $panelName }}Tab(messenger.screen, true); err != nil {
			// Return the message with the error information.
			msg.Error = true
			msg.ErrorMessage = "{{ .PackageName }}.txrx.Messenger.receiveSpawn{{ .PackageName }}Tab: _tabs_.New{{ $panelName }}Tab error: " + err.Error()
			_txrxchans_.Send(msg)
			return
		}
		// Added a new Hello tab.
		// Set the state from the message.
		var state *_{{ $DOT.FuncsLowerCase $panelName }}panel_.State = panel.State()
		state.Set(
			// Tab settings.
			state.SetTabLabel(msg.TabLabel),
			state.SetTabIcon(msg.TabIcon),
			// Panel settings.
			state.SetHeading(msg.{{ $panelName }}PanelHeading),
			state.SetDescription(msg.{{ $panelName }}PanelDescription),
		)
		// Return the message with new tab information.
		msg.SpawnedMessengerID = panel.ID()
		_txrxchans_.Send(msg)
{{- end }}
	}
}

func (messenger *Messenger) unSpawn(msg *_message_.Spawned{{ .PackageName }}Tab) {
	// There was a back-end error so cancel the spawn.
	switch msg.Tab {
{{- range $panelName := .LocalPanelNames }}
	case _tabbar_model_.Tab{{ $panelName }}:
		// Remove this {{ $panelName }} tab.
		messenger.screen.Layout.RemoveID(msg.{{ .PackageName }}MessengerID)
{{- end }}
	}
}
`
)
