package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FRemoveAccordionTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	B2FRemoveAccordionTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

import (
	"context"

	_api_ "{{ .ImportPrefix }}/backend/txrx/api"
	_message_ "{{ .ImportPrefix }}/shared/message/{{ .PackageName }}"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

const {{ call .Funcs.DeCap .MessageStructName }}F = "receive{{ .MessageStructName }}: %s"

func init() {
	_api_.AddReceiver(_message_.{{ .MessageStructName }}ID, receive{{ .MessageStructName }})
}

// Send{{ .MessageStructName }} sends this message from the back-end to the front-end.
// This message is telling the front-end that a accordion needs to be removed.
// The front-end must return the message to indicate success or failure.
func Send{{ .MessageStructName }} (
	accordionMessengerID string,
) {
	var {{ call .Funcs.DeCap .MessageStructName }} *_message_.{{ .MessageStructName }} = _message_.New{{ .MessageStructName }}(
		accordionMessengerID,
	)
	_api_.Send({{ call .Funcs.DeCap .MessageStructName }})
}

// receive{{ .MessageStructName }} informs the back-end that the front-end has removed the accordion.
func receive{{ .MessageStructName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {
	{{ call .Funcs.DeCap .MessageStructName }}Msg := msg.(*_message_.{{ .MessageStructName }})
	if {{ call .Funcs.DeCap .MessageStructName }}Msg.Error {
		// The front-end was unable to remove the accordion.
		return
	}
	// The front-end removed the accordion so do it here.
	_message_.RemoveAccordionMessengerID({{ call .Funcs.DeCap .MessageStructName }}Msg.AccordionMessengerID)
}
`
)
