package types

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/types.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("types.Build: %w", err)
		}
	}()

	var oPath string
	var templateData any

	// frontend/types/interfaces.go.
	oPath = filepath.Join(folderPaths.FrontendTypes, interfacesFileName)
	if err = _utils_.ProcessTemplate(interfacesFileName, oPath, interfacesTemplate, nil); err != nil {
		return
	}

	// frontend/types/accordionItemContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, accordionItemContentConsumerFileName)
	if err = _utils_.ProcessTemplate(accordionItemContentConsumerFileName, oPath, accordionItemContentConsumerTemplate, nil); err != nil {
		return
	}

	// frontend/types/appTabItemContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, appTabItemContentConsumerFileName)
	if err = _utils_.ProcessTemplate(appTabItemContentConsumerFileName, oPath, appTabItemContentConsumerTemplate, nil); err != nil {
		return
	}

	// frontend/types/docTabItemContentConsumer.go
	templateData = &docTabItemContentConsumerTemplateData{
		ImportPrefix: importPrefix,
	}
	oPath = filepath.Join(folderPaths.FrontendTypes, docTabItemContentConsumerFileName)
	if err = _utils_.ProcessTemplate(docTabItemContentConsumerFileName, oPath, docTabItemContentConsumerTemplate, templateData); err != nil {
		return
	}

	// frontend/types/windowContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, windowContentConsumerFileName)
	if err = _utils_.ProcessTemplate(windowContentConsumerFileName, oPath, windowContentConsumerTemplate, nil); err != nil {
		return
	}

	return
}
