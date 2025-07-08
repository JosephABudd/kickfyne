package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type B2FAddAccordionTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	B2FAddAccordionTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

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
// This message is telling the front-end that a new accordion needs to be added.
// The front-end must return the message to indicate success or failure.
func Send{{ .MessageStructName }} (
	accordionMessengerID string,
) {
	var {{ call .Funcs.DeCap .MessageStructName }} *_message_.{{ .MessageStructName }} = _message_.New{{ .MessageStructName }}(
		accordionMessengerID,
	)
	_api_.Send({{ call .Funcs.DeCap .MessageStructName }})
}

// receive{{ .MessageStructName }} receives this message from the front-end.
// The front-end is responding to this message which was originally sent by the back-end.
// The front-end is informing the back-end about the success or failure adding a new {{ .PackageName }} accordionItem.
func receive{{ .MessageStructName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {
	{{ call .Funcs.DeCap .MessageStructName }}Msg := msg.(*_message_.{{ .MessageStructName }})
	if {{ call .Funcs.DeCap .MessageStructName }}Msg.Error {
		// The accordionItem was not spawned.
		return
	}
	// The front-end added the accordion so do it here.
	_message_.AddAccordionMessengerID({{ call .Funcs.DeCap .MessageStructName }}Msg.AccordionMessengerID)
}
`
)
