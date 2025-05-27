package backend

import (
	"fmt"
	"path/filepath"

	_txrx_ "github.com/JosephABudd/kickfyne/source/backend/txrx"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

const (
	FolderName = "backend"
)

// CreateFramework creates the backend/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backend.CreateFramework: %w", err)
		}
	}()

	// backend/backend.go
	oPath := filepath.Join(folderPaths.Backend, fileName)
	data := templateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(fileName, oPath, template, data); err != nil {
		return
	}

	// backend/txrx/
	if err = _txrx_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
