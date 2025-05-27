package txrx

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's backend/txrx/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.CreateFramework: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// txrx/doc.go
	oPath = filepath.Join(folderPaths.BackendTXRX, docFileName)
	if err = _utils_.ProcessTemplate(docFileName, oPath, docTemplate, nil); err != nil {
		return
	}

	// txrx/txrx.go
	oPath = filepath.Join(folderPaths.BackendTXRX, tXRXFileName)
	data = tXRXTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(tXRXFileName, oPath, tXRXTemplate, data); err != nil {
		return
	}

	// txrx/Init.go
	fname := _utils_.MessageFileName(_utils_.InitMessageName)
	oPath = filepath.Join(folderPaths.BackendTXRX, fname)
	data = initRXTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(fname, oPath, initRXTemplate, data); err != nil {
		return
	}

	return
}
