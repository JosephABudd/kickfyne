package backend

const (
	fileName = "backend.go"
)

type templateData struct {
	ImportPrefix string
}

var template = `package backend

import (
	"context"

	_txrx_ "{{ .ImportPrefix }}/backend/txrx"
	_store_ "{{ .ImportPrefix }}/shared/store"
)

// Start starts the backend.
func Start(ctx context.Context, ctxCancel context.CancelFunc) {
	stores := _store_.New()
	// Messages.
	_txrx_.StartReceiver(ctx, ctxCancel, stores)
}

`
