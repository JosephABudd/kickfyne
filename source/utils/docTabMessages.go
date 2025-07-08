package utils

import (
	"path/filepath"
)

const (
	DocTabMessageB2FAddTabbar    = "B2FAddTabbar"
	DocTabMessageB2FAddTab       = "B2FAddTab"
	DocTabMessageB2FRemoveTabbar = "B2FRemoveTabbar"
	DocTabMessageB2FRemoveTab    = "B2FRemoveTab"

	DocTabMessageF2BAddTabbar    = "F2BAddTabbar"
	DocTabMessageF2BAddTab       = "F2BAddTab"
	DocTabMessageF2BRemoveTabbar = "F2BRemoveTabbar"
	DocTabMessageF2BRemoveTab    = "F2BRemoveTab"
)

var docTabMessageNames = []string{
	DocTabMessageB2FAddTabbar,
	DocTabMessageB2FAddTab,
	DocTabMessageB2FRemoveTabbar,
	DocTabMessageB2FRemoveTab,

	DocTabMessageF2BAddTabbar,
	DocTabMessageF2BAddTab,
	DocTabMessageF2BRemoveTabbar,
	DocTabMessageF2BRemoveTab,
}

var docTabMessageDescription = map[string]string{
	DocTabMessageB2FAddTabbar: "The back-end is telling the front-end to add a new DocTabs tabbar.",
	DocTabMessageF2BAddTabbar: "The front-end is telling the back-end to add a new DocTabs tabbar.",

	DocTabMessageB2FRemoveTabbar: "The back-end is telling the front-end to remove a DocTabs tabbar.",
	DocTabMessageF2BRemoveTabbar: "The front-end is telling the back-end to remove a DocTabs tabbar.",

	DocTabMessageB2FAddTab: "The back-end is telling the front-end to add a new tab to a DocTabs tabbar.",
	DocTabMessageF2BAddTab: "The front-end is telling the back-end to add a new tab to a DocTabs tabbar.",

	DocTabMessageB2FRemoveTab: "The back-end is telling the front-end to remove a tab from a DocTabs tabbar.",
	DocTabMessageF2BRemoveTab: "The front-end is telling the back-end to remove a tab from a DocTabs tabbar.",
}

func DocTabMessageNameDescription() (messageNameDescription map[string]string) {
	messageNameDescription = docTabMessageDescription
	return
}

func DocTabMessageNameStructName(packageName string) (messageNameStructName map[string]string) {
	messageNameStructName = make(map[string]string, len(docTabMessageNames))
	for _, name := range docTabMessageNames {
		messageNameStructName[name] = docTabMessageStructName(packageName, name)
	}
	return
}

func DocTabSharedMessageFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.SharedMessage, packageName)
	return
}

func DocTabBackendMessengerFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.BackendTXRX, packageName)
	return
}

func DocTabMessageNameFileName(packageName string) (messageFileNames map[string]string) {
	messageFileNames = make(map[string]string, len(docTabMessageNames))
	for _, docTabMessageName := range docTabMessageNames {
		messageFileNames[docTabMessageName] = docTabMessageFileName(packageName, docTabMessageName)
	}
	return
}

func docTabMessageFileName(packageName, docTabMessageName string) (fileName string) {
	fileName = DeCap(
		docTabMessageStructName(packageName, docTabMessageName) + GoFileExt,
	)
	return
}

func docTabMessageStructName(packageName, docTabMessageName string) (messageStructName string) {
	messageStructName = docTabMessageName + packageName
	return
}
