package doctabs

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BAddTabbarTemplateData struct {
	PackageName                string
	ImportPrefix               string
	MessageStructName          string
	B2FAddTabMessageStructName string
	Funcs                      _utils_.Funcs
}

const (
	F2BAddTabbarTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

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
// The message is telling the back-end that a new tabbar needs to be added.
// The back-end must return the message to indicate success or failure.
func receive{{ .MessageStructName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {

	{{ call .Funcs.DeCap .MessageStructName }}Msg := msg.(*_message_.{{ .MessageStructName }})

	// The front-end added the tabbar so do it here.
	_message_.AddTabbarMessengerID({{ call .Funcs.DeCap .MessageStructName }}Msg.TabbarMessengerID)
}
`
)
