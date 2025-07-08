package accordionitems

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	FileName = "accordionItems.go"

	Template = `{{ $DOT := . -}}
package accordionitems

import (
	"fyne.io/fyne/v2/widget"

	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

func AddExampleAccordionItems(screen *_misc_.Miscellaneous) (err error) {
	var tabMessengerID string
{{- range $panelName := .LocalPanelNames}}
	if tabMessengerID, err = Open{{ $panelName }}AccordionItem(screen, "{{ $panelName }}"); err != nil {
		return
	}
{{- end }}
	_ = tabMessengerID

	var startupMessenger _types_.StartupMessenger
{{ range $panelName := .RemotePanelNames}}
	if startupMessenger, err = Open{{ $panelName }}AccordionItem(screen, "{{ $panelName }}"); err != nil {
		return
	}
{{- end }}
	_ = startupMessenger

	return
}

{{- range $panelName := .LocalPanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, accordionItemTitle string) (tabMessengerID string, err error) {
	accordionItem := widget.NewAccordionItem(accordionItemTitle, widget.NewLabel("This is the {{ $panelName }} panel."))
	accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.Accordion(), accordionItem)
	// accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.AccordionConsumer(), screen.Layout.Accordion(), accordionItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, accordionItemContentConsumer, accordionItem); err != nil {
		return
	}
	panelProducer := panel.Producer()
	accordionItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerAccordionItemConsumer(panel, accordionItem, accordionItemContentConsumer)
	tabMessengerID = panel.ID()

	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, title string) (startupMessenger _types_.StartupMessenger, err error) {
	accordionItem := widget.NewAccordionItem(title, widget.NewLabel("This is the {{ $panelName }} panel."))
	api := _screenmap_.Map["{{ $panelName }}"]
	var accordionItemContentConsumer *_types_.AccordionItemContentConsumer

	accordionItemContentConsumer, startupMessenger, err = api.NewAccordionItemContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		screen.Layout.Accordion(),
		accordionItem,
	)
	if err == nil {
		screen.Layout.AddPanelerAccordionItemConsumer(nil, accordionItem, accordionItemContentConsumer)
	}
	return
}
{{- end }}

func CloseAccordionItem(screen *_misc_.Miscellaneous, accordionItem *widget.AccordionItem) {
	if id, found := screen.Layout.AccordionItemID(accordionItem); found {
		screen.Layout.RemoveID(id)
	}
}

func CloseAccordionItemID(screen *_misc_.Miscellaneous, id string) {
	screen.Layout.RemoveID(id)
}

`
)
