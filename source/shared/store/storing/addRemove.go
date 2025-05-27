package storing

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// AddRecord adds a record storing file to shared/store/storing/.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storing.AddRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	storerFilePath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	data := templateData{
		RecordName:     recordName,
		ImportPrefix:   importPrefix,
		StorerFilePath: storerFilePath,
	}
	err = _utils_.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record storing file from shared/store/storing/.
func RemoveRecord(
	recordName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storing.RemoveRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	err = os.Remove(oPath)
	return
}
