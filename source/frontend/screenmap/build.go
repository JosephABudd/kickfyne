package screenmap

import (
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// Rebuild rebuids frontend/screenmap/screenMap.go
// Call it after a screen is added or removed.
func Rebuild(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	var screenPackageNames []string
	if screenPackageNames, err = _utils_.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	fPath := filepath.Join(folderPaths.FrontendScreenMap, screenMapFileName)
	data := &screenMapTemplateData{
		ImportPrefix: importPrefix,
		ScreenNames:  screenPackageNames,
	}
	err = _utils_.ProcessTemplate(screenMapFileName, fPath, screenMapTemplate, data)
	return
}
