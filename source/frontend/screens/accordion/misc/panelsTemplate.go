package misc

import (
	"github.com/JosephABudd/kickfyne/source/utils"
)

type PanelsTemplateData struct {
	PackageName      string
	PanelNames       []string // All panels are local. [] is default.
	DefaultPanelName string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            utils.Funcs
}

const (
	PanelsFileName = "panels.go"

	PanelsTemplate = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_panels_ "{{ .ImportPrefix }}/frontend/gui/screens/tabbar/{{ .PackageName }}/panels"

)

var screen_panels: *Panels = nil

// Panels is this screen's panels.
// This screen has {{ len .LocalPanelNames }} panels.
// The default panel is {{ .DefaultPanelName }}.
type Panels struct {
	// Construct each panel building panel content.
{{- range $i, $panelName := .LocalPanelNames }}
 {{- if eq $i 0 }}
	default *_panels_.{{ $panelName }}Panel
 {{- end }}
	{{ $panelName }} *_panels_.{{ $panelName }}Panel
{{- end }}
	default *_panels_.{{ .DefaultPanelName }}Panel
}

// NewPanels constructs a Panels.
func NewPanels(screenComponents *ScreenComponents) (panels *Panels, err error) {
	defer func() {
		if err == nil {
			panels = screen_panels
		}
	}

	if screen_panels != nil {
		return
	}

	// Construct panels.
	screen_panels = &Panels{}
{{- range $i, $panelName := .LocalPanelNames }}
	if screen_panels.{{ $panelName }}, err = _panels_.New{{ $panelName }}Panel(screenComponents); err != nil {
		return
	}
{{- end }}
	screen_panels.default = panels.{{ .DefaultPanelName }}
	return
}
`
)
