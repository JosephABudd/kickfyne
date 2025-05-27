package txrx

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// AddMessageHandler adds a message handler to the back-end txrx folder.
func AddMessageHandler(
	messageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.AddMessageHandler: %w", err)
		}
	}()

	// This is a new unique message name.
	fName := _utils_.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.BackendTXRX, fName)
	data := handlerTemplateData{
		ImportPrefix: importPrefix,
		MessageName:  messageName,
		Funcs:        _utils_.GetFuncs(),
	}
	if err = _utils_.ProcessTemplate(fName, oPath, handlerTemplate, data); err != nil {
		return
	}

	return
}

// RemoveMessageHandler removes a message handler from the back-end txrx folder.
func RemoveMessageHandler(
	messageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.RemoveMessageHandler: %w", err)
		}
	}()

	fName := _utils_.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.BackendTXRX, fName)
	err = os.Remove(oPath)
	return
}
