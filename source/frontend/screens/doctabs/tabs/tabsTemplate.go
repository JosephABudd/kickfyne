package tabs

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
	FileName = "tabs.go"

	Template = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

func AddExampleTabs(screen *_misc_.Miscellaneous) (err error) {
	var tabMessengerID string
{{- range $panelName := .LocalPanelNames}}
	if tabMessengerID, err = Open{{ $panelName }}Tab(screen, nil, "{{ $panelName }}"); err != nil {
		return
	}
{{- end }}
	_ = tabMessengerID

	var startupMessenger _types_.StartupMessenger
{{ range $panelName := .RemotePanelNames}}
	if startupMessenger, err = Open{{ $panelName }}Tab(screen, nil, "{{ $panelName }}"); err != nil {
		return
	}
{{- end }}
	_ = startupMessenger

	return
}

{{- range $panelName := .LocalPanelNames}}

// Open{{ $panelName }}Tab constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}Tab(screen *_misc_.Miscellaneous, tabIcon fyne.Resource, tabLabel string) (tabMessengerID string, err error) {
	tabItem := container.NewTabItemWithIcon(tabLabel, tabIcon, widget.NewLabel("This is the {{ $panelName }} panel."))
	tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.Tabbar(), tabItem)
	// tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.TabbarConsumer(), screen.Layout.Tabbar(), tabItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, tabItemContentConsumer, tabItem); err != nil {
		return
	}
	panelProducer := panel.Producer()
	tabItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerTabItemConsumer(panel, tabItem, tabItemContentConsumer)
	tabMessengerID = panel.ID()

	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}Tab constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}Tab(screen *_misc_.Miscellaneous, tabIcon fyne.Resource, tabLabel string) (startupMessenger _types_.StartupMessenger, err error) {
	tabItem := container.NewTabItemWithIcon(tabLabel, tabIcon, widget.NewLabel("This is the {{ $panelName }} panel."))
	api := _screenmap_.Map["{{ $panelName }}"]
	var docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer

	docTabsTabItemContentConsumer, startupMessenger, err = api.NewDocTabsTabItemContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		screen.Layout.Tabbar(),
		tabItem,
	)
	if err == nil {
		screen.Layout.AddPanelerTabItemConsumer(nil, tabItem, docTabsTabItemContentConsumer)
	}
	return
}
{{- end }}

func CloseTabItem(screen *_misc_.Miscellaneous, tabItem *container.TabItem) {
	if id, found := screen.Layout.TabID(tabItem); found {
		screen.Layout.RemoveID(id)
	}
}

func CloseTabID(screen *_misc_.Miscellaneous, id string) {
	screen.Layout.RemoveID(id)
}

`
)
