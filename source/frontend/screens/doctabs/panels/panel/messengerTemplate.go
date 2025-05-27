package panel

type MessengerTemplateData = struct {
	PackageName  string
	PanelName    string
	ImportPrefix string
}

const (
	MessengerTemplate = `package {{ .PanelName }}Panel

import (
	"fmt"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	// _message_ "{{ .ImportPrefix }}/shared/message"
	// _record_ "{{ .ImportPrefix }}/shared/store/record"
)

const (
	packageName = "{{ .PackageName }}"
)

// type Messenger communicates with the backend using messages.
// It is an implementation of the _txrxchans_.Receiver interface.
type Messenger struct {
	screen    *_misc_.Miscellaneous
	state     *State
	showPanel func()
}

// NewMessenger constructs this Messenger.
func NewMessenger(screen *_misc_.Miscellaneous, state *State, showPanelFunc func()) (messenger *Messenger, err error) {
	messenger = &Messenger{
		screen:    screen,
		state:     state,
		showPanel: showPanelFunc,
	}
	err = messenger.startReceiving()
	return
}

func (messenger *Messenger) ID() (id string) {
	getters := messenger.state.Get().(Getters)
	id = getters.ID()
	return
}

// Send{{ .PackageName }}{{ .PanelName }}SpawnMessage(data any) requests initialization data from the back-end for a {{ .PanelName }} tab and it's panel.
// It is only used if this screen can be content for a spawned tab item or accordion item.
// If this func is to be used:
// * A message will need to be created and defined.
// * Param data should be publically defined above.
// * This func should be completed to send the message.
// * Another func should be completed to receive the message.
func (messenger *Messenger) Send{{ .PackageName }}{{ .PanelName }}SpawnMessage(data any) {
	_ = data
	// msg := _message_.New{{ .PackageName }}{{ .PanelName }}TabSpawnMessage(messenger.ID(), data)
	// _txrxchans_.Send(msg)
}

func (messenger *Messenger) StopReceiving() {
	_txrxchans_.UnSpawnReceiver(messenger)
}

// ScreenPackage returns this screen's package name.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) ScreenPackage() (name string) {
	name = packageName
	return
}

// startReceiving begins the receiving process.
// It adds this message receiver as a listener to messages from the back-end.
func (messenger *Messenger) startReceiving() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Messenger.startReceiving: %w", err)
		}
	}()

	/* KICKFYNE TODO:
	Add this listener for each message that is used by this package.
	The messages are in {{ .ImportPrefix }}/shared/message/ folder.
	Use a message's ID.

	Example:
	if err = _txrxchans_.AddReceiver(
		messenger,
		_message_.GetContactID, // GetContact message ID.
		_message_.FooID,        // Foo message ID.
		_message_.BarID         // Bar message ID.
	); err != nil {
		return
	}

	*/

	return
}

// Receive receives dispatched messages.
// It is part of the _txrxchans_.Receiver interface implementation.
func (messenger *Messenger) Receive(msg interface{}) {

	/* KICKFYNE TODO:
	Add a switch with cases for each message type and corresonding receiver.

	Example:
	switch msg := msg.(type) {
	case *_message_.GetContact:
		messenger.receiveGetContact(msg)
	case *_message_.Foo:
		messenger.receiveFoo(msg) // Not shown for brevity.
	case *_message_.Bar:
		messenger.receiveBar(msg) // Not shown for brevity.
	}

	*/
}

/* KICKFYNE TODO:
Add send funcs for each message sent.
Add receiver funcs for each message received.

Example:
// GetContact message.

// sendGetContact sends a GetContact message to the back-end.
func (messenger *Messenger) sendGetContact(r *_record_.GetContact) {
	msg := _message_.NewGetContact(messenger.ID(), r)
	_txrxchans_.Send(msg)
}

// receiveGetContact handles a received GetContact message from the back-end.
func (messenger *Messenger) receiveGetContact(msg *_message_.GetContact) {
	// Ignore this message if this messenger did not send it.
	if msg.MessengerID != messenger.ID() {
		return
	}
	// This messenger sent this message so process it.
	if msg.Error {
		// This message was sent by another screen.
		// That screen will deal with the error.
		return
	}
	// No error so process the message sent by the back-end.
	// It's a contact record for the EditPanel.
	contactRecord := msg.ContactRecord
	state := messenger.screen.Panelers.Edit.State()
	state.Set(
		state.SetName(contactRecord.Name),
		state.SetAddress(contactRecord.Address),
		state.SetCity(contactRecord.City),
		state.SetState(contactRecord.State),
		state.SetZip(contactRecord.Zip),
	)
	// Go back to the Edit panel.
	messenger.screen.Panelers.Edit.Show()
	// Tell the user there is a new heading.
	dialog.ShowInformation("Success", "The Edit panel has a contact record for you to edit.", messenger.screen.window)
}

*/

`
)
