package frontend

const (
	frontendFileName = "frontend.go"
)

type frontendTemplateData struct {
	ImportPrefix string
}

var frontendTemplate = `package frontend

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	_mainmenu_ "{{ .ImportPrefix }}/frontend/mainmenu"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"
)

func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Initialize main menu.
	// The developer must ensure that all panel groups should get initialized from main menu.
	if err = _mainmenu_.Init(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}

	// Start communications with the back-end.
	// The receiver will run as a concurrent process.
	_txrxchans_.StartReceiver(ctx, ctxCancelFunc)
	return
}

`
