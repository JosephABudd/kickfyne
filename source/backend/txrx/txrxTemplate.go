package txrx

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

const (
	tXRXFileName = "txrx.go"
)

type tXRXTemplateData struct {
	ImportPrefix    string
	TXRXFolderNames []string
	Funcs           _utils_.Funcs
}

var tXRXTemplate = `{{ $DOT := . -}}
package txrx

import (
	"context"
	"log"

	_api_ "{{ .ImportPrefix }}/backend/txrx/api"
	_message_ "{{ .ImportPrefix }}/shared/message"
	_store_ "{{ .ImportPrefix }}/shared/store"
{{- if ne (len .TXRXFolderNames) 0 }}

{{ range $folderName := .TXRXFolderNames }}	_ "{{ $DOT.ImportPrefix }}/backend/txrx/{{ $folderName }}"
{{ end }}
{{- end}}
)

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
				var receivers []_api_.Receiver
				var found bool
				if receivers, found = _api_.MessageReceivers[id]; !found {
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
