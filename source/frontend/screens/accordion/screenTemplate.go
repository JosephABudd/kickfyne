package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type screenTemplateData struct {
	PackageName  string
	ImportPrefix string
}

const (
	screenFileName = _utils_.ScreenFileName

	screenTemplate = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

var screen *_misc_.ScreenComponents
var initErr error

// New returns the screen's canvas object provider and the error. 
// If needed, it constructs this screen.
// * It constructs each panel in this screen.
// * It constructs the messenger.
// * It uses the default panel contents to create the screen content.
// If this screen has already been constructed then it uses the already constructed screen.
func New(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (canvasObjectProvider _types_.CanvasObjectProvider, err error) {

	if screen != nil {
		canvasObjectProvider = screen.canvasObjectProvider
		err = initErr
		return
	}

	defer func() {
		if err == nil {
			canvasObjectProvider = screen.canvasObjectProvider
			initErr = err
		}
		if err != nil {
			screen = nil
			err = fmt.Errorf("{{ .PackageName }}.New: %w", err)
			initErr = err
		}
	}()

	screen, err = _misc_.NewScreenComponents(ctx, ctxCancel, app, w)
	return
}

// Spawn returns the screen's canvas object provider, and messenger and the error.
// Messenger must implement _types_.MessageSpawner
// * It constructs this screen.
// * It constructs each panel in this screen.
// * It constructs the messenger.
// * It uses the layout contents to create the screen content.
func Spawn(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (canvasObjectProvider _types_.CanvasObjectProvider, messenger _types_.MessageSpawner, unspawner _types_.UnSpawner, err error) {

	var spawn_screen *_misc_.ScreenComponents
	defer func() {
		if err == nil {
			canvasObjectProvider = spawn_screen.canvasObjectProvider
			messenger = spawn_screen.Messenger
			unspawner = spawn_screen
		}
		if err != nil {
			spawn_screen = nil
			err = fmt.Errorf("{{ .PackageName }}.Spawn: %w", err)
		}
	}()

	spawn_screen, err = _misc_.NewScreenComponents(ctx, ctxCancel, app, w)
	return
}

`
)
