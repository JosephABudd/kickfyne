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
	_txrx_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/txrx"
	_types_ "{{ .ImportPrefix }}/frontend/types"
	_message_ "{{ .ImportPrefix }}/shared/message"
)


type SpawnData = _misc_.SpawnData

// NewWindowContentConsumer constructs a new screen and returns a window content consumer of the screen's content.
func NewWindowContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (windowContentConsumer *_types_.WindowContentConsumer, err error) {
	var packageScreen *_misc_.Miscellaneous
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewWindowContentConsumer: %w", err)
		} else {
			// IS THIS NEEDED?
			// packageScreen.Layout.Producer().Refresh()
			_ = packageScreen
		}
	}()

	const notSpawned bool = false
	windowContentConsumer = _types_.NewWindowContentConsumer(w)
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, w, windowContentConsumer, notSpawned); err != nil {
		return
	}
	err = announceNew{{ .PackageName }}(packageScreen)
	return
}

// NewTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewDocTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabbar *container.DocTabs, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer) (tabItemConsumer *_types_.DocTabItemContentConsumer, err error) {
	tabItemConsumer, err = newDocTabItemContentConsumer(ctx, ctxCancel, app, w, tabbar, tabItem, screenConsumer, nil)
	return
}

// NewSpawnedTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewSpawnedDocTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabbar *container.DocTabs, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer, spawnData *SpawnData) (tabItemConsumer *_types_.DocTabItemContentConsumer, err error) {
	tabItemConsumer, err = newDocTabItemContentConsumer(ctx, ctxCancel, app, w, tabbar, tabItem, screenConsumer, spawnData)
	return
}

// newDocTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func newDocTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabbar *container.DocTabs, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer, spawnData *SpawnData) (tabItemConsumer *_types_.DocTabItemContentConsumer, err error) {
	var spawned bool = (spawnData != nil)
	var packageScreen *_misc_.Miscellaneous
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newDocTabItemContentConsumer: %w", err)
		}
	}()

	tabItemConsumer = _types_.NewDocTabItemContentConsumer(tabbar, tabItem, spawned)
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, w, tabItemConsumer, spawned); err != nil {
		return
	}
	err = announceNew{{ .PackageName }}(packageScreen)

	return
}

// NewAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, err error) {
	accordionItemContentConsumer, err = newAccordionItemContentConsumer(ctx, ctxCancel, app, w, accordionItem, screenConsumer, nil)
	return
}

// NewSpawnedAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewSpawnedAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer, spawnData *SpawnData) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, messenger *_txrx_.Messenger, err error) {
	accordionItemContentConsumer, err = newAccordionItemContentConsumer(ctx, ctxCancel, app, w, accordionItem, screenConsumer, spawnData)
	return
}

// newAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func newAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer, spawnData *SpawnData) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, err error) {
	var spawned = (spawnData != nil)
	var packageScreen *_misc_.Miscellaneous
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newAccordionItemContentConsumer: %w", err)
		}
	}()

	accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordionItem, screenConsumer, spawned)
	if packageScreen, err = buildLayout(ctx, ctxCancel, app, w, accordionItemContentConsumer, spawned); err != nil {
		return
	}
	err = announceNew{{ .PackageName }}(packageScreen)
	return
}

func buildLayout(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, consumer _types_.ContentConsumer, spawned bool) (screen *_misc_.Miscellaneous, err error) {
	// Build the content producer.
	producer := _types_.NewScreenContentProducer(spawned, consumer)
	consumer.Bind(producer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(producer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, w, layout); err != nil {
		return
	}
	return
}

func announceNew{{ .PackageName }}(packageScreen *_misc_.Miscellaneous) (err error) {
	var messenger *_txrx_.Messenger
	if messenger, err = _txrx_.NewMessenger(packageScreen); err != nil {
		return
	}
	// Tell the back-end that this {{ .PackageName }} screen is active.
	messenger.SendOpenedNew{{ .PackageName }}(_message_.{{ .PackageName }}RemoteData{})
	return
}
`
)
