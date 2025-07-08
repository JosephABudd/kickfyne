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
	"fyne.io/fyne/v2/container"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_content_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ .PanelName }}Panel"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// {{ .PanelName }}Panel is a {{ .PanelName }} panel.
// It's content is at {{ .PanelName }}Panel/content.go.
// It's content's state is at {{ .PanelName }}Panel/state.go.
type {{ .PanelName }}Panel struct {
	content   *_content_.Content
	state     *_content_.State
	messenger *_content_.Messenger
	tabItem   *container.TabItem
	screen    *_misc_.Miscellaneous
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.Miscellaneous, tabItemContentConsumer *_types_.AppTabsTabItemContentConsumer, tabItem *container.TabItem) (panel *{{ .PanelName }}Panel, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.New{{ .PanelName }}Panel: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Panel{
		screen:  screen,
		tabItem: tabItem,
	}
	if panel.content, err = _content_.NewContent(tabItemContentConsumer, screen, tabItem, panel); err != nil {
		return
	}
	if panel.state, err = _content_.NewState(
		panel.content,
		screen.ScreenID,
	); err != nil {
		return
	}
	if panel.messenger, err = _content_.NewMessenger(
		panel.screen,
		panel.state,
	); err != nil {
		return
	}

	return
}

// ID returns the panel's id.
func (panel *{{ .PanelName }}Panel) ID() (id string) {
	getters := panel.state.Get().(_content_.Getters)
	id = getters.ID()
	return
}

// Show doesn't do anything becuase this panel is always shown by it's tabItem.
func (panel *{{ .PanelName }}Panel) Show(isMainThread bool) {}

// Producer returns the panel's producer.
// func (panel *{{ .PanelName }}Panel) Producer() (producer *_producer_.ContentProducer) {
func (panel *{{ .PanelName }}Panel) Producer() (producer _types_.ContentProducer) {
	producer = panel.content.Producer()
	return
}

// CanvasObject returns the panel's content.
func (panel *{{ .PanelName }}Panel) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panel.content.CanvasObject()
	return
}

// Required cleanup after unbinding.
func (panel *{{ .PanelName }}Panel) UnBindCleanUP() {
	panel.messenger.StopReceiving()
}

// State returns the panel's state.
func (panel *{{ .PanelName }}Panel) State() (state _types_.Stater) {
	state = panel.state
	return
}
`
)
