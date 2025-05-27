package store

import (
	"fmt"

	_record_ "github.com/JosephABudd/kickfyne/source/shared/store/record"
	_storer_ "github.com/JosephABudd/kickfyne/source/shared/store/storer"
	_storing_ "github.com/JosephABudd/kickfyne/source/shared/store/storing"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// AddRecord add the files for the new record and then rebuilds stores.go.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.AddRecord: %w", err)
		}
	}()

	if err = _record_.AddRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = _storer_.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	if err = _storing_.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	err = rebuildStoresGo(importPrefix, folderPaths)
	return
}

// RemoveRecord add the files for the new record and then rebuilds stores.go.
func RemoveRecord(
	recordName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.RemoveRecord: %w", err)
		}
	}()

	if err = _record_.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = _storer_.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = _storing_.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	err = rebuildStoresGo(importPrefix, folderPaths)
	return
}
