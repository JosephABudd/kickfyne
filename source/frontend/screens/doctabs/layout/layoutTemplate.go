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
	"fyne.io/fyne/v2/container"


	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// Layout this screen's layout of a container.DocTabs.
type Layout struct {
	contentProducer _types_.ContentProducer
	doctabs         *container.DocTabs // This is the canvas object.
	tabItemConsumer map[*container.TabItem]*_types_.DocTabItemContentConsumer
	panelerTabItem  map[_types_.Paneler]*container.TabItem
	panelIDPaneler  map[string]_types_.Paneler
}

// NewLayout constructs this layout.
func NewLayout(producer _types_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	layout = &Layout{
		contentProducer: producer,
		doctabs:         container.NewDocTabs(),
		tabItemConsumer: make(map[*container.TabItem]*_types_.DocTabItemContentConsumer),
		panelerTabItem:  make(map[_types_.Paneler]*container.TabItem),
		panelIDPaneler:  make(map[string]_types_.Paneler),
	}
	producer.SetCanvasObject(layout.doctabs)

	return
}

func (layout *Layout) AddPanelerTabItemConsumer(paneler _types_.Paneler, tabItem *container.TabItem, consumer *_types_.DocTabItemContentConsumer) {
	fyne.DoAndWait(func() { layout.doctabs.Append(tabItem) })
	layout.tabItemConsumer[tabItem] = consumer
	layout.panelerTabItem[paneler] = tabItem
	layout.panelIDPaneler[paneler.ID()] = paneler
}

func (layout *Layout) RemoveTabItem(removeTabItem *container.TabItem) {
	for paneler, tabItem := range layout.panelerTabItem {
		if tabItem == removeTabItem {
			fyne.DoAndWait(func() { layout.doctabs.Remove(removeTabItem) })
			paneler.UnBindCleanUP()
			delete(layout.panelerTabItem, paneler)
			delete(layout.tabItemConsumer, removeTabItem)
			delete(layout.panelIDPaneler, paneler.ID())
			return
		}
	}
}

func (layout *Layout) RemovePaneler(removePaneler _types_.Paneler) {
	var paneler _types_.Paneler
	var tabItem *container.TabItem
	for paneler, tabItem = range layout.panelerTabItem {
		if paneler == removePaneler {
			fyne.DoAndWait(func() { layout.doctabs.Remove(tabItem) })
			paneler.UnBindCleanUP()
			delete(layout.panelerTabItem, paneler)
			delete(layout.tabItemConsumer, tabItem)
			delete(layout.panelIDPaneler, paneler.ID())
			return
		}
	}
}

func (layout *Layout) RemoveID(removeID string) {
	var paneler _types_.Paneler
	var tabItem *container.TabItem
	for paneler, tabItem = range layout.panelerTabItem {
		if paneler.ID() == removeID {
			fyne.DoAndWait(func() { layout.doctabs.Remove(tabItem) })
			paneler.UnBindCleanUP()
			delete(layout.panelerTabItem, paneler)
			delete(layout.tabItemConsumer, tabItem)
			delete(layout.panelIDPaneler, paneler.ID())
			return
		}
	}
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.contentProducer
	return
}

// Refresh refreshes the tabbar tabs only. Not their content.
func (layout *Layout) Refresh() (producer _types_.ContentProducer) {
	layout.doctabs.Refresh()
	return
}

func (layout *Layout) Tabbar() (tabbar *container.DocTabs) {
	tabbar = layout.doctabs
	return
}
`
)
