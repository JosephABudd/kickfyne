package simple

import (
	"fmt"
	"os"
	"path/filepath"

	_layout_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/layout"
	_misc_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/misc"
	_panelers_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/panelers"
	_panels_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/panels"
	_panel_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/panels/panel"
	_producer_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/producer"
	_tabs_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/tabs"
	_txrx_ "github.com/JosephABudd/kickfyne/source/frontend/screens/doctabs/txrx"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// Build builds a type Simple screen package.
func Build(
	packageName string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("doctabs.Build: %w", err)
		}
	}()

	defaultPanelName := localPanelNames[0]

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)
	if err = os.Mkdir(packagePath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/misc
	packageMiscPath := filepath.Join(packagePath, _utils_.FolderNameMisc)
	if err = os.Mkdir(packageMiscPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/panelers
	packagePanelersPath := filepath.Join(packagePath, _utils_.FolderNamePanelers)
	if err = os.Mkdir(packagePanelersPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/producer
	packageProducerPath := filepath.Join(packagePath, _utils_.FolderNameProducer)
	if err = os.Mkdir(packageProducerPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)
	if err = os.Mkdir(packagePanelsPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/txrx
	packageTXRXPath := filepath.Join(packagePath, _utils_.FolderNameTXRX)
	if err = os.Mkdir(packageTXRXPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/layout
	packageLayoutPath := filepath.Join(packagePath, _utils_.FolderNameLayout)
	if err = os.Mkdir(packageLayoutPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/tabs
	packageTabsPath := filepath.Join(packagePath, _utils_.FolderNameTabs)
	if err = os.Mkdir(packageTabsPath, _utils_.DMode); err != nil {
		return
	}

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	successMessage := docTemplateSuccessMessage(packageName, localPanelNames, folderPaths)
	data = &docTemplateData{
		PackageName: packageName,
		PackageDoc:  packageDoc,
		Files:       successMessage,
		Funcs:       funcs,
	}
	if err = _utils_.ProcessTemplate(docFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/api.go
	messageNameStructName := _utils_.DocTabMessageNameStructName(packageName)
	fPath = filepath.Join(packagePath, aPIFileName)
	data = &aPITemplateData{
		PackageName:             packageName,
		AllPanelNames:           allPanelNames,
		DefaultPanelName:        defaultPanelName,
		F2BAddTabbarMessageName: messageNameStructName[_utils_.DocTabMessageF2BAddTabbar],
		ImportPrefix:            importPrefix,
		Funcs:                   funcs,
	}
	if err = _utils_.ProcessTemplate(aPIFileName, fPath, aPITemplate, data); err != nil {
		return
	}

	// layout/ folder.

	// frontend/screens/«screen-package-name»/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	// tabs/ folder.

	// frontend/screens/«screen-package-name»/tabs/tabs.go
	fPath = filepath.Join(packageTabsPath, _tabs_.FileName)
	data = &_tabs_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(_tabs_.FileName, fPath, _tabs_.Template, data); err != nil {
		return
	}

	// misc/ folder.

	// frontend/screens/«screen-package-name»/misc/miscellaneous.go
	fPath = filepath.Join(packageMiscPath, _misc_.MiscellaneousFileName)
	data = &_misc_.MiscellaneousTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_misc_.MiscellaneousFileName, fPath, _misc_.MiscellaneousTemplate, data); err != nil {
		return
	}

	// txrx folder.

	// frontend/screens/«screen-package-name»/txrx/messenger.go
	docTabMessageNameStructName := _utils_.DocTabMessageNameStructName(packageName)
	fPath = filepath.Join(packageTXRXPath, _txrx_.MessengerFileName)
	data = &_txrx_.MessengerTemplateData{
		PackageName:                  packageName,
		ImportPrefix:                 importPrefix,
		LocalPanelNames:              localPanelNames,
		Funcs:                        funcs,
		DocTabMessageB2FAddTabbar:    docTabMessageNameStructName[_utils_.DocTabMessageB2FAddTabbar],
		DocTabMessageB2FAddTab:       docTabMessageNameStructName[_utils_.DocTabMessageB2FAddTab],
		DocTabMessageB2FRemoveTabbar: docTabMessageNameStructName[_utils_.DocTabMessageB2FRemoveTabbar],
		DocTabMessageB2FRemoveTab:    docTabMessageNameStructName[_utils_.DocTabMessageB2FRemoveTab],
		DocTabMessageF2BAddTabbar:    docTabMessageNameStructName[_utils_.DocTabMessageF2BAddTabbar],
		DocTabMessageF2BAddTab:       docTabMessageNameStructName[_utils_.DocTabMessageF2BAddTab],
		DocTabMessageF2BRemoveTabbar: docTabMessageNameStructName[_utils_.DocTabMessageF2BRemoveTabbar],
		DocTabMessageF2BRemoveTab:    docTabMessageNameStructName[_utils_.DocTabMessageF2BRemoveTab],
	}
	if err = _utils_.ProcessTemplate(_txrx_.MessengerFileName, fPath, _txrx_.MessengerTemplate, data); err != nil {
		return
	}
	var messageFileName string
	messageNameFileName := _utils_.DocTabMessageNameFileName(packageName)
	for messageName, structName := range messageNameStructName {
		messageFileName = messageNameFileName[messageName]
		fPath = filepath.Join(packageTXRXPath, messageFileName)
		switch messageName {
		case _utils_.DocTabMessageB2FAddTabbar:
			data = _txrx_.B2FAddTabbarTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.B2FAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FAddTab:
			data = _txrx_.B2FAddTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				LocalPanelNames:   localPanelNames,
				Funcs:             funcs,
				MessageStructName: structName,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.B2FAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTabbar:
			data = _txrx_.B2FRemoveTabbarTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.B2FRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageB2FRemoveTab:
			data = _txrx_.B2FRemoveTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.B2FRemoveTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTabbar:
			data = _txrx_.F2BAddTabbarTemplateData{
				PackageName:       packageName,
				AllPanelNames:     allPanelNames,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.F2BAddTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BAddTab:
			data = _txrx_.F2BAddTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.F2BAddTabTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTabbar:
			data = _txrx_.F2BRemoveTabbarTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.F2BRemoveTabbarTemplate, data); err != nil {
				return
			}
		case _utils_.DocTabMessageF2BRemoveTab:
			data = _txrx_.F2BRemoveTabTemplateData{
				PackageName:       packageName,
				ImportPrefix:      importPrefix,
				MessageStructName: structName,
				Funcs:             funcs,
			}
			if err = _utils_.ProcessTemplate(messageFileName, fPath, _txrx_.F2BRemoveTabTemplate, data); err != nil {
				return
			}
		}
	}
	// panelers folder.

	// frontend/screens/«screen-package-name»/panelers/panelers.go
	fPath = filepath.Join(packagePanelersPath, _panelers_.PanelersFileName)
	data = &_panelers_.PanelersTemplateData{
		ImportPrefix:    importPrefix,
		LocalPanelNames: localPanelNames,
	}
	if err = _utils_.ProcessTemplate(_panelers_.PanelersFileName, fPath, _panelers_.PanelersTemplate, data); err != nil {
		return
	}

	// producer folder.

	// frontend/screens/«screen-package-name»/producer/docTabs.go
	fPath = filepath.Join(packageProducerPath, _producer_.DocTabsContentProducerFileName)
	data = &_producer_.DocTabsContentProducerTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_producer_.DocTabsContentProducerFileName, fPath, _producer_.DocTabsContentProducerTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/producer/tabItem.go
	fPath = filepath.Join(packageProducerPath, _producer_.TabItemContentProducerFileName)
	data = &_producer_.TabItemContentProducerTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_producer_.TabItemContentProducerFileName, fPath, _producer_.TabItemContentProducerTemplate, data); err != nil {
		return
	}

	// panels folder.

	// frontend/screens/«screen-package-name»/panels.log
	fPath = filepath.Join(packagePanelsPath, _panels_.LogFileName)
	fmt.Printf("fPath is %s\n", fPath)
	content := _panels_.LogContent(packageName, localPanelNames, folderPaths)
	if err = _utils_.WriteFile(fPath, []byte(content)); err != nil {
		return
	}

	// Add each panel's file and sub folder.
	fmt.Printf("localPanelNames is %+v", localPanelNames)
	for _, panelName := range localPanelNames {
		// Panel file.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		fileName = panelName + _panels_.PanelFileNameSuffix
		fPath = filepath.Join(packagePanelsPath, fileName)
		data = &_panels_.PanelTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
			Funcs:        funcs,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panels_.PanelTemplate, data); err != nil {
			return
		}

		// Panel sub folder holding content and state.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		panelFolderName := panelName + "Panel"
		fmt.Printf("making panel folder %s\n", panelFolderName)
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		// content.go
		fileName = _panel_.ContentFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		data = &_panel_.ContentTemplateData{
			PackageName:     packageName,
			PanelName:       panelName,
			LocalPanelNames: localPanelNames,
			ImportPrefix:    importPrefix,
			Funcs:           funcs,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ContentTemplate, data); err != nil {
			return
		}
		// state.go
		fileName = _panel_.StateFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		data = &_panel_.StateTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.StateTemplate, data); err != nil {
			return
		}
		// messenger.go
		fPath = filepath.Join(panelFolderPath, _utils_.MessengerFileName)
		data = &_panel_.MessengerTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
		}
		if err = _utils_.ProcessTemplate(_utils_.MessengerFileName, fPath, _panel_.MessengerTemplate, data); err != nil {
			return
		}
	}
	return
}
