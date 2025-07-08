package utils

import (
	"path/filepath"
)

const (
	AppTabMessageB2FAddTabbar    = "B2FAddTabbar"
	AppTabMessageB2FAddTab       = "B2FAddTab"
	AppTabMessageB2FRemoveTabbar = "B2FRemoveTabbar"
	AppTabMessageB2FRemoveTab    = "B2FRemoveTab"

	AppTabMessageF2BAddTabbar    = "F2BAddTabbar"
	AppTabMessageF2BAddTab       = "F2BAddTab"
	AppTabMessageF2BRemoveTabbar = "F2BRemoveTabbar"
	AppTabMessageF2BRemoveTab    = "F2BRemoveTab"
)

var appTabMessageNames = []string{
	AppTabMessageB2FAddTabbar,
	AppTabMessageB2FAddTab,
	AppTabMessageB2FRemoveTabbar,
	AppTabMessageB2FRemoveTab,

	AppTabMessageF2BAddTabbar,
	AppTabMessageF2BAddTab,
	AppTabMessageF2BRemoveTabbar,
	AppTabMessageF2BRemoveTab,
}

var appTabMessageDescription = map[string]string{
	AppTabMessageB2FAddTabbar: "The back-end is telling the front-end to add a new AppTabs tabbar.",
	AppTabMessageF2BAddTabbar: "The front-end is telling the back-end to add a new AppTabs tabbar.",

	AppTabMessageB2FRemoveTabbar: "The back-end is telling the front-end to remove a AppTabs tabbar.",
	AppTabMessageF2BRemoveTabbar: "The front-end is telling the back-end to remove a AppTabs tabbar.",

	AppTabMessageB2FAddTab: "The back-end is telling the front-end to add a new tab to a AppTabs tabbar.",
	AppTabMessageF2BAddTab: "The front-end is telling the back-end to add a new tab to a AppTabs tabbar.",

	AppTabMessageB2FRemoveTab: "The back-end is telling the front-end to remove a tab from a AppTabs tabbar.",
	AppTabMessageF2BRemoveTab: "The front-end is telling the back-end to remove a tab from a AppTabs tabbar.",
}

func AppTabMessageNameDescription() (messageNameDescription map[string]string) {
	messageNameDescription = appTabMessageDescription
	return
}

func AppTabMessageNameStructName(packageName string) (messageNameStructName map[string]string) {
	messageNameStructName = make(map[string]string, len(appTabMessageNames))
	for _, name := range appTabMessageNames {
		messageNameStructName[name] = appTabMessageStructName(packageName, name)
	}
	return
}

func AppTabSharedMessageFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.SharedMessage, packageName)
	return
}

func AppTabBackendMessengerFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.BackendTXRX, packageName)
	return
}

func AppTabMessageNameFileName(packageName string) (messageFileNames map[string]string) {
	messageFileNames = make(map[string]string, len(appTabMessageNames))
	for _, appTabMessageName := range appTabMessageNames {
		messageFileNames[appTabMessageName] = appTabMessageFileName(packageName, appTabMessageName)
	}
	return
}

func appTabMessageFileName(packageName, appTabMessageName string) (fileName string) {
	fileName = DeCap(
		appTabMessageStructName(packageName, appTabMessageName) + GoFileExt,
	)
	return
}

func appTabMessageStructName(packageName, appTabMessageName string) (messageStructName string) {
	messageStructName = appTabMessageName + packageName
	return
}
