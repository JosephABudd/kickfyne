package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BAddAccordionItemTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	F2BAddAccordionItemTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

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

// receive{{ .MessageStructName }} processes the {{ .MessageStructName }}.
// The {{ .MessageStructName }} message is sent from the front-end to the back-end.
// The front-end is informing the backend that a new accordionItem needs to be added to a accordion.
// The back-end must return the message to indicate success or failure.
func receive{{ .MessageStructName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {

	{{ call .Funcs.DeCap .MessageStructName }}Msg := msg.(*_message_.{{ .MessageStructName }})

	// The front-end added the accordionItem so do it here.
	_message_.AddAccordionMessengerIDAccordionItemMessengerID({{ call .Funcs.DeCap .MessageStructName }}Msg.AccordionMessengerID, {{ call .Funcs.DeCap .MessageStructName }}Msg.AccordionItemMessengerID)
}
`
)
