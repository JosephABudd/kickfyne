package storer

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// AddRecord adds a record storer file to shared/store/storer/.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storer.AddRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	storingFilePath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	data := templateData{
		RecordName:      recordName,
		ImportPrefix:    importPrefix,
		StoringFilePath: storingFilePath,
	}
	err = _utils_.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record storer file from shared/store/storer/.
func RemoveRecord(
	recordName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storer.RemoveRecord: %w", err)
		}
	}()

	fName := _utils_.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	err = os.Remove(oPath)
	return
}
