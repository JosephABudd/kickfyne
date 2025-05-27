package accordion

import _utils_ "github.com/JosephABudd/kickfyne/source/utils"

type accordionLayoutTemplateData struct {
	PackageName      string
	PanelNames       []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
	LocalPanelNames  []string
	RemotePanelNames []string
}

const (
	accordionLayoutFileName = "layout.go"

	accordionLayoutTemplate = `package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/types"

{{- if eq (len .RemotePanelNames) 1}}
	// A remote screen which is content for a tab.
{{- else }}
{{- if gt (len .RemotePanelNames) 1}}
	// Remote screens which are content for tabs.
{{- end }}
{{- $panelName := range .RemotePanelNames }}
	_{{ $panelName }}_ = "{{ .ImportPrefix }}/frontend/screens/{{ $panelName }}"
{{- end }}
)

// LayoutComponents is an Accordion layout.
// It is this screen's layout and content.
type LayoutComponents struct {
	content        fyne.CanvasObject
	screen         *screenComponents
	screenWatchers map[*widget.AccordionItem]*_types_.AccordionItemScreenCanvasObjectWatcher
}

// newLayoutComponents constructs this layout.
// It creates the accordion that makes up the layout.
// Returns the layout and error.
func newLayoutComponents(screen *screenComponents) (layout *LayoutComponents, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.New{{ call .Funcs.Cap .PanelName }}: %w", err)
		}
	}()

	layout = &LayoutComponents{
		screen:         screen,
		screenWatchers: make(map[*widget.AccordionItem]*_types_.AccordionItemScreenCanvasObjectWatcher),
	}

	// Build the accordionItems.
	var accordionItems []*widget.AccordionItem
	if accordionItems, err = layout.accordionItems(screen.ctx, screen.ctxCancel, screen.app, screen.window); err != nil {
		return
	}
	// Build the accordion.
	accordion := widget.NewAccordion(
		accordionItems...,
	)
	// Build the layout content.
	layout.content = container.New(
		layout.NewMaxLayout(),
		accordion,
	)
	return
}

func (layout *LayoutComponents) accordionItems(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (items []*widget.AccordionItem, err error) {

	defer func() {
		if len(items) == 0 {
			err = fmt.Errorf("the Accordion in {{ .PackageName }}.{{ .PanelName }} must have items")
		}
		if err != nil {
			err = fmt.Errorf("layout.accordionItems: %w", err)
		}
	}()

	// Accordion items using a local panel for content.
	items = make([]*widget.AccordionItem, 0, {{ len .LocalPanelNames }})
{{- $panelName := range .LocalPanelNames }}
	FOR EACH ACCORDION ITEM USING CONTENT FROM ANOTHER PANEL IN THIS SCREEN, use the 1 step example code below.
	items = append(
		items,
		widget.NewAccordionItem("{{ $panelName }}", layout.screen.panels.{{ $panelName }}.content),
	)
{{- end }}

	// Accordion items using another screen for content.
{{- $panelName := range .RemotePanelNames }}
	var otherScreen _types_.CanvasObjectProvider
	// Construct the other screen package.
	if otherScreen, err = {{ $panelName }}.New(layout.screen.ctx, layout.screen.ctxCancel, layout.screen.app, layout.screen.window); err != nil {
		return
	}
	// Build and add the accordionItem with the other screen's canvas object provider.
	items = append(
		items,
		layout.addScreenWatcherItem("{{ $panelName }}", otherScreen),
	)

{{- end }}
	return
}

// addScreenWatcherItem creates and adds an AccordionItem with a canvas object provided by another screen.
func (layout *LayoutComponents) addScreenWatcherItem(label string, otherScreen _types_.CanvasObjectProvider) (accordionItem *widget.AccordionItem) {
	var watcher *_types_.AccordionItemScreenCanvasObjectWatcher
	accordionItem, watcher = _types_.NewAccordionItemScreenCanvasObjectWatcher(label, otherScreen)
	layout.screenWatchers[accordionItem] = watcher
	return
}

// Show shows this layout.
func (layout *LayoutComponents) Show() {
	layout.screen.canvasObjectProvider.UpdateCanvasObject(layout.content)
}

`
)
