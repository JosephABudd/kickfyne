package screenmap

type screenMapTemplateData struct {
	ImportPrefix string
	ScreenNames  []string
}

const (
	screenMapFileName = "screenmap.go"
	screenMapTemplate = `{{ $DOT := . -}}
package screenmap

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

{{ range $screenName := .ScreenNames }}
	_{{ $screenName }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}"
{{- end }}

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type API struct {
	NewWindowContentConsumer        func(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (consumer *_types_.WindowContentConsumer, err error)
	NewTabItemContentConsumer       func(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer) (tabItemConsumer *_types_.TabItemContentConsumer, err error)
	NewAccordionItemContentConsumer func(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer) (consumer *_types_.AccordionItemContentConsumer, err error)
}

var Map = map[string]*API{
{{ range $screenName := .ScreenNames }}
	"{{ $screenName }}": {
		NewWindowContentConsumer:        _{{ $screenName }}_.NewWindowContentConsumer,
		NewTabItemContentConsumer:       _{{ $screenName }}_.NewTabItemContentConsumer,
		NewAccordionItemContentConsumer: _{{ $screenName }}_.NewAccordionItemContentConsumer,
	},
{{- end }}
}
`
)
