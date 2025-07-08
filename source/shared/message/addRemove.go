package message

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_accordion_ "github.com/JosephABudd/kickfyne/source/shared/message/accordion"
	_doctabs_ "github.com/JosephABudd/kickfyne/source/shared/message/doctabs"
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
		MessageStructName: messageName,
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

func AddAccordion(
	packageName string,
	allPanelNames []string,
	localPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	messageFolderPath := _utils_.AccordionSharedMessageFolderPath(packageName, folderPaths)
	// These doctabs messages have their own folder.
	if err = os.Mkdir(messageFolderPath, _utils_.DMode); err != nil {
		return
	}
	messageNameFileName := _utils_.AccordionMessageNameFileName(packageName)
	messageNameStructName := _utils_.AccordionMessageNameStructName(packageName)
	var templateData any
	var template string
	var funcs = _utils_.GetFuncs()
	for messageName, messageFileName := range messageNameFileName {
		messagePath := filepath.Join(messageFolderPath, messageFileName)
		messageStructName := messageNameStructName[messageName]
		messageFileName := messageNameFileName[messageName]
		switch messageName {
		case _utils_.AccordionMessageB2FAddAccordion:
			templateData = &_accordion_.B2FAddAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _accordion_.B2FAddAccordionTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FAddAccordionItem:
			templateData = &_accordion_.B2FAddAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _accordion_.B2FAddAccordionItemTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FRemoveAccordion:
			templateData = &_accordion_.B2FRemoveAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _accordion_.B2FRemoveAccordionTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageB2FRemoveAccordionItem:
			templateData = &_accordion_.B2FRemoveAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _accordion_.B2FRemoveAccordionItemTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BAddAccordion:
			templateData = &_accordion_.F2BAddAccordionTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _accordion_.F2BAddAccordionTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BAddAccordionItem:
			templateData = &_accordion_.F2BAddAccordionItemTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _accordion_.F2BAddAccordionItemTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BRemoveAccordion:
			templateData = &_accordion_.F2BRemoveAccordionTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _accordion_.F2BRemoveAccordionTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AccordionMessageF2BRemoveAccordionItem:
			templateData = &_accordion_.F2BRemoveAccordionItemTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _accordion_.F2BRemoveAccordionItemTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		}
	}
	// Panel state for this package.
	templateData = &_accordion_.PaneslStateTemplateData{
		PackageName:     packageName,
		LocalPanelNames: localPanelNames,
		ImportPrefix:    importPrefix,
		Funcs:           funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.PanelsStateFileName,
		filepath.Join(messageFolderPath, _utils_.PanelsStateFileName),
		_accordion_.PaneslStateTemplate,
		templateData,
	); err != nil {
		return
	}
	// AccordionItem.
	tabsPath := filepath.Join(messageFolderPath, _utils_.IDsFileName)
	tabsData := _accordion_.IDsTemplateData{
		PackageName:   packageName,
		AllPanelNames: allPanelNames,
		Funcs:         funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.IDsFileName,
		tabsPath,
		_accordion_.IDsTemplate,
		tabsData,
	); err != nil {
		return
	}

	return
}

