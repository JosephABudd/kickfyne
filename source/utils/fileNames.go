package utils

import (
	"path"
	"path/filepath"
)

const (
	APIFileName         = "api.go"
	ContentFileName     = "content.go"
	StateFileName       = "state.go"
	panelFileSuffix     = "Panel.go"
	GoFileExt           = ".go"
	FyneAppTOMLFileName = "FyneApp.toml"
	ScreenFileName      = "screen.go"
	DocFileName         = "doc.go"
	LayoutFileName      = "layout.go"
	MessengerFileName   = "messenger.go"
	PanelsFileName      = "panels.go"
	ComponentsFileName  = "components.go"

	ralativeFilePathSuffix = ":1:1"
)

func Clickable(path string) (clickable string) {
	clickable = path + ralativeFilePathSuffix
	return
}

// FyneAppTOMLFilePath
func FyneAppTOMLFilePath(folderPaths *FolderPaths) (metaDataTOMLFilePath string) {
	metaDataTOMLFilePath = filepath.Join(folderPaths.App, FyneAppTOMLFileName)
	return
}

// PanelFileName returns the file name for a panel file.
func PanelFileName(panelName string) (fileName string) {
	fileName = panelName + GoFileExt
	return
}

// PanelContentFolderName returns the content folder name for a panel.
func PanelContentFolderName(panelName string) (fileName string) {
	fileName = panelName + "Panel"
	return
}

// PanelContentFilePath returns the relative path for a panel content file.
func PanelContentFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, contentFolderName, ContentFileName)
	return
}

// PanelStateFilePath returns the relative path for a panel's content state file.
func PanelStateFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, contentFolderName, StateFileName)
	return
}

// ScreenFileRelativeFilePath returns the relative path for a screen's screen.go file.
func ScreenFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, FolderNameScreens, screenPackageName, ScreenFileName+ralativeFilePathSuffix)
	return
}

// DocFileRelativeFilePath returns the relative path for a screen's screen.go file.
func DocFileRelativeFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, DocFileName)
	return
}

// ScreenDocFilePath returns the relative path for a screen's doc.go file.
func ScreenDocFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, DocFileName)
	return
}

// LayoutFileRelativeFilePath returns the relative path for a screen's layout.go file.
func LayoutFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, FolderNameScreens, screenPackageName, LayoutFileName+ralativeFilePathSuffix)
	return
}

// MessageFileName returns the file name for a messsage.
func MessageFileName(messageName string) (fileName string) {
	fileName = DeCap(messageName) + GoFileExt
	return
}

// MessageFileRelativeFilePath returns the relative path for a message file.
func MessageFileRelativeFilePath(messageName string) (relativeFilePath string) {
	fName := MessageFileName(messageName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameShared, folderNameMessage, fName)
	return
}

// MessageHandlerFileRelativeFilePath returns the relative path for a message file.
func MessageHandlerFileRelativeFilePath(messageName string) (relativeFilePath string) {
	fName := MessageFileName(messageName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameBackend, FolderNameTXRX, fName)
	return
}

// RecordFileName returns the file name for a record.
func RecordFileName(recordName string) (fileName string) {
	fileName = DeCap(recordName) + GoFileExt
	return
}

// RecordFileRelativeFilePath returns the relative path for a record file.
func RecordFileRelativeFilePath(recordName string) (relativeFilePath string) {
	fName := MessageFileName(recordName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameShared, folderNameStore, folderNameRecord, fName)
	return
}

// RecordStorerFileRelativeFilePath returns the relative path for a record's storer file.
func RecordStorerFileRelativeFilePath(recordName string) (relativeFilePath string) {
	fName := MessageFileName(recordName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameShared, folderNameStore, folderNameStorer, fName)
	return
}

// RecordStoringFileRelativeFilePath returns the relative path for a record's storer file.
func RecordStoringFileRelativeFilePath(recordName string) (relativeFilePath string) {
	fName := MessageFileName(recordName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameShared, folderNameStore, folderNameStoring, fName)
	return
}
