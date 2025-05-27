package shared

import (
	"fmt"
	"path/filepath"

	_message_ "github.com/JosephABudd/kickfyne/source/shared/message"
	_metadata_ "github.com/JosephABudd/kickfyne/source/shared/metadata"
	_paths_ "github.com/JosephABudd/kickfyne/source/shared/paths"
	_store_ "github.com/JosephABudd/kickfyne/source/shared/store"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

const (
	FolderName = "shared"
)

// CreateFramework creates the shared/ files.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shared.CreateFramework: %w", err)
		}
	}()

	// shared/shared.go
	path := filepath.Join(folderPaths.Shared, sharedFileName)
	data := sharedTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(sharedFileName, path, sharedTemplate, data); err != nil {
		return
	}

	// shared/message/
	if err = _message_.CreateFramework(folderPaths); err != nil {
		return
	}

	// shared/metadata/
	if err = _metadata_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// shared/paths/
	if err = _paths_.CreateFramework(appName, folderPaths); err != nil {
		return
	}

	// shared/store/
	if err = _store_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
