package txrx

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"

	_api_ "github.com/JosephABudd/kickfyne/source/backend/txrx/api"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

func RebuildTXRX(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
	funcs _utils_.Funcs,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.RebuildTXRX: %w", err)
		}
	}()

	// txrx/txrx.go
	oPath := filepath.Join(folderPaths.BackendTXRX, tXRXFileName)
	var txrxFolderNames []string
	if txrxFolderNames, err = _utils_.BackendTXRXFolderNames(folderPaths); err != nil {
		return
	}
	log.Printf("%d txrxFolderNames", len(txrxFolderNames))
	sort.Strings(txrxFolderNames)
	data := tXRXTemplateData{
		ImportPrefix:    importPrefix,
		TXRXFolderNames: txrxFolderNames,
		Funcs:           funcs,
	}
	err = _utils_.ProcessTemplate(tXRXFileName, oPath, tXRXTemplate, data)
	return
}

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

	// txrx/Init.go
	fname := _utils_.MessageFileName(_utils_.InitMessageName)
	oPath = filepath.Join(folderPaths.BackendTXRX, fname)
	data = initRXTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(fname, oPath, initRXTemplate, data); err != nil {
		return
	}

	// txrx/txrx.go
	funcs := _utils_.GetFuncs()
	if err = RebuildTXRX(importPrefix, folderPaths, funcs); err != nil {
		return
	}

	// txrx/api/api.go

	oPath = filepath.Join(folderPaths.BackendTXRXAPI, _utils_.APIFileName)
	var txrxFolderNames []string
	if txrxFolderNames, err = _utils_.BackendTXRXFolderNames(folderPaths); err != nil {
		return
	}
	data = _api_.APITemplateData{
		ImportPrefix:    importPrefix,
		TXRXFolderNames: txrxFolderNames,
		Funcs:           funcs,
	}
	if err = _utils_.ProcessTemplate(_utils_.APIFileName, oPath, _api_.APITemplate, data); err != nil {
		return
	}
	return
}
