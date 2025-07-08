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

// type Messenger communicates with the backend using messages.
// It is an implementation of the _txrxchans_.Receiver interface.
type Messenger struct {
	screen    *_misc_.Miscellaneous
	state     *State
}

// NewMessenger constructs this Messenger.
func NewMessenger(screen *_misc_.Miscellaneous, state *State) (messenger *Messenger, err error) {
	messenger = &Messenger{
		screen:    screen,
		state:     state,
	}
	err = messenger.startReceiving()
	return
}

// ID returns this messenger's id. Same as this messenger's panel's id.
func (messenger *Messenger) ID() (id string) {
	getters := messenger.state.Get().(Getters)
	id = getters.ID()
	return
}

// StopReceiving ends the receiving process.
func (messenger *Messenger) StopReceiving() {
	_txrxchans_.UnSpawnReceiver(messenger)
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
		_message_.GetHeadingDescriptionID, // GetHeadingDescription message ID.
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
	case *_message_.GetHeadingDescription:
		messenger.receiveGetHeadingDescription(msg) // See the example.
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
// GetHeadingDescription message.

// sendGetHeadingDescription sends a GetHeadingDescription message to the back-end.
func (messenger *Messenger) sendGetHeadingDescription(heading, description string) {
	msg := _message_.NewGetHeadingDescription(messenger.ID(), heading, description)
	_txrxchans_.Send(msg)
}

// receiveGetHeadingDescription handles a received GetHeadingDescription message from the back-end.
func (messenger *Messenger) receiveGetHeadingDescription(msg *_message_.GetHeadingDescription) {
	// Ignore this message if this messenger did not send it.
	if msg.MessengerID != messenger.ID() {
		// This message was sent by another screen so ignore it.
		return
	}
	// This messenger sent this message so process it.
	if msg.Error {
		dialog.ShowInformation("Error", msg.ErrorMessage, messenger.screen.window)
		return
	}
	// No error so process the message sent by the back-end.
	// It's a contact record for the EditPanel.
	state := messenger.screen.Panelers.Edit.State()
	state.Set(
		state.SetHeading(msg.Heading),
		state.SetDescription(msg.Descriptioin),
	)
	// Go back to the {{ .PanelName }} panel.
	messenger.screen.Panelers.{{ .PanelName }}.Show()
	// Tell the user there is a new heading.
	dialog.ShowInformation("Success", "The {{ .PanelName }} panel has a heading and description.", messenger.screen.window)
}

*/

`
)
