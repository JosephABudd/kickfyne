package screenmap

type screenMapTemplateData struct {
	ImportPrefix string
}

const (
	screenMapFileName = "screenmap.go"
	screenMapTemplate = `{{ $DOT := . -}}
package screenmap

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type API struct {
	NewWindowContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		isInMainMenu bool,
	) (
		screenID string,
		windowContentConsumer *_types_.WindowContentConsumer,
		startupMessenger _types_.StartupMessenger,
		err error,
	)

	NewAppTabsTabItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		tabbar *container.AppTabs,
		tabItem *container.TabItem,
	) (
		docTabsTabItemContentConsumer *_types_.AppTabsTabItemContentConsumer,
		startupMessenger _types_.StartupMessenger,
		err error,
	)

	NewDocTabsTabItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		tabbar *container.DocTabs,
		tabItem *container.TabItem,
	) (
		docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer,
		startupMessenger _types_.StartupMessenger,
		err error,
	)

	NewAccordionItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		accordion *widget.Accordion,
		accordionItem *widget.AccordionItem,
	) (
		accordionItemContentConsumer *_types_.AccordionItemContentConsumer,
		startupMessenger _types_.StartupMessenger,
		err error,
	)

}

var Map = make(map[string]*API)
`
)
