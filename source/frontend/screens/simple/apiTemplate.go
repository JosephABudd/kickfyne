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
)

var unSpawnedPackageScreen *_misc_.Miscellaneous
var unSpawnedPackageMessenger *_txrx_.Messenger
var unSpawnedWindowConsumer *_types_.WindowContentConsumer
var initErr error

// NewWindowContentConsumer constructs a new screen and returns a window content consumer of the screen's content.
func NewWindowContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (consumer *_types_.WindowContentConsumer, err error) {
	if initErr != nil {
		err = initErr
		return
	}

	defer func() {
		if err != nil {
			unSpawnedPackageScreen = nil
			unSpawnedPackageMessenger = nil
			err = fmt.Errorf("{{ .PackageName }}.New: %w", err)
			initErr = err
		} else {
			initErr = nil
			consumer = unSpawnedWindowConsumer
			unSpawnedPackageScreen.Layout.Producer().Refresh()
		}
	}()

	if unSpawnedWindowConsumer != nil {
		// This unspawned packages has already been constructed so reuse it.
		return
	}

	const notSpawned bool = false
	unSpawnedWindowConsumer = _types_.NewWindowContentConsumer(w)
	if unSpawnedPackageScreen, unSpawnedPackageMessenger, err = newScreenMessenger(ctx, ctxCancel, app, w, notSpawned, unSpawnedWindowConsumer); err != nil {
		return
	}

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	unSpawnedPackageScreen.Panelers.DefaultPanel.Show()

	return
}

// NewTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer) (tabItemConsumer *_types_.TabItemContentConsumer, err error) {
	tabItemConsumer, _, err = newTabItemContentConsumer(ctx, ctxCancel, app, w, tabItem, screenConsumer, false)
	return
}

// NewSpawnedTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewSpawnedTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer) (tabItemConsumer *_types_.TabItemContentConsumer, messenger *_txrx_.Messenger, err error) {
	tabItemConsumer, messenger, err = newTabItemContentConsumer(ctx, ctxCancel, app, w, tabItem, screenConsumer, true)
	return
}

// newTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func newTabItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, tabItem *container.TabItem, screenConsumer _types_.ContentConsumer, spawned bool) (tabItemConsumer *_types_.TabItemContentConsumer, messenger *_txrx_.Messenger, err error) {
	if !spawned {
		if initErr != nil {
			err = initErr
			return
		}

		if unSpawnedPackageScreen != nil {
			// unSpawnedPackageScreen.Layout.producer already exists.
			// Bind this tabItem consumer and the producer.
			tabItemConsumer = _types_.NewTabItemContentConsumer(tabItem, spawned)
			messenger = unSpawnedPackageMessenger
			producer := unSpawnedPackageScreen.Layout.Producer()
			tabItemConsumer.Bind(producer)
			return
		}
	}

	var packageScreen *_misc_.Miscellaneous
	var packageMessenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
			initErr = err
		} else {
			if !spawned {
				unSpawnedPackageMessenger = packageMessenger
				unSpawnedPackageScreen = packageScreen
			}
			initErr = nil
		}
	}()

	tabItemConsumer = _types_.NewTabItemContentConsumer(tabItem, spawned)
	if packageScreen, packageMessenger, err = newScreenMessenger(ctx, ctxCancel, app, w, spawned, tabItemConsumer); err != nil {
		return
	}

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	packageScreen.Panelers.DefaultPanel.Show()

	return
}

// NewAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, err error) {
	accordionItemContentConsumer, _, err = newAccordionItemContentConsumer(ctx, ctxCancel, app, w, accordionItem, screenConsumer, false)
	return
}

// NewSpawnedAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewSpawnedAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, messenger *_txrx_.Messenger, err error) {
	accordionItemContentConsumer, messenger, err = newAccordionItemContentConsumer(ctx, ctxCancel, app, w, accordionItem, screenConsumer, true)
	return
}

// newAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func newAccordionItemContentConsumer(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, accordionItem *widget.AccordionItem, screenConsumer _types_.ContentConsumer, spawned bool) (accordionItemContentConsumer *_types_.AccordionItemContentConsumer, messenger *_txrx_.Messenger, err error) {
	if !spawned {
		if initErr != nil {
			err = initErr
			return
		}

		if unSpawnedPackageScreen != nil {
			// unSpawnedPackageScreen.Layout.producer already exists.
			// Bind this accordionItem consumer and the producer.
			accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordionItem, screenConsumer, spawned)
			messenger = unSpawnedPackageMessenger
			producer := unSpawnedPackageScreen.Layout.Producer()
			accordionItemContentConsumer.Bind(producer)
			return
		}
	}

	var packageScreen *_misc_.Miscellaneous
	var packageMessenger *_txrx_.Messenger
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newAccordionItemContentConsumer: %w", err)
			initErr = err
		} else {
			if !spawned {
				unSpawnedPackageMessenger = packageMessenger
				unSpawnedPackageScreen = packageScreen
			}
			initErr = nil
		}
	}()

	accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordionItem, screenConsumer, spawned)
	if packageScreen, packageMessenger, err = newScreenMessenger(ctx, ctxCancel, app, w, spawned, accordionItemContentConsumer); err != nil {
		return
	}

	// This screen only show 1 of it's panels at a time.
	// Show the default panel.
	packageScreen.Panelers.DefaultPanel.Show()

	return
}

func newScreenMessenger(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, spawned bool, consumer _types_.ContentConsumer) (screen *_misc_.Miscellaneous, messenger *_txrx_.Messenger, err error) {
	// Build the content & producer.
	producer := _producer_.NewContentProducer(spawned, consumer)
	producer.Bind(consumer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(producer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, w, layout); err != nil {
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
{{- range $panelName := .LocalPanelNames }}
	{{ call $DOT.Funcs.DeCap $panelName }}Panel.SetMessenger(messenger)
{{- end }}

	return
}
`
)
