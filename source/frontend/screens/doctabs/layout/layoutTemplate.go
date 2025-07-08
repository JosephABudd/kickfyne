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
	_thread_ "{{ .ImportPrefix }}/shared/thread"
)

// Layout this screen's layout of a container.DocTabs.
type Layout struct {
	tabbarConsumer         _types_.ContentConsumer
	tabItemContentProducer _types_.ContentProducer
	docTabs                *container.DocTabs // This is the canvas object.
	tabItemConsumer        map[*container.TabItem]*_types_.DocTabsTabItemContentConsumer
	tabItemPaneler         map[*container.TabItem]_types_.Paneler
	panelIDPaneler         map[string]_types_.Paneler
	panelIDTabItem         map[string]*container.TabItem
}

// NewLayout constructs this layout.
func NewLayout(tabbarConsumer _types_.ContentConsumer, tabItemContentProducer _types_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	docTabs := container.NewDocTabs()
	layout = &Layout{
		tabbarConsumer:         tabbarConsumer,
		tabItemContentProducer: tabItemContentProducer,
		docTabs:                docTabs,
		tabItemConsumer:        make(map[*container.TabItem]*_types_.DocTabsTabItemContentConsumer),
		tabItemPaneler:         make(map[*container.TabItem]_types_.Paneler),
		panelIDPaneler:         make(map[string]_types_.Paneler),
		panelIDTabItem:         make(map[string]*container.TabItem),
	}
	docTabs.CloseIntercept = func(tabItem *container.TabItem) {
		var id string
		var found bool
		if id, found = layout.TabID(tabItem); found {
			layout.RemoveID(id)
		}
	}
	tabItemContentProducer.SetCanvasObject(layout.docTabs)

	return
}

func (layout *Layout) TabbarConsumer() (tabbarConsumer _types_.ContentConsumer) {
	tabbarConsumer = layout.tabbarConsumer
	return
}

func (layout *Layout) AddPanelerTabItemConsumer(paneler _types_.Paneler, tabItem *container.TabItem, consumer *_types_.DocTabsTabItemContentConsumer) {
	if _thread_.IsMainThread() {
		layout.docTabs.Append(tabItem)
	} else {
		fyne.DoAndWait(func() { layout.docTabs.Append(tabItem) })
	}
	layout.tabItemConsumer[tabItem] = consumer
	layout.tabItemPaneler[tabItem] = paneler
	if paneler != nil {
		panelID := paneler.ID()
		layout.panelIDPaneler[panelID] = paneler
		layout.panelIDTabItem[panelID] = tabItem
	}
}

func (layout *Layout) TabID(tabItem *container.TabItem) (id string, found bool) {
	if paneler := layout.tabItemPaneler[tabItem]; paneler != nil {
		id = paneler.ID()
		found = true
	}
	return
}

func (layout *Layout) RemoveID(removeID string) {
	var paneler _types_.Paneler
	var tabItem *container.TabItem
	if paneler = layout.panelIDPaneler[removeID]; paneler == nil {
		return
	}
	tabItem = layout.panelIDTabItem[removeID]
	layout.docTabs.Remove(tabItem)
	paneler.UnBindCleanUP()
	delete(layout.tabItemPaneler, tabItem)
	delete(layout.tabItemConsumer, tabItem)
	panelID := paneler.ID()
	delete(layout.panelIDPaneler, panelID)
	delete(layout.panelIDTabItem, panelID)
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.tabItemContentProducer
	return
}

// Refresh refreshes the tabbar tabs only. Not their content.
func (layout *Layout) Refresh() {
	layout.docTabs.Refresh()
}

func (layout *Layout) Tabbar() (tabbar *container.DocTabs) {
	tabbar = layout.docTabs
	return
}
`
)
