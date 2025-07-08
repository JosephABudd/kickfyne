package thread

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the shared/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("thread.CreateFramework: %w", err)
		}
	}()

	oPath := filepath.Join(folderPaths.SharedThread, fileName)
	err = _utils_.WriteFile(oPath, []byte(template))
	return
}
