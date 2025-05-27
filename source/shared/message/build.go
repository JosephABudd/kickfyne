package message

import (
	"errors"
	"os"
	"path/filepath"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

func addDocTabs(
	packageName string,
	localPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	{
		// SpawnedTab message.
		fileName := SpawnedTabFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		templateData := &SpawnedTabTemplateData{
			DocTabPackageName: packageName,
			ImportPrefix:      importPrefix,
		}
		template := SpawnedTabTemplate
		if err = _utils_.ProcessTemplate(fileName, oPath, template, templateData); err != nil {
			return
		}
	}
	{
		// SpawnedTabbar message.
		fileName := SpawnedTabbarFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		templateData := &SpawnedTabbarTemplateData{
			DocTabPackageName: packageName,
			ImportPrefix:      importPrefix,
		}
		template := SpawnedTabbarTemplate
		if err = _utils_.ProcessTemplate(fileName, oPath, template, templateData); err != nil {
			return
		}
	}
	{
		// SpawnTab message.
		fileName := SpawnTabFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		templateData := &SpawnTabTemplateData{
			DocTabPackageName: packageName,
			ImportPrefix:      importPrefix,
		}
		template := SpawnTabTemplate
		if err = _utils_.ProcessTemplate(fileName, oPath, template, templateData); err != nil {
			return
		}
	}
	{
		// SpawnTabbar message.
		fileName := SpawnTabbarFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		templateData := &SpawnTabbarTemplateData{
			DocTabPackageName: packageName,
			ImportPrefix:      importPrefix,
		}
		template := SpawnTabbarTemplate
		if err = _utils_.ProcessTemplate(fileName, oPath, template, templateData); err != nil {
			return
		}
	}
	return
}

func removeDocTabs(
	packageName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (errs error) {
	{
		// SpawnedTab message.
		fileName := SpawnedTabFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		if err := os.Remove(oPath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	{
		// SpawnedTabbar message.
		fileName := SpawnedTabbarFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		if err := os.Remove(oPath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	{
		// SpawnTab message.
		fileName := SpawnTabFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		if err := os.Remove(oPath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	{
		// SpawnTabbar message.
		fileName := SpawnTabbarFileName(packageName)
		oPath := filepath.Join(folderPaths.SharedMessage, fileName)
		if err := os.Remove(oPath); err != nil && !os.IsNotExist(err) {
			errs = errors.Join(errs, err)
			return
		}
	}
	return
}

// CreateFramework creates the shared/message/ files.
func CreateFramework(
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var oPath string

	// message/chans.go
	oPath = filepath.Join(folderPaths.SharedMessage, chansFileName)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = _utils_.ProcessTemplate(chansFileName, oPath, chansTemplate, nil); err != nil {
		return
	}

	// message/init.go
	oPath = filepath.Join(folderPaths.SharedMessage, initFileName)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = _utils_.ProcessTemplate(initFileName, oPath, initTemplate, nil); err != nil {
		return
	}

	return
}
