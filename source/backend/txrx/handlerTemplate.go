package txrx

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

type handlerTemplateData struct {
	ImportPrefix string
	MessageName  string
	Funcs        _utils_.Funcs
}

var handlerTemplate = `{{ $dCMessageName := call .Funcs.DeCap .MessageName }}package txrx

import (
	"context"
	"fmt"

	_message_ "{{ .ImportPrefix }}/shared/message"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

const {{ $dCMessageName }}F = "receive{{ .MessageName }}: %s"

func init() {
	addReceiver(_message_.{{ .MessageName }}ID, receive{{ .MessageName }})
}

func receive{{ .MessageName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {

	{{ $dCMessageName }}Msg := msg.(*_message_.{{ .MessageName }})
	var err, fatal error
	defer func() {
		switch {
		case fatal != nil:
			{{ $dCMessageName }}Msg.Fatal = true
			{{ $dCMessageName }}Msg.ErrorMessage = fmt.Sprintf({{ $dCMessageName }}F, fatal.Error())
			Send({{ $dCMessageName }}Msg)
		case err != nil:
			{{ $dCMessageName }}Msg.Error = true
			{{ $dCMessageName }}Msg.ErrorMessage = fmt.Sprintf({{ $dCMessageName }}F, err.Error())
			Send({{ $dCMessageName }}Msg)
		default:
			// No errors so return the {{ $dCMessageName }}Msg.
			Send({{ $dCMessageName }}Msg)
		}
	}()

	/* KICKFYNE TODO:
	Do something with this message.
	Use fatal for unrecoverable errors.
	Use err for user error messages.
	*/
}

`
