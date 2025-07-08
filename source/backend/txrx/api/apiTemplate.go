package txrx

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

type APITemplateData struct {
	ImportPrefix    string
	TXRXFolderNames []string
	Funcs           _utils_.Funcs
}

var APITemplate = `{{ $DOT := . -}}
package api

import (
	"context"

	_message_ "{{ .ImportPrefix }}/shared/message"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

type Receiver func(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{})

var (
	MessageReceivers = make(map[string][]Receiver, 20)
)

// Send sends a message to the front-end.
func Send(msg _message_.MSGer) {
	_message_.BackEndToFrontEnd <- msg
}

// AddReceiver adds the number of receivers.
func AddReceiver(msgID string, receiver Receiver) (err error) {
	var receivers []Receiver
	var found bool
	if receivers, found = MessageReceivers[msgID]; !found {
		receivers = make([]Receiver, 0, 5)
	}
	receivers = append(receivers, receiver)
	MessageReceivers[msgID] = receivers
	return
}

`
