package misc

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type ScreenComponentsTemplateData struct {
	PackageName      string
	PanelNames       []string // All panels are local.
	DefaultPanelName string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	ScreenComponentsFileName = "screenComponents.go"

	ScreenComponentsTemplate = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// ScreenComponents is this screen, it's panels and messenger.
// This screen has {{ len .PanelNames }} panels.
// The default panel is {{ index .PanelNames 0 }}.
type ScreenComponents struct {
	CTX       context.Context
	CTXCancel context.CancelFunc
	APP       fyne.App
	Window    fyne.Window

	CanvasObjectProvider _types_.CanvasObjectProvider
	Panels               *Panels
	Messenger            *Messenger
}

// NewScreenComponents returns the screen's canvas object provider and the error. 
// If needed, it constructs this screen.
// * It constructs each panel in this screen.
// * It constructs the messenger.
// * It uses the default panel contents to create the screen content.
func NewScreenComponents(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (components *ScreenComponents, err error) {
	components = &ScreenComponents{
		CTX:                  ctx,
		CTXCancel:            ctxCancel,
		APP:                  app,
		Window:               w,
		CanvasObjectProvider: _types_.NewScreenCanvasManager(),
	}

	// Panels.
	if components.Panels, err = NewPanels(packageScreen); err != nil {
		return
	}
	components.canvasObjectProvider.UpdateCanvasObject(newScreen.Panels.{{ .DefaultPanelName }}.Content())

	// Messenger.
	components.Messenger = newMessageHandler(newScreen)
	return
}


// UnSpawn is an implementation of _types_.UnSpawner.
func (components *ScreenComponents) UnSpawn() {
	components.Messenger.StopReceiving()
}

`
)
