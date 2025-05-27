package mainmenu

type mainMenuTemplateData struct {
	ImportPrefix string
}

const (
	mainMenuFileName = "mainmenu.go"

	mainMenuTemplate = `package mainmenu

import (
	"context"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"

	_ScreenMap_ "example.com/okp/frontend/screenmap"
	_types_ "example.com/okp/frontend/types"
)

// Init builds the main menu and adds it to the app.
func Init(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.Init: %w", err)
		}
	}()

	// Build the sub menu items.
	var items []*fyne.MenuItem
	if items, err = subMenuItems(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}
	// Build the sub menu.
	subMenu := fyne.NewMenu(app.Metadata().Name, items...)
	// Build the main menu.
	mainmenu := fyne.NewMainMenu(subMenu)
	window.SetMainMenu(mainmenu)
	return
}

func subMenuItems(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (items []*fyne.MenuItem, err error) {
	items = make([]*fyne.MenuItem, 0, 5)

	var screenNames []string = strings.Split(app.Metadata().Custom["MainMenu"], " ")
	// consumer receives the screen's content from the screen's producer and gives it to the window.
	for i, screenName := range screenNames {
		var api *_ScreenMap_.API
		if api = _ScreenMap_.Map[screenName]; api == nil {
			fmt.Printf("FyneApp.toml Development.MainMenu: %q is not a real screen name.\n", screenName)
			continue
		}
		var consumer _types_.ContentConsumer
		if consumer, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, window); err != nil {
			return
		}
		item := fyne.NewMenuItem(screenName, consumer.Show)
		items = append(items, item)
		if i == 0 {
			// The first screen is the opening screen.
			consumer.Show()
		}
	}

	return
}
`
)
