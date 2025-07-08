package txrx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_accordion_ "github.com/JosephABudd/kickfyne/source/backend/txrx/accordion"
	_tabbar_ "github.com/JosephABudd/kickfyne/source/backend/txrx/tabbar"
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

func AddAccordion(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var oPath string
	var data interface{}

	funcs := _utils_.GetFuncs()
	// Make the shared/message/«packageName»/ folder.
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if err = os.Mkdir(folderPath, _utils_.DMode); err != nil {
		return
	}
	accordionMessageNameFileName := _utils_.AccordionMessageNameFileName(packageName)
	accordionMessageNameStructName := _utils_.AccordionMessageNameStructName(packageName)
	for messageName, fileName := range accordionMessageNameFileName {
		oPath = filepath.Join(folderPath, fileName)
		switch messageName {
		case _utils_.AccordionMessageB2FAddAccordion:
			data = _accordion_.B2FAddAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.B2FAddAccordionTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FAddAccordionItem:
			data = _accordion_.B2FAddAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.B2FAddAccordionItemTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FRemoveAccordion:
			data = _accordion_.B2FRemoveAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.B2FRemoveAccordionTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FRemoveAccordionItem:
			data = _accordion_.B2FRemoveAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.B2FRemoveAccordionItemTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BAddAccordion:
			data = _accordion_.F2BAddAccordionTemplateData{
				PackageName:                          packageName,
				MessageStructName:                    accordionMessageNameStructName[messageName],
				B2FAddAccordionItemMessageStructName: accordionMessageNameStructName[_utils_.AccordionMessageB2FAddAccordionItem],
				ImportPrefix:                         importPrefix,
				Funcs:                                funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.F2BAddAccordionTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BAddAccordionItem:
			data = _accordion_.F2BAddAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.F2BAddAccordionItemTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BRemoveAccordion:
			data = _accordion_.F2BRemoveAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.F2BRemoveAccordionTemplate, data); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BRemoveAccordionItem:
			data = _accordion_.F2BRemoveAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: accordionMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _accordion_.F2BRemoveAccordionItemTemplate, data); err != nil {
				return
			}
		}
	}

	// txrx/txrx.go
	err = RebuildTXRX(importPrefix, folderPaths, funcs)

	return
}

func RemoveAccordion(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	var removeErr error
	var rebuildErr error
	defer func() {
		err = fmt.Errorf(
			"backend.RemoveAccordion:%w",
			errors.Join(removeErr, rebuildErr),
		)
	}()
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if removeErr = os.RemoveAll(folderPath); removeErr != nil {
		if os.IsNotExist(removeErr) {
			// Folder does not exist.
			removeErr = nil
		}
	}
	funcs := _utils_.GetFuncs()
	rebuildErr = RebuildTXRX(importPrefix, folderPaths, funcs)
	return
}

func AddAppTab(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var oPath string
	var data interface{}

	funcs := _utils_.GetFuncs()
	// Make the shared/message/«packageName»/ folder.
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if err = os.Mkdir(folderPath, _utils_.DMode); err != nil {
		return
	}
	appTabMessageNameFileName := _utils_.AppTabMessageNameFileName(packageName)
	appTabMessageNameStructName := _utils_.AppTabMessageNameStructName(packageName)
	for messageName, fileName := range appTabMessageNameFileName {
		oPath = filepath.Join(folderPath, fileName)
		switch messageName {
		case _utils_.AppTabMessageB2FAddTabbar:
			data = _tabbar_.B2FAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FAddTab:
			data = _tabbar_.B2FAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FRemoveTabbar:
			data = _tabbar_.B2FRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FRemoveTab:
			data = _tabbar_.B2FRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FRemoveTabTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BAddTabbar:
			data = _tabbar_.F2BAddTabbarTemplateData{
				PackageName:                packageName,
				MessageStructName:          appTabMessageNameStructName[messageName],
				B2FAddTabMessageStructName: appTabMessageNameStructName[_utils_.AppTabMessageB2FAddTab],
				ImportPrefix:               importPrefix,
				Funcs:                      funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BAddTab:
			data = _tabbar_.F2BAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BRemoveTabbar:
			data = _tabbar_.F2BRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BRemoveTab:
			data = _tabbar_.F2BRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: appTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BRemoveTabTemplate, data); err != nil {
				return
			}
		}
	}

	// txrx/txrx.go
	err = RebuildTXRX(importPrefix, folderPaths, funcs)

	return
}

func RemoveAppTab(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	var removeErr error
	var rebuildErr error
	defer func() {
		err = fmt.Errorf(
			"backend.RemoveAppTab:%w",
			errors.Join(removeErr, rebuildErr),
		)
	}()
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if removeErr = os.RemoveAll(folderPath); removeErr != nil {
		if os.IsNotExist(removeErr) {
			// Folder does not exist.
			removeErr = nil
		}
	}
	funcs := _utils_.GetFuncs()
	rebuildErr = RebuildTXRX(importPrefix, folderPaths, funcs)
	return
}

func AddDocTab(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var oPath string
	var data interface{}

	funcs := _utils_.GetFuncs()
	// Make the shared/message/«packageName»/ folder.
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if err = os.Mkdir(folderPath, _utils_.DMode); err != nil {
		return
	}
	docTabMessageNameFileName := _utils_.DocTabMessageNameFileName(packageName)
	docTabMessageNameStructName := _utils_.DocTabMessageNameStructName(packageName)
	for messageName, fileName := range docTabMessageNameFileName {
		oPath = filepath.Join(folderPath, fileName)
		switch messageName {
		case _utils_.DocTabMessageB2FAddTabbar:
			data = _tabbar_.B2FAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FAddTab:
			data = _tabbar_.B2FAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTabbar:
			data = _tabbar_.B2FRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTab:
			data = _tabbar_.B2FRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.B2FRemoveTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTabbar:
			data = _tabbar_.F2BAddTabbarTemplateData{
				PackageName:                packageName,
				MessageStructName:          docTabMessageNameStructName[messageName],
				B2FAddTabMessageStructName: docTabMessageNameStructName[_utils_.DocTabMessageB2FAddTab],
				ImportPrefix:               importPrefix,
				Funcs:                      funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTab:
			data = _tabbar_.F2BAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTabbar:
			data = _tabbar_.F2BRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTab:
			data = _tabbar_.F2BRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: docTabMessageNameStructName[messageName],
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(fileName, oPath, _tabbar_.F2BRemoveTabTemplate, data); err != nil {
				return
			}
		}
	}

	// txrx/txrx.go
	err = RebuildTXRX(importPrefix, folderPaths, funcs)

	return
}

func RemoveDocTab(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	var removeErr error
	var rebuildErr error
	defer func() {
		err = fmt.Errorf(
			"backend.RemoveDocTab:%w",
			errors.Join(removeErr, rebuildErr),
		)
	}()
	folderPath := filepath.Join(folderPaths.BackendTXRX, packageName)
	if removeErr = os.RemoveAll(folderPath); removeErr != nil {
		if os.IsNotExist(removeErr) {
			// Folder does not exist.
			removeErr = nil
		}
	}
	funcs := _utils_.GetFuncs()
	rebuildErr = RebuildTXRX(importPrefix, folderPaths, funcs)
	return
}
