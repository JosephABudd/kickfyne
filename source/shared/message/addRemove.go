package message

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

func AddMessage(
	messageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.Add: %w", err)
		}
	}()

	fname := _utils_.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.SharedMessage, fname)
	data := messageTemplateData{
		MessageName: messageName,
	}
	err = _utils_.ProcessTemplate(fname, oPath, messageTemplate, data)
	return
}

func RemoveMessage(
	messageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.RemoveMessage: %w", err)
		}
	}()

	fName := _utils_.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.SharedMessage, fName)
	err = os.Remove(oPath)
	return
}
