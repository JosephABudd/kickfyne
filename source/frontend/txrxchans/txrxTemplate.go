package txrxchans

type txrxTemplateData struct {
	ImportPrefix string
}

const (
	txrxFileName = "txrxchans.go"

	txrxTemplate = `package txrxchans

import (
	"context"
	"fmt"
	"log"

	_message_ "{{ .ImportPrefix }}/shared/message"
)

var messageReceivers = make(map[string][]Receiver, 20)
var Dispatcher = make(chan interface{}, 1024)

type Receiver interface {
	Receive(msg interface{})
	ID() string
}

// AddReceiver adds the receiver to the messages indicated by msgIDs.
func AddReceiver(receiver Receiver, msgIDs ...string) (err error) {
	for _, msgID := range msgIDs {
		if err = addReceiver(receiver, msgID); err != nil {
			return
		}
	}
	return
}

// RemoveReceiver removes the receiver from the message indicated by msgID.
func RemoveReceiver(msgID string, receiver Receiver) {
	if _, found := messageReceivers[msgID]; !found {
		return
	}
	removeReceiver(receiver, msgID)
}

// UnSpawnReceiver removes the receiver from all messages.
func UnSpawnReceiver(receiver Receiver) {
	for msgID := range messageReceivers {
		removeReceiver(receiver, msgID)
	}
}

func removeReceiver(receiver Receiver, msgID string) {
	receivers := messageReceivers[msgID]
	for i, l := range receivers {
		if l == receiver {
			messageReceivers[msgID] = receivers[0:i]
			if i++; i<len(receivers) {
				messageReceivers[msgID] = append(messageReceivers[msgID], receivers[i:]...)
			}
			return
		}
	}
}

func addReceiver(receiver Receiver, msgID string) (err error) {
	var receivers []Receiver
	var found bool
	if receivers, found = messageReceivers[msgID]; !found {
		receivers = make([]Receiver, 0, 20)
	}
	// Don't allow for duplicates.
	for _, l := range receivers {
		if l == receiver {
			err = fmt.Errorf("The screen package %q is already receiving to %q.", receiver.ID(), msgID)
			return
		}
	}
	receivers = append(receivers, receiver)
	messageReceivers[msgID] = receivers
	return
}

// StartReceiver starts receiving messages from the back-end and dispathes them to the front-end.
// If a fatal message is received it logs it and closes the app.
func StartReceiver(ctx context.Context, ctxCancel context.CancelFunc) {
	go func(ctx context.Context, ctxCancel context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Frontend Receiver DONE")
				return
			case msg := <-_message_.BackEndToFrontEnd:
				if isFatal, errorMessage, screenPackage := msg.FatalError(); isFatal {
					log.Printf("frontend: txrx.Receive: Fatal from back-end: %q; ScreenPackge: %q.", errorMessage, screenPackage)
					ctxCancel()
					return
				}
				id := msg.ID()
				var receivers []Receiver
				var found bool
				if receivers, found = messageReceivers[id]; !found {
					// No receivers for this _message_.
					continue
				}
				// Dispatch the _message_.
				realMSG := msg.AsInterface()
				for _, l := range receivers {
					go l.Receive(realMSG)
				}
			}
		}
	}(ctx, ctxCancel)
}

// Send sends a message to the back-end.
func Send(msg _message_.MSGer) {
	_message_.FrontEndToBackEnd <- msg
}
`
)
