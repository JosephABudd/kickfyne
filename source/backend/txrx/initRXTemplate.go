package txrx

type initRXTemplateData struct {
	ImportPrefix string
}

var initRXTemplate = `package txrx

import (
	"context"
	"fmt"

	_api_ "{{ .ImportPrefix }}/backend/txrx/api"
	_message_ "{{ .ImportPrefix }}/shared/message"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

const initF = "receiveInit: %s"

func init() {
	_api_.AddReceiver(_message_.InitID, receiveInit)
}

func receiveInit(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {

	initMsg := msg.(*_message_.Init)
	var err, fatal error
	defer func() {
		switch {
		case err != nil:
			initMsg.Error = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, err.Error())
			_api_.Send(initMsg)
		case fatal != nil:
			initMsg.Fatal = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, fatal.Error())
			_api_.Send(initMsg)
		default:
			// No errors so don't send back the Init message.
		}
	}()

	/* KICKFYNE TODO:
	
	1. The GUI just got displayed.
	   If parts of the back-end need initialized then initialize them now.
	2. The front-end screens are ready to receive messages for initialization.
	   Send those messages now.
	*/

}

`
