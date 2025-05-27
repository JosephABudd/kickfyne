package accordion

import (
	"fmt"
	"os"
	"path/filepath"

	_misc_ "github.com/JosephABudd/kickfyne/source/frontend/screens/accordion/misc"
	_panels_ "github.com/JosephABudd/kickfyne/source/frontend/screens/accordion/panels"
	_content_ "github.com/JosephABudd/kickfyne/source/frontend/screens/accordion/panels/content"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// Build builds a type Simple screen package.
func Build(
	allPanelNames, localPanelNames, remotePanelNames []string,
	packageName string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.Build: %w", err)
		}
	}()

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

	// frontend/screens/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)
	if err = os.Mkdir(packagePanelsPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/panels/content
	packagePanelsContentPath := filepath.Join(packagePath, _utils_.FolderNameContent)
	if err = os.Mkdir(packagePanelsContentPath, _utils_.DMode); err != nil {
		return
	}

	var fPath string
	var data interface{}
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	data = &docTemplateData{
		PackageName: packageName,
		PackageDoc:  packageDoc,
		Funcs:       funcs,
	}
	if err = _utils_.ProcessTemplate(docFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/screen.go
	fPath = filepath.Join(packagePath, screenFileName)
	data = &screenTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(screenFileName, fPath, screenTemplate, data); err != nil {
		return
	}

	fPath = filepath.Join(packagePath, accordionLayoutFileName)
	data = &accordionLayoutTemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(accordionLayoutFileName, fPath, accordionLayoutTemplate, data); err != nil {
		return
	}

	// Add files to the package's misc/ folder.

	// frontend/screens/«screen-package-name»/misc/messageHandler.go
	fPath = filepath.Join(packageMiscPath, _misc_.MessengerFileName)
	data = &_misc_.MessengerTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_misc_.MessengerFileName, fPath, _misc_.MessengerTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/misc/panels.go
	fPath = filepath.Join(packageMiscPath, _utils_.PanelsFileName)
	data = &_misc_.PanelsTemplateData{
		PackageName:      packageName,
		PanelNames:       allPanelNames,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
		DefaultPanelName: allPanelNames[0],
	}
	if err = _utils_.ProcessTemplate(_misc_.PanelsFileName, fPath, _misc_.PanelsTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/misc/screenComponents.go
	fPath = filepath.Join(packageMiscPath, _misc_.ScreenComponentsFileName)
	data = &_misc_.ScreenComponentsTemplateData{
		PackageName:      packageName,
		PanelNames:       allPanelNames,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
		DefaultPanelName: allPanelNames[0],
	}
	if err = _utils_.ProcessTemplate(_misc_.ScreenComponentsFileName, fPath, _misc_.ScreenComponentsTemplate, data); err != nil {
		return
	}

	// Add files to the package's panels/ folder.
	for _, panelName := range localPanelNames {
		// frontend/screens/«screen-package-name»/panels/«panel-name».go
		fileName = _panels_.PanelFileName(panelName)
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
	}

	// Add files to the package's panels/content/ folder.
	for _, panelName := range localPanelNames {
		// frontend/screens/«screen-package-name»/panels/content/«panel-name».go
		fileName = _content_.PanelContentFileName(panelName)
		fPath = filepath.Join(packagePanelsContentPath, fileName)
		data = &_content_.PanelContentTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
			Funcs:        funcs,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _content_.PanelContentTemplate, data); err != nil {
			return
		}
	}

	return
}
