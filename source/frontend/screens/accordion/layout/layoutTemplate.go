package misc

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type LayoutTemplateData struct {
	PackageName      string
	AllPanelNames    []string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/shared/thread"
)

// Layout this screen's layout of a widget.AccordionItem.
type Layout struct {
	accordionConsumer             _types_.ContentConsumer
	accordionItemContentProducer  _types_.ContentProducer
	accordion                     *widget.Accordion // This is the canvas object.
	accordionItemConsumer         map[*widget.AccordionItem]*_types_.AccordionItemContentConsumer
	accordionItemPaneler          map[*widget.AccordionItem]_types_.Paneler
	panelIDPaneler                map[string]_types_.Paneler
	panelIDAccordionItem          map[string]*widget.AccordionItem
}

// NewLayout constructs this layout.
func NewLayout(accordionConsumer _types_.ContentConsumer, accordionItemContentProducer _types_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	layout = &Layout{
		accordionConsumer:            accordionConsumer,
		accordionItemContentProducer: accordionItemContentProducer,
		accordion:                    widget.NewAccordion(),

		accordionItemConsumer:        make(map[*widget.AccordionItem]*_types_.AccordionItemContentConsumer),
		accordionItemPaneler:         make(map[*widget.AccordionItem]_types_.Paneler),
		panelIDPaneler:               make(map[string]_types_.Paneler),
		panelIDAccordionItem:         make(map[string]*widget.AccordionItem),
	}
	accordionItemContentProducer.SetCanvasObject(layout.accordion)

	return
}

func (layout *Layout) AccordionConsumer() (accordionConsumer _types_.ContentConsumer) {
	accordionConsumer = layout.accordionConsumer
	return
}

func (layout *Layout) AddPanelerAccordionItemConsumer(paneler _types_.Paneler, accordionItem *widget.AccordionItem, consumer *_types_.AccordionItemContentConsumer) {
	if _thread_.IsMainThread() {
		layout.accordion.Append(accordionItem)
	} else {
		fyne.DoAndWait(func() { layout.accordion.Append(accordionItem) })
	}
	layout.accordionItemConsumer[accordionItem] = consumer
	layout.accordionItemPaneler[accordionItem] = paneler
	if paneler != nil {
		panelID := paneler.ID()
		layout.panelIDPaneler[panelID] = paneler
		layout.panelIDAccordionItem[panelID] = accordionItem
	}
}

func (layout *Layout) AccordionItemID(accordionItem *widget.AccordionItem) (id string, found bool) {
	if paneler := layout.accordionItemPaneler[accordionItem]; paneler != nil {
		id = paneler.ID()
		found = true
	}
	return
}

func (layout *Layout) RemoveID(removeID string) {
	var paneler _types_.Paneler
	var accordionItem *widget.AccordionItem
	if paneler = layout.panelIDPaneler[removeID]; paneler == nil {
		return
	}
	accordionItem = layout.panelIDAccordionItem[removeID]
	layout.accordion.Remove(accordionItem)
	paneler.UnBindCleanUP()
	delete(layout.accordionItemPaneler, accordionItem)
	delete(layout.accordionItemConsumer, accordionItem)
	panelID := paneler.ID()
	delete(layout.panelIDPaneler, panelID)
	delete(layout.panelIDAccordionItem, panelID)
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.accordionItemContentProducer
	return
}

// Refresh refreshes the accordion tabs only. Not their content.
func (layout *Layout) Refresh() {
	layout.accordion.Refresh()
}

func (layout *Layout) Accordion() (accordion *widget.Accordion) {
	accordion = layout.accordion
	return
}
`
)
