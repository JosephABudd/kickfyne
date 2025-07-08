package message

import (
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the shared/message/ files.
func CreateFramework(
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var oPath string

	// message/chans.go
	oPath = filepath.Join(folderPaths.SharedMessage, chansFileName)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = _utils_.ProcessTemplate(chansFileName, oPath, chansTemplate, nil); err != nil {
		return
	}

	// message/init.go
	oPath = filepath.Join(folderPaths.SharedMessage, initFileName)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = _utils_.ProcessTemplate(initFileName, oPath, initTemplate, nil); err != nil {
		return
	}

	return
}
