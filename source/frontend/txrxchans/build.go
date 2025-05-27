package txrxchans

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/txrx/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.CreateFramework: %w", err)
		}
	}()

	// frontend/txrx/listen.go
	oPath := filepath.Join(folderPaths.FrontendTXRXChans, txrxFileName)
	data := txrxTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(txrxFileName, oPath, txrxTemplate, data); err != nil {
		return
	}

	return
}
