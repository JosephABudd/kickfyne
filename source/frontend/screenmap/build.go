package screenmap

import (
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework builds frontend/screenmap/screenMap.go
// Call it after a screen is added or removed.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	fPath := filepath.Join(folderPaths.FrontendScreenMap, screenMapFileName)
	data := &screenMapTemplateData{
		ImportPrefix: importPrefix,
	}
	err = _utils_.ProcessTemplate(screenMapFileName, fPath, screenMapTemplate, data)
	return
}
