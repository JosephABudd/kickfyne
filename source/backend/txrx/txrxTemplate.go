package txrx

const (
	tXRXFileName = "txrx.go"
)

type tXRXTemplateData struct {
	ImportPrefix string
}

var tXRXTemplate = `package txrx

import (
	"context"
	"fmt"
	"log"

	_message_ "{{ .ImportPrefix }}/shared/message"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

type Receiver func(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{})

var (
	messageReceivers = make(map[uint64][]Receiver, 20)
)

// Send sends a message to the front-end.
func Send(msg _message_.MSGer) {
	_message_.BackEndToFrontEnd <- msg
}

// addReceiver adds the number of receivers.
func addReceiver(msgID uint64, receiver Receiver) (err error) {
	if !_message_.IsValidID(msgID) {
		err = fmt.Errorf("_store_.AddReceiver: message id not found")
		return
	}
	var receivers []Receiver
	var found bool
	if receivers, found = messageReceivers[msgID]; !found {
		receivers = make([]Receiver, 0, 5)
	}
	receivers = append(receivers, receiver)
	messageReceivers[msgID] = receivers
	return
}

// StartReceiver starts receiving messages from the front-end and dispatches them to the back-end.
func StartReceiver(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores) {
	go func(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Backend Receiver DONE")
				return
			case msg := <-_message_.FrontEndToBackEnd:
				id := msg.ID()
				name := msg.Name()
				var receivers []Receiver
				var found bool
				if receivers, found = messageReceivers[id]; !found {
					log.Printf("backend receivers not found for *_store_.%s", name)
					continue
				}
				realMSG := msg.AsInterface()
				for _, f := range receivers {
					go f(ctx, ctxCancel, stores, realMSG)
				}
			}
		}
	}(ctx, ctxCancel, stores)
}

`