func RemoveAccordion(
	packageName string,
	folderPaths *_utils_.FolderPaths,
) (errs error) {
	for _, messagePath := range _utils_.AccordionMessageNameFileName(packageName) {
		if err := os.Remove(messagePath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	return
}

func AddAppTabs(
	packageName string,
	allPanelNames []string,
	localPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	messageFolderPath := _utils_.AppTabSharedMessageFolderPath(packageName, folderPaths)
	// These doctabs messages have their own folder.
	if err = os.Mkdir(messageFolderPath, _utils_.DMode); err != nil {
		return
	}
	messageNameFileName := _utils_.AppTabMessageNameFileName(packageName)
	messageNameStructName := _utils_.AppTabMessageNameStructName(packageName)
	var templateData any
	var template string
	var funcs = _utils_.GetFuncs()
	for messageName, messageFileName := range messageNameFileName {
		messagePath := filepath.Join(messageFolderPath, messageFileName)
		messageStructName := messageNameStructName[messageName]
		messageFileName := messageNameFileName[messageName]
		switch messageName {
		case _utils_.AppTabMessageB2FAddTabbar:
			templateData = &_doctabs_.B2FAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _doctabs_.B2FAddTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FAddTab:
			templateData = &_doctabs_.B2FAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.B2FAddTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FRemoveTabbar:
			templateData = &_doctabs_.B2FRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.B2FRemoveTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageB2FRemoveTab:
			templateData = &_doctabs_.B2FRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _doctabs_.B2FRemoveTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BAddTabbar:
			templateData = &_doctabs_.F2BAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BAddTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BAddTab:
			templateData = &_doctabs_.F2BAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BAddTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BRemoveTabbar:
			templateData = &_doctabs_.F2BRemoveTabbarTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BRemoveTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.AppTabMessageF2BRemoveTab:
			templateData = &_doctabs_.F2BRemoveTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BRemoveTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		}
	}
	// Panel state for this package.
	templateData = &_doctabs_.PaneslStateTemplateData{
		PackageName:     packageName,
		LocalPanelNames: localPanelNames,
		ImportPrefix:    importPrefix,
		Funcs:           funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.PanelsStateFileName,
		filepath.Join(messageFolderPath, _utils_.PanelsStateFileName),
		_doctabs_.PaneslStateTemplate,
		templateData,
	); err != nil {
		return
	}
	// Tabs.
	tabsPath := filepath.Join(messageFolderPath, _utils_.IDsFileName)
	tabsData := _doctabs_.IDsTemplateData{
		PackageName:   packageName,
		AllPanelNames: allPanelNames,
		Funcs:         funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.IDsFileName,
		tabsPath,
		_doctabs_.TabsIDsTemplate,
		tabsData,
	); err != nil {
		return
	}

	return
}

func RemoveAppTabs(
	packageName string,
	folderPaths *_utils_.FolderPaths,
) (errs error) {
	for _, messagePath := range _utils_.AppTabMessageNameFileName(packageName) {
		// SpawnedTab message.
		if err := os.Remove(messagePath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	return
}

func AddDocTabs(
	packageName string,
	allPanelNames []string,
	localPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	messageFolderPath := _utils_.DocTabSharedMessageFolderPath(packageName, folderPaths)
	// These doctabs messages have their own folder.
	if err = os.Mkdir(messageFolderPath, _utils_.DMode); err != nil {
		return
	}
	messageNameFileName := _utils_.DocTabMessageNameFileName(packageName)
	messageNameStructName := _utils_.DocTabMessageNameStructName(packageName)
	var templateData any
	var template string
	var funcs = _utils_.GetFuncs()
	for messageName, messageFileName := range messageNameFileName {
		messagePath := filepath.Join(messageFolderPath, messageFileName)
		messageStructName := messageNameStructName[messageName]
		messageFileName := messageNameFileName[messageName]
		switch messageName {
		case _utils_.DocTabMessageB2FAddTabbar:
			templateData = &_doctabs_.B2FAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _doctabs_.B2FAddTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FAddTab:
			templateData = &_doctabs_.B2FAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.B2FAddTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTabbar:
			templateData = &_doctabs_.B2FRemoveTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.B2FRemoveTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTab:
			templateData = &_doctabs_.B2FRemoveTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template = _doctabs_.B2FRemoveTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTabbar:
			templateData = &_doctabs_.F2BAddTabbarTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BAddTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTab:
			templateData = &_doctabs_.F2BAddTabTemplateData{
				PackageName:       packageName,
				MessageStructName: messageStructName,
				ImportPrefix:      importPrefix,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BAddTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTabbar:
			templateData = &_doctabs_.F2BRemoveTabbarTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BRemoveTabbarTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTab:
			templateData = &_doctabs_.F2BRemoveTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: messageStructName,
				Funcs:             funcs,
			}
			template := _doctabs_.F2BRemoveTabTemplate
			if err = _utils_.ProcessTemplate(
				messageFileName,
				messagePath,
				template,
				templateData,
			); err != nil {
				return
			}
		}
	}
	// Panel state for this package.
	templateData = &_doctabs_.PaneslStateTemplateData{
		PackageName:     packageName,
		LocalPanelNames: localPanelNames,
		ImportPrefix:    importPrefix,
		Funcs:           funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.PanelsStateFileName,
		filepath.Join(messageFolderPath, _utils_.PanelsStateFileName),
		_doctabs_.PaneslStateTemplate,
		templateData,
	); err != nil {
		return
	}
	// Tabs.
	tabsPath := filepath.Join(messageFolderPath, _utils_.IDsFileName)
	tabsData := _doctabs_.IDsTemplateData{
		PackageName:   packageName,
		AllPanelNames: allPanelNames,
		Funcs:         funcs,
	}
	if err = _utils_.ProcessTemplate(
		_utils_.IDsFileName,
		tabsPath,
		_doctabs_.TabsIDsTemplate,
		tabsData,
	); err != nil {
		return
	}

	return
}

func RemoveDocTabs(
	packageName string,
	folderPaths *_utils_.FolderPaths,
) (errs error) {
	for _, messagePath := range _utils_.DocTabMessageNameFileName(packageName) {
		// SpawnedTab message.
		if err := os.Remove(messagePath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	return
}
