package misc

type MessengerTemplateData struct {
	PackageName  string
	ImportPrefix string
}

const (
	MessengerFileName = "messenger.go"

	MessengerTemplate = `package content

import (
	"fmt"

	"_message_ {{ .ImportPrefix }}/deps/message"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
	)

const (
	packageName = "{{ .PackageName }}"
)

// type Messenger communicates with the backend using messages.
// It is an implementation of the _txrxchans_.Receiver interface.
type Messenger struct{
	screen *ScreenComponents
}

// newMessenger constructs this message handler.
func newMessenger(screen *ScreenComponents) (messenger *Messenger) {
	messenger = &Messenger{
		screen: screen,
	}
	messenger.startReceiving()
	return
}

// StopReceiving stops the messenger from receiving messages. 
func (m *Messenger) StopReceiving() {
	_txrxchans_.UnSpawnReceiver(m)
}

// SendSpawnedMessage(data interface) requests initialization data from the back-end.
// It is only used if this screen can be content for a spawned tab item or accordion item.
// It is an implementation of the _types_.MessageSpawner interface.
// If this func is to be used:
// * A message will need to be created and defined.
// * Param data should be publically defined above.
// * This func should be completed to send the message.
// * Another func should be completed to receive the message.
func (m *Messenger) SendSpawnedMessage(data interface) {}

// ScreenPackage returns this screen's package name.
// It is part of the _txrxchans_.Receiver interface implementation.
func (m *Messenger) ScreenPackage() (name string) {
	name = packageName
	return
}

// ScreenPackage returns this screen's package name.
// It is part of the _txrxchans_.Receiver interface implementation.
func (m *Messenger) ScreenPackage() (name string) {
	name = packageName
	return
}

// startReceiving begins the receiving process.
// It adds this message receiver as a listener to messages from the back-end.
func (m *Messenger) startReceiving() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Messenger.listen: %w", err)
		}
	}()

	/* KICKFYNE TODO:
	Add this listener for each message that is used by this package.
	The messages are in {{ .ImportPrefix }}/shared/message/ folder.
	Use a message's ID.

	Example:
	if err = _txrxchans_.AddReceiver(
		m,
		_message_.GetSomethingID,
		_message_.DoSomethingID,
	); err != nil {
		return
	}

	*/

	return
}

// Receive receives dispatched messages.
// It is part of the _txrxchans_.Receiver interface implementation.
func (m *Messenger) Receive(msg interface{}) {

	/* KICKFYNE TODO:
	Add a switch with cases for each message type and corresonding receiver.

	Example:
	switch msg := msg.(type) {
	case *_message_.GetSomething:
		m.receiveGetSomething(msg)
	case *_message_.DoSomething:
		m.receiveDoSomething(msg)
	}

	*/
}

/* KICKFYNE TODO:
Add send funcs for each message sent.
Add receiver funcs for each message received.

Example:
// GetContact message.

// sendGetContact sends a GetContact message to the back-end.
func (m *Messenger) sendGetContact(r *record.GetContact) {
	msg := _message_.HelloTX(packageName, r)
	_txrxchans_.Send(msg)
}

// receiveGetContact handles a received GetContact message from the back-end.
func (m *Messenger) receiveGetContact(msg *_message_.GetContact) {
	if msg.Error {
		if msg.ScreenPackage != packageName {
			// This message was sent by another screen.
			// That screen will deal with the error.
			return
		}
		// This screen sent the message so receive the error here.
		dialog.ShowInformation("Error", msg.ErrorMessage, m.screen.window)
		return
	}
	m.screen.panels.editPanel.FillForm(msg.Contact)
	m.screen.panels.editPanel.show()
}

// DoSomething message.

// sendDoSomething sends an DoSomething message to the back-end.
func (m *Messenger) sendDoSomething(r *record.ContactEdit) {
	msg := _message_.NewDoSomething(packageName, r)
	_txrxchans_.Send(msg)

}

// receiveDoSomething handles a received DoSomething message from the back-end.
func (m *Messenger) receiveDoSomething(msg *_message_.DoSomething) {
	if msg.ScreenPackage != packageName {
		// This particular message is ignored here if not sent by this screen.
		return
	}
	// This screen sent the message so it will receive it.
	if msg.Error {
		dialog.ShowInformation("Error", msg.ErrorMessage, window)
		return
	}
	dialog.ShowInformation("Success", "Did something.", m.screen.window)
	// Go back to the select panel.
	m.screen.panels.selectPanel.show()
}

*/

`
)
