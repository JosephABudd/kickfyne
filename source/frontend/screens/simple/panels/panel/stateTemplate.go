package screens

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
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type Getters struct {
	Heading     func() string
	Description func() string
}

// State is the state for the {{ .PanelName }} panel.
type State struct {
	content *Content
	refresh func()

}

// NewState constructs a new content state.
// It may or may not make some initial settings.
func NewState(content *Content, refresh func()) (state *State, err error) {
	state = &State{
		content: content,
		refresh: refresh,
	}
	// Initial settings.
	state.initialSet(
		state.SetHeading("{{ .PackageName }} screen : {{ .PanelName }} panel."),
		state.SetDescription("Using this heading and description as an examples."),
	)
	return
}

// initialSet sets state for the constructor func NewState.
func (state *State) initialSet(setters ..._types_.StateSetter) {
	for _, setter := range setters {
		setter()
	}
}

// Set sets the state.
func (state *State) Set(setters ..._types_.StateSetter) {
	var refreshCanvasObject bool
	for _, setter := range setters {
		if setter() {
			refreshCanvasObject = true
		}
	}
	if refreshCanvasObject {
		state.content.producer.SetCanvasObject(state.content.content)
	}
	state.content.producer.Refresh()
}

func (state *State) Get() (getters any) {
	getters = Getters{
		Heading:     state.getHeading,
		Description: state.getDescription,
	}
	return
}

// Heading is a widget with variable state.

// SetHeading returns a _types_.Setter that sets the content's heading widget's text.
func (state *State) SetHeading(heading string) (setter _types_.StateSetter) {
	setter = func() {
		state.content.heading.Text = heading
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
	setter = func() {
		state.content.description.Text = description
	}
	return
}

func (state *State) getDescription() (description string) {
	description = state.content.description.Text
	return
}
`
)
