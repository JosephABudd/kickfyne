package doctabs

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type F2BRemoveTabTemplateData struct {
	PackageName       string
	ImportPrefix      string
	MessageStructName string
	Funcs             _utils_.Funcs
}

const (
	F2BRemoveTabTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

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
// The front-end is informing the backend that a tabbar needs to be removed.
// The back-end must return the message to indicate success or failure.
func receive{{ .MessageStructName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *_store_.Stores, msg interface{}) {

	{{ call .Funcs.DeCap .MessageStructName }}Msg := msg.(*_message_.{{ .MessageStructName }})

	// The front-end removed the tab so do it here.
	_message_.RemoveTabMessengerID({{ call .Funcs.DeCap .MessageStructName }}Msg.TabbarMessengerID, {{ call .Funcs.DeCap .MessageStructName }}Msg.TabMessengerID)
}
`
)
