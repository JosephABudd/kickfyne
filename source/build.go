package source

import (
	"fmt"
	"os"

	_backend_ "github.com/JosephABudd/kickfyne/source/backend"
	_frontend_ "github.com/JosephABudd/kickfyne/source/frontend"
	_root_ "github.com/JosephABudd/kickfyne/source/root"
	_shared_ "github.com/JosephABudd/kickfyne/source/shared"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

func HasAppFolder(currentWP, appName string) (hasAppFolder bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.HasAppFolder: %w", err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(currentWP); err != nil {
		return
	}
	var dirEntry os.DirEntry
	for _, dirEntry = range dirEntrys {
		if dirEntry.IsDir() {
			dName := dirEntry.Name()
			if hasAppFolder = dName == appName; hasAppFolder {
				return
			}
		}
	}
	return
}

// CreateFramework builds the framework in an appName folder in this parent folder.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.CreateFramework: %w", err)
		}
	}()

	// App folder.
	if err = _root_.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Shared
	if err = _shared_.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Backend
	if err = _backend_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// Frontend
	if err = _frontend_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
