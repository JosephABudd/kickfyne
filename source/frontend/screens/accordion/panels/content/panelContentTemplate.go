package content

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type PanelContentTemplateData struct {
	PanelName    string
	PackageName  string
	ImportPrefix string
	Funcs        _utils_.Funcs
}

func PanelContentFileName(panelName string) (panelFileName string) {
	return panelName + ".go"
}

const (
	PanelContentTemplate = `package content

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_misc_ "{{ .ImportPrefix }}/frontend/gui/screens/simple/{{ .PackageName }}/misc"
)

// type {{ .PanelName }}PanelContentAction indicates some changes in state.
// KICKFYNE TODO: Add members for changes that need further attention in {{ .PanelName }}PanelContent.SetState(...).
// See func {{ .PanelName }}PanelContent.SetState(...).
type {{ .PanelName }}PanelContentAction struct {
	// SetHeading is for example only.: No more attention needed for the heading.
	// SetHeading bool
}

// type {{ .PanelName }}PanelContentValues indicates some changes in state.
// KICKFYNE TODO: Add members for changes that need further attention in {{ .PanelName }}PanelContent.SetState(...).
// See func {{ .PanelName }}PanelContent.Get().
type {{ .PanelName }}PanelContentValues struct {
	Heading string
}

type {{ .PanelName }}PanelContentSetter func(actions *{{ .PanelName }}PanelContentAction)

// {{ .PanelName }}PanelContent is the content for the {{ .PanelName }}Panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type {{ .PanelName }}PanelContent struct {
	canvasObject fyne.CanvasObject
	screen  *_misc_.ScreenComponents

	// Widgets.
	// KICKFYNE TODO: Add widgets.
	// Each widget needs:
	// 1. a widget. & optional binding value.
	// 2. a Setter func which returns a {{ .PanelName }}PanelContentSetter
	// 3. a Getter func which returns the widget's value.

	// The heading is the text for the panel's heading.
	headingWidget *widget.Label
	headingWidgetBindingValue binding.ExternalString
	// heading setter.
	// Returns a {{ .PanelName }}PanelContentSetter.
	func (content *{{ .PanelName }}PanelContent) HeadingSetter(text string) {
		return func(actions *{{ .PanelName }}PanelContentAction) {
			content.headingWidgetBindingValue.Set(text)
			// actions.SetHeading is for example only.: No more attention needed for the heading.
			// actions.SetHeading = true
			return
		}
	}
	// heading getter.
	func (content *{{ .PanelName }}PanelContent) headingGetter() (text string) {
			text = content.headingWidgetBindingValue.Get()
			return
		}
	}
}

// New{{ .PanelName }}PanelContent constructs this content.
// Returns the {{ .PanelName }}PanelContent and the error.
func New{{ .PanelName }}PanelContent(screen *_misc_.ScreenComponents) (panelContent *{{ .PanelName }}PanelContent, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.new{{ .PanelName }}: %w", err)
		}
	}()
	
	// 1. Construct the content.
	panelContent = &{{ .PanelName }}PanelContent{
		screen: screen,
		headingWidget: *widget.Label,
	}

	// 2. KICKFYNE TODO: Construct widgets. Set default values.
	panelContent.heading = widget.NewLabelWithData(panelContent.headingWidgetBindingValue)
	panelContent.headingWidgetBindingValue.Set("{{ .PanelName }} Panel")

	// 3. KICKFYNE TODO: Layout the widgets.
	vbox := container.NewVBox(
		panelContent.heading,
	)

	// Here the vertical stack will scroll vertically.
	scroller := container.NewVScroll(vbox)
	panelContent.canvasObject = container.NewMax(scroller)
	return
}

func (panelContent *{{ .PanelName }}PanelContent) Get() (values {{ .PanelName }}PanelContentValues) {
	values = {{ .PanelName }}PanelContentValues{}
	values.Heading = panelContent.headingGetter()

	return
}

// Content returns the panel's content.
func (panelContent *{{ .PanelName }}PanelContent) Content() (content fyne.CanvasObject) {
	content = panelContent.canvasObject
	return
}

// Set sets the content state.
// The widget states are then updated according.
func (panelContent *{{ .PanelName }}PanelContent) Set(setters ...contents.{{ .PanelName }}PanelContentSetter) {
	var action contents.{{ .PanelName }}PanelContentAction
	actions := panelContent.state.Set(setters)

	// By default, no manual updates.
	// The following widgets are bound to their data.
	// 1. panelContent.heading is bound to panelContent.state.heading.
	_= actions

	// KICKFYNE TODO: Manually update widgets.
	// This is only for widgets that are not updated automatically.
	// if actions.UpdatedFooBar {
	//     // update the panelContent.FooBar widget.
	// }
}

// GetState returns a snapshot of state.
func (panelContent *{{ .PanelName }}PanelContent) GetState() (state *{{ .PanelName }}PanelContent) {
	state = panelContent.state.Get()
	return
}

`
)
