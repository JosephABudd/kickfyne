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

// SendSpawnedMessage(data interface) requests initialization data from the back-end.
// It is only used if this screen can be content for a spawned tab item or accordion item.
// It is an implementation of the _types_.MessageSpawner interface.
// If this func is to be used:
// * A message will need to be created and defined.
// * Param data should be publically defined above.
// * This func should be completed to send the message.
// * Another func should be completed to receive the message.
func (messenger *Messenger) SendSpawnedMessage(data any) {}

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
		_message_.GetContactID,
		_message_.FooID,
		_message_.BarID
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
		messenger.receiveFoo(msg)
	case *_message_.Bar:
		messenger.receiveBar(msg)
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
	msg := _message_.GetContact(packageName, r)
	_txrxchans_.Send(msg)
}

// receiveGetContact handles a received GetContact message from the back-end.
func (messenger *Messenger) receiveGetContact(msg *_message_.GetContact) {
	if msg.Error {
		if msg.ScreenPackage != packageName {
			// This message was sent by another screen.
			// That screen will deal with the error.
			return
		}
		// This screen sent the message so receive the error here.
		dialog.ShowInformation("Error", msg.ErrorMessage, messenger.screen.window)
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
