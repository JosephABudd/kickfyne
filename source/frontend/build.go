package frontend

import (
	"fmt"
	"path/filepath"

	_mainmenu_ "github.com/JosephABudd/kickfyne/source/frontend/mainmenu"
	_screenmap_ "github.com/JosephABudd/kickfyne/source/frontend/screenmap"
	_screens_ "github.com/JosephABudd/kickfyne/source/frontend/screens"
	_txrxchans_ "github.com/JosephABudd/kickfyne/source/frontend/txrxchans"
	_types_ "github.com/JosephABudd/kickfyne/source/frontend/types"
	_widget_ "github.com/JosephABudd/kickfyne/source/frontend/widget"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.CreateFramework: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// frontend/frontend.go
	oPath = filepath.Join(folderPaths.Frontend, frontendFileName)
	data = frontendTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(frontendFileName, oPath, frontendTemplate, data); err != nil {
		return
	}

	// gui/mainmenu/ package
	if err = _mainmenu_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/txrx/ package
	if err = _txrxchans_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/types/
	if err = _types_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/widget/
	err = _widget_.CreateFramework(folderPaths)

	// Add the HelloWorld screen.
	if err = _screens_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}
	// Rebuild frontend/screenmap/screenmap.go
	err = _screenmap_.Rebuild(importPrefix, folderPaths)

	return

}
