package simple

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type aPITemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	DefaultPanelName string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	// aPIFileName = _utils_.ScreenFileName
	aPIFileName = _utils_.APIFileName

	aPITemplate = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/layout"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panelers_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panelers"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/producer"
	_txrx_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/txrx"
	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/shared/thread"
)

var screenCount uint = 0
func nextScreenCount() (count uint) {
	count = screenCount
	screenCount++
	return
}

// NewWindowContentConsumer constructs a new screen and returns a window content consumer of the screen's content.
func NewWindowContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	isInMainMenu bool,
) (
	screenID string,
	windowContentConsumer *_types_.WindowContentConsumer,
	startupMessenger _types_.StartupMessenger,
	err error,
){
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewWindowContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID = fmt.Sprintf("{{ .PackageName }}:Window:%d", nextScreenCount())
	// Consumer.
	windowContentConsumer = _types_.NewWindowContentConsumer(window, isInMainMenu)
	// Screen & messenger.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, messenger, err = newScreenMessenger(ctx, ctxCancel, app, window, windowContentConsumer, screenID); err != nil {
		return
	}
	startupMessenger = messenger

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	isMainThread := _thread_.IsMainThread()
	packageScreen.Panelers.DefaultPanel.Show(isMainThread)
	packageScreen.Layout.Producer().Refresh(isMainThread)

	return
}

// NewAppTabsTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewAppTabsTabItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	appTabs *container.AppTabs,
	tabItem *container.TabItem,
) (
	appTabsTabItemContentConsumer *_types_.AppTabsTabItemContentConsumer,
	startupMessenger _types_.StartupMessenger,
	err error,
) {
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewAppTabsTabItemContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID := fmt.Sprintf("{{ .PackageName }}:AppTabsTabItem:%d", nextScreenCount())
	// Consumer.
	appTabsTabItemContentConsumer = _types_.NewAppTabsTabItemContentConsumer(appTabs, tabItem)
	// Screen & messenger.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, messenger, err = newScreenMessenger(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID); err != nil {
		return
	}
	startupMessenger = messenger

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	isMainThread := _thread_.IsMainThread()
	packageScreen.Panelers.DefaultPanel.Show(isMainThread)
	packageScreen.Layout.Producer().Refresh(isMainThread)

	return
}

// NewDocTabsTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewDocTabsTabItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	docTabs *container.DocTabs,
	tabItem *container.TabItem,
) (
	docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer,
	startupMessenger _types_.StartupMessenger,
	err error,
) {
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewDocTabsTabItemContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID := fmt.Sprintf("{{ .PackageName }}:DocTabsTabItem:%d", nextScreenCount())
	// Consumer.
	docTabsTabItemContentConsumer = _types_.NewDocTabsTabItemContentConsumer(docTabs, tabItem)
	// Screen & messenger.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, messenger, err = newScreenMessenger(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID); err != nil {
		return
	}
	startupMessenger = messenger

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	isMainThread := _thread_.IsMainThread()
	packageScreen.Panelers.DefaultPanel.Show(isMainThread)
	packageScreen.Layout.Producer().Refresh(isMainThread)

	return
}

// NewAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewAccordionItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	accordion *widget.Accordion,
	accordionItem *widget.AccordionItem,
) (
		accordionItemContentConsumer *_types_.AccordionItemContentConsumer,
		startupMessenger _types_.StartupMessenger,
		err error,
) {
	var screenID string
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewAccordionItemContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID = fmt.Sprintf("{{ .PackageName }}:AccordionItem:%d", nextScreenCount())
	// Consumer.
	accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordion, accordionItem)
	// Screen & messenger.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, messenger, err = newScreenMessenger(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID); err != nil {
		return
	}

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	isMainThread := _thread_.IsMainThread()
	packageScreen.Panelers.DefaultPanel.Show(isMainThread)
	packageScreen.Layout.Producer().Refresh(isMainThread)
	startupMessenger = messenger
	return
}

func newScreenMessenger(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	id string,
) (screen *_misc_.Miscellaneous, messenger *_txrx_.Messenger, err error) {
	// Build the content & producer.
	producer := _producer_.NewContentProducer(consumer)
	producer.Bind(consumer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(producer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, id); err != nil {
		return
	}

	// Build each panel.
{{- range $panelName := .LocalPanelNames }}
	// {{ $panelName }} panel.
	var {{ call $DOT.Funcs.DeCap $panelName }}Panel *_panels_.{{ $panelName }}Panel
	if {{ call $DOT.Funcs.DeCap $panelName }}Panel, err = _panels_.New{{ $panelName }}Panel(screen); err != nil {
		return
	}
	screen.Panelers.{{ $panelName }} = {{ call $DOT.Funcs.DeCap $panelName }}Panel
{{- end }}

	// Build Panelers.
	screen.Panelers = &_panelers_.Panelers{}
{{- range $panelName := .LocalPanelNames }}
	screen.Panelers.{{ $panelName }} = {{ call $DOT.Funcs.DeCap $panelName }}Panel
{{- end }}
	screen.Panelers.DefaultPanel = {{ call .Funcs.DeCap .DefaultPanelName }}Panel

	// Messenger.
	if messenger, err = _txrx_.NewMessenger(screen); err != nil {
		return
	}

	return
}
`
)
