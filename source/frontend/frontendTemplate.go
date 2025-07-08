package frontend

const (
	frontendFileName = "frontend.go"
)

type frontendTemplateData struct {
	ImportPrefix string
	ScreenNames  []string
}

var frontendTemplate = `{{ $DOT := . -}}
package frontend

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	_mainmenu_ "{{ .ImportPrefix }}/frontend/mainmenu"
	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
	_txrxchans_ "{{ .ImportPrefix }}/frontend/txrxchans"

{{ range $screenName := .ScreenNames }}
	_{{ $screenName }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}"
{{- end }}
)

func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Set the screen map.
{{ range $screenName := .ScreenNames }}
	_screenmap_.Map["{{ $screenName }}"] = &_screenmap_.API{
		NewWindowContentConsumer:         _{{ $screenName }}_.NewWindowContentConsumer,
		NewAppTabsTabItemContentConsumer: _{{ $screenName }}_.NewAppTabsTabItemContentConsumer,
		NewDocTabsTabItemContentConsumer: _{{ $screenName }}_.NewDocTabsTabItemContentConsumer,
		NewAccordionItemContentConsumer:  _{{ $screenName }}_.NewAccordionItemContentConsumer,
	}
{{- end }}

	// Initialize main menu.
	// The developer must ensure that all panel groups should get initialized from main menu.
	_mainmenu_.Init(ctx, ctxCancelFunc, app, window)

	// Start communications with the back-end.
	// The receiver will run as a concurrent process.
	_txrxchans_.StartReceiver(ctx, ctxCancelFunc)

	return
}
`
