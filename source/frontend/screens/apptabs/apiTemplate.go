package simple

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type aPITemplateData struct {
	PackageName             string
	AllPanelNames           []string
	DefaultPanelName        string
	F2BAddTabbarMessageName string
	ImportPrefix            string
	Funcs                   _utils_.Funcs
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
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/producer"
	_tabs_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/tabs"
	_txrx_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/txrx"
	_types_ "{{ .ImportPrefix }}/frontend/types"
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
	// PackageScreen.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, windowContentConsumer, screenID); err != nil {
		return
	}
	// Messenger.
	if messenger, err = _txrx_.NewMessenger(packageScreen); err != nil {
		return
	}
	messenger.Send{{ .F2BAddTabbarMessageName }}()
	startupMessenger = messenger
	err = _tabs_.AddExampleTabs(packageScreen)

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
	var screenID string
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID = fmt.Sprintf("{{ .PackageName }}:TabItem:%d", nextScreenCount())
	// Consumer.
	appTabsTabItemContentConsumer = _types_.NewAppTabsTabItemContentConsumer(appTabs, tabItem)
	// PackageScreen.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID); err != nil {
		return
	}
	// Messenger.
	if messenger, err = _txrx_.NewMessenger(packageScreen); err != nil {
		return
	}
	messenger.Send{{ .F2BAddTabbarMessageName }}()
	startupMessenger = messenger

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
	var screenID string
	var messenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
		}
	}()

	// Screen ID.
	screenID = fmt.Sprintf("{{ .PackageName }}:TabItem:%d", nextScreenCount())
	// Consumer.
	docTabsTabItemContentConsumer = _types_.NewDocTabsTabItemContentConsumer(docTabs, tabItem)
	// PackageScreen.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID); err != nil {
		return
	}
	// Messenger.
	if messenger, err = _txrx_.NewMessenger(packageScreen); err != nil {
		return
	}
	messenger.Send{{ .F2BAddTabbarMessageName }}()
	startupMessenger = messenger

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
	// PackageScreen.
	var packageScreen *_misc_.Miscellaneous
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID); err != nil {
		return
	}
	// Messenger.
	if messenger, err = _txrx_.NewMessenger(packageScreen); err != nil {
		return
	}
	messenger.Send{{ .F2BAddTabbarMessageName }}()
	startupMessenger = messenger
	return
}

func buildLayout(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	screenID string,
) (screen *_misc_.Miscellaneous, err error) {
	// Build the AppTabs content producer.
	appTabsProducer := _producer_.NewAppTabsContentProducer(consumer)
	consumer.Bind(appTabsProducer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(consumer, appTabsProducer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, screenID); err != nil {
		return
	}
	return
}
`
)
