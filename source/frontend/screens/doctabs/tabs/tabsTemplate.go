package tabs

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName     string
	LocalPanelNames []string
	ImportPrefix    string
	Funcs           _utils_.Funcs
}

const (
	Template = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

{{- range $panelName := .LocalPanelNames}}

// New{{ $panelName }}Tab constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the local {{ $panelName }} panel for content.
func New{{ $panelName }}Tab(screen *_misc_.Miscellaneous, spawned bool) (panel *_panels_.{{ $panelName }}Panel, err error) {
	tabItem := container.NewTabItem("{{ $panelName }}", widget.NewLabel("{{ $panelName }}"))
	tabItemContentConsumer := _types_.NewDocTabItemContentConsumer(screen.Layout.Tabbar(), tabItem, spawned)
	// The {{ $panelName }} panel.
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, tabItemContentConsumer, tabItem); err != nil {
		return
	}
	panelProducer := panel.Producer()
	tabItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerTabItemConsumer(panel, tabItem, tabItemContentConsumer)

	return
}
{{ end }}
`
)
