package panel

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
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/producer"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	producer *_producer_.AccordionItemContentProducer
	content  fyne.CanvasObject
	screen   *_misc_.Miscellaneous
	accordionItem  *widget.AccordionItem

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
}

// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(accordionItemContentConsumer *_types_.AccordionItemContentConsumer, screen *_misc_.Miscellaneous, accordionItem *widget.AccordionItem, paneler _types_.Paneler) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.NewContent: %w", err)
		}
	}()

	_ = paneler

	// Create the components of this panel's content.
	panelContent = &Content{
		producer:      _producer_.NewAccordionItemContentProducer(accordionItemContentConsumer),
		screen:        screen,
		accordionItem: accordionItem,

		// Widgets with variable state. See state.go.
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		panelContent.description,
	)
	return
}

// CanvasObject returns the panel's content as a fyne.CanvasObject.
func (panelContent *Content) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panelContent.content
	return
}

func (panelContent *Content) Producer() (producer _types_.ContentProducer) {
	producer = panelContent.producer
	return
}
`
)
