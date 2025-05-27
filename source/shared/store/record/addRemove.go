package record

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// AddRecord adds a record file to shared/store/record/.
func AddRecord(
	recordName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.AddRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreRecord, fName)
	data := templateData{
		RecordName: recordName,
		Funcs:      _utils_.GetFuncs(),
	}
	err = _utils_.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record file from shared/store/record/.
func RemoveRecord(
	recordName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.RemoveRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreRecord, fName)
	err = os.Remove(oPath)
	return
}
