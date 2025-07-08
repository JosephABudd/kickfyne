package panel

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type StateTemplateData struct {
	PackageName  string
	PanelName    string
	ImportPrefix string
}

const (
	StateFileName = _utils_.StateFileName

	StateTemplate = `package {{ .PanelName }}Panel

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/shared/thread"
)

type Getters struct {
	ID          func() string
	TabLabel    func() string
	TabIcon     func() fyne.Resource
	Heading     func() string
	Description func() string
}

// State is the state for the {{ .PanelName }} panel.
// Panel, Tab, & Messenger have the state id.
type State struct {
	id       string
	tabLabel string
	tabIcon  fyne.Resource
	content  *Content
}

// NewState constructs a new content state.
// It may or may not make some initial settings.
func NewState(
	content  *Content,
	screenID string,
) (state *State, err error) {
	state = &State{
		id:      screenID + ".{{ .PanelName }}",
		content: content,
	}
	// Initial settings.
	state.initialSet(
		state.SetTabIcon(nil),
		state.SetTabLabel("{{ .PanelName }}"),
		state.SetHeading("Tabbar screen : {{ .PanelName }} panel."),
		state.SetDescription("Using this heading and description as examples."),
	)
	return
}

// initialSet sets state for the constructor func NewState.
func (state *State) initialSet(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	for _, setter := range setters {
		_ = setter(isMainThread)
	}
}

// Set sets the state.
func (state *State) Set(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	var refreshCanvasObject bool
	for _, setter := range setters {
		refreshCanvasObject = refreshCanvasObject || setter(isMainThread)
	}
	producer := state.content.screen.Layout.Producer()
	if refreshCanvasObject {
		producer.SetCanvasObject(state.content.content)
	}
	producer.Refresh(isMainThread)
}

func (state *State) Get() (getters any) {
	getters = Getters{
		ID:          state.getID,
		TabLabel:    state.getTabLabel,
		TabIcon:     state.getTabIcon,
		Heading:     state.getHeading,
		Description: state.getDescription,
	}
	return
}

// The panel, tab, & messenger use this for an ID.
func (state *State) getID() (id string) {
	id = state.id
	return
}

// Tab label.
func (state *State) SetTabLabel(label string) (setter _types_.StateSetter) {
	state.tabLabel = label
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.producer.SetLabel(state.tabLabel)
		} else {
			fyne.Do(
				func() {
					state.content.producer.SetLabel(state.tabLabel);
				},
			)
		}
		return
	}
	return
}

func (state *State) getTabLabel() (label string) {
	label = state.tabLabel
	return
}

// Tab icon.
func (state *State) SetTabIcon(icon fyne.Resource) (setter _types_.StateSetter) {
	state.tabIcon = icon
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.producer.SetIcon(state.tabIcon)
		} else {
			fyne.Do(
				func() {
					state.content.producer.SetIcon(state.tabIcon);
				},
			)
		}
		return
	}
	return
}

func (state *State) getTabIcon() (icon fyne.Resource) {
	icon = state.tabIcon
	return
}

// Heading is a widget with variable state.

// SetHeading returns a _types_.Setter that sets the content's heading widget's text.
func (state *State) SetHeading(heading string) (setter _types_.StateSetter) {
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.heading.Text = heading
		} else {
			fyne.Do(
				func() {
					state.content.heading.Text = heading;
				},
			)
		}
		refreshCanvasObject = true
		return
	}
	return
}

func (state *State) getHeading() (heading string) {
	heading = state.content.heading.Text
	return
}

// Description is a widget with variable state.

// SetDescription returns a _types_.Setter that sets the content's description widget's text.
func (state *State) SetDescription(description string) (setter _types_.StateSetter) {
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.description.Text = description
		} else {
			fyne.Do(
				func() {
					state.content.description.Text = description;
				},
			)
		}
		refreshCanvasObject = true
		return
	}
	return
}

func (state *State) getDescription() (description string) {
	description = state.content.description.Text
	return
}
`
)
