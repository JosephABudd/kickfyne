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

import(
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/shared/thread"
)

type Getters struct {
	ID          func() string
	Heading     func() string
	Description func() string
}

// State is the state for the {{ .PanelName }} panel.
type State struct {
	id      string
	content *Content

}

// NewState constructs a new content state.
// It may or may not make some initial settings.
func NewState(
	content *Content,
	screenID string,
) (state *State, err error) {
	state = &State{
		content: content,
		id:      screenID + ".{{ .PanelName }}",
	}
	// Initial settings.
	state.initialSet(
		state.SetHeading("{{ .PackageName }} screen : {{ .PanelName }} panel."),
		state.SetDescription("Using this heading and description as examples."),
	)
	return
}

// initialSet sets state for the constructor func NewState.
func (state *State) initialSet(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	for _, setter := range setters {
		setter(isMainThread)
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
		Heading:     state.getHeading,
		Description: state.getDescription,
	}
	return
}

// ID is this panel's id.
func (state *State) getID() (id string) {
	id = state.id
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
