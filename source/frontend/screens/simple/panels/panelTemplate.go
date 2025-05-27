package panels

import (
	"github.com/JosephABudd/kickfyne/source/utils"
)

type PanelTemplateData struct {
	PanelName    string
	PackageName  string
	ImportPrefix string
	Funcs        utils.Funcs
}

const (
	PanelFileNameSuffix = "Panel.go"

	PanelTemplate = `package panels

import (
	"fmt"

	"fyne.io/fyne/v2"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_content_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ .PanelName }}Panel"
	_txrx_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/txrx"
	_types_ "example.com/okp/frontend/types"
)


// {{ .PanelName }}Panel is a {{ .PanelName }} panel.
// It's content is at {{ .PanelName }}Panel/content.go.
// It's content's state is at {{ .PanelName }}Panel/state.go.
type {{ .PanelName }}Panel struct {
	content   *_content_.Content
	state     *_content_.State
	screen    *_misc_.Miscellaneous
	messenger *_txrx_.Messenger
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.Miscellaneous) (panel *{{ .PanelName }}Panel, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.New{{ .PanelName }}Panel: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Panel{
		screen: screen,
	}
	if panel.content, err = _content_.NewContent(screen); err != nil {
		return
	}
	if panel.state, err = _content_.NewState(panel.content, panel.screen.Layout.Producer().Refresh); err != nil {
		return
	}

	return
}

// SetMessenger does just that.
func (panel *{{ .PanelName }}Panel) SetMessenger(messenger *_txrx_.Messenger) {
	panel.messenger = messenger
}

// Show shows this panel and hides the others.
func (panel *{{ .PanelName }}Panel) Show() {
	panel.screen.Layout.Set{{ .PanelName }}PanelCanvasObject(panel.content.CanvasObject())
	panel.screen.Layout.Producer().Refresh()
}

// Returns the panel's state.
func (panel *{{ .PanelName }}Panel) State() (state _types_.Stater) {
	state = panel.state
	return
}

// CanvasObject returns the panel's content.
func (panel *{{ .PanelName }}Panel) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panel.content.CanvasObject()
	return
}
`
)
