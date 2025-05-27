package screens

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type ContentTemplateData struct {
	PackageName     string
	PanelName       string
	LocalPanelNames []string
	ImportPrefix    string
	Funcs           _utils_.Funcs
}

const (
	ContentFileName = _utils_.ContentFileName

	ContentTemplate = `{{ $DOT := . -}}
package {{ .PanelName }}Panel

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	content fyne.CanvasObject
	screen  *_misc_.Miscellaneous

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
}

// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(screen *_misc_.Miscellaneous) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.NewContent: %w", err)
		}
	}()

	// Create the components of this panel's content.
	panelContent = &Content{
		screen:      screen,
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		panelContent.description,
{{- range $panelName := .LocalPanelNames }}
 {{- if ne $panelName $DOT.PanelName }}
		widget.NewButton("Switch to the {{ $panelName }} Panel", panelContent.switchTo{{ $panelName }}),
 {{- end }}
{{- end }}
	)
	return
}

// CanvasObject returns the panel's content as a fyne.CanvasObject.
func (panelContent *Content) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panelContent.content
	return
}

{{- range $panelName := .LocalPanelNames }}
 {{- if ne $panelName $DOT.PanelName }}

func (panelContent *Content) switchTo{{ $panelName }}() {
	panelContent.screen.Panelers.{{ $panelName }}.Show()
}
 {{- end }}
{{- end }}
`
)
