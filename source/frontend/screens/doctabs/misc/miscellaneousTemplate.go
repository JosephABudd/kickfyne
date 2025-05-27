package misc

type MiscellaneousTemplateData struct {
	PackageName  string
	ImportPrefix string
}

const (
	MiscellaneousFileName = "miscellaneous.go"

	MiscellaneousTemplate = `{{ $DOT := . -}}
package misc

import (
	"context"

	"fyne.io/fyne/v2"

	_layout_ "example.com/okp/frontend/screens/{{ .PackageName }}/layout"
	_panelers_ "example.com/okp/frontend/screens/{{ .PackageName }}/panelers"
	_producer_ "example.com/okp/frontend/screens/{{ .PackageName }}/producer"
)

// Miscellaneous is a variety of components for this layout, panels and messenger.
type Miscellaneous struct {
	CTX       context.Context
	CTXCancel context.CancelFunc
	APP       fyne.App
	Window    fyne.Window

	Layout   *_layout_.Layout
	Panelers *_panelers_.Panelers
	Producer *_producer_.ContentProducer
}

// NewMiscellaneous constrtucts a Miscellaneous.
// Its parts of the screen that can be shared in one struct.
// So it does not include the messenger.
func NewMiscellaneous(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, layout *_layout_.Layout) (components *Miscellaneous, err error) {
	components = &Miscellaneous{
		CTX:       ctx,
		CTXCancel: ctxCancel,
		APP:       app,
		Window:    w,
		Panelers:  &_panelers_.Panelers{},
		Layout:    layout,
	}
	return
}
`
)
