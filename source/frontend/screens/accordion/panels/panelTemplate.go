package panels

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type PanelTemplateData struct {
	PanelName    string
	PackageName  string
	ImportPrefix string
	Funcs        _utils_.Funcs
}

func PanelFileName(panelName string) (panelFileName string) {
	return panelName + ".go"
}

const (
	PanelTemplate = `package {{ .PackageName }}

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	_content_ "{{ .ImportPrefix }}/frontend/gui/screens/simple/{{ .PackageName }}/panels/content"
	_misc_ "{{ .ImportPrefix }}/frontend/gui/screens/simple/{{ .PackageName }}/misc"
)

// {{ .PanelName }}Panel is a panel.
// It's content is at content/{{ .PanelName }}Content.go
// KICKFYNE TODO: Correct this panel's doc comment.
type {{ .PanelName }}Panel struct {
	content *_content_.{{ .PanelName }}PanelContent
	screen  *_misc_.ScreenComponents
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.ScreenComponents) (panel *{{ .PanelName }}Panel, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.new{{ .PanelName }}: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Panel{
		screen: screen,
	}
	if panel.content, err = _content_.New{{ .PanelName }}PanelContent(screen); err != nil {
		return error
	}
	return
}

// SetState sets the state of this panel's content.
func (panel *{{ .PanelName }}Panel) SetState(setters ...Content.{{ .PanelName }}PanelContentStateSetter) {
	panel.content.Set(setters)
}

// GetState gets a copy the state of this panel's content.
func (panel *{{ .PanelName }}Panel) GetState(state *{{ .PanelName }}PanelContentState) {
	state = panel.content.Get()
	return
}

// Show shows this panel and hides the others.
func (panel *{{ .PanelName }}Panel) Show() {
	panel.screen.canvasObjectProvider.UpdateCanvasObject(panel.Content())
}

// Content returns the panel's content.
func (panel *{{ .PanelName }}Panel) Content() (content fyne.CanvasObject) {
	content = panel.content.Content()
	return
}

`
)
