package utils

import (
	"path/filepath"
)

const (
	AccordionMessageB2FAddAccordion        = "B2FAddAccordion"
	AccordionMessageB2FAddAccordionItem    = "B2FAddAccordionItem"
	AccordionMessageB2FRemoveAccordion     = "B2FRemoveAccordion"
	AccordionMessageB2FRemoveAccordionItem = "B2FRemoveAccordionItem"

	AccordionMessageF2BAddAccordion        = "F2BAddAccordion"
	AccordionMessageF2BAddAccordionItem    = "F2BAddAccordionItem"
	AccordionMessageF2BRemoveAccordion     = "F2BRemoveAccordion"
	AccordionMessageF2BRemoveAccordionItem = "F2BRemoveAccordionItem"
)

var accordionMessageNames = []string{
	AccordionMessageB2FAddAccordion,
	AccordionMessageB2FAddAccordionItem,
	AccordionMessageB2FRemoveAccordion,
	AccordionMessageB2FRemoveAccordionItem,

	AccordionMessageF2BAddAccordion,
	AccordionMessageF2BAddAccordionItem,
	AccordionMessageF2BRemoveAccordion,
	AccordionMessageF2BRemoveAccordionItem,
}

var accordionMessageDescription = map[string]string{
	AccordionMessageB2FAddAccordion: "The back-end is telling the front-end to add an Accordion.",
	AccordionMessageF2BAddAccordion: "The front-end is telling the back-end to add an Accordion.",

	AccordionMessageB2FRemoveAccordion: "The back-end is telling the front-end to remove a Accordion.",
	AccordionMessageF2BRemoveAccordion: "The front-end is telling the back-end to remove a Accordion.",

	AccordionMessageB2FAddAccordionItem: "The back-end is telling the front-end to add an AccordionItem to a Accordion.",
	AccordionMessageF2BAddAccordionItem: "The front-end is telling the back-end to add an AccordionItem to a Accordion.",

	AccordionMessageB2FRemoveAccordionItem: "The back-end is telling the front-end to remove an AccordionItem from a Accordion.",
	AccordionMessageF2BRemoveAccordionItem: "The front-end is telling the back-end to remove an AccordionItem from a Accordion.",
}

func AccordionMessageNameDescription() (messageNameDescription map[string]string) {
	messageNameDescription = accordionMessageDescription
	return
}

func AccordionMessageNameStructName(packageName string) (messageNameStructName map[string]string) {
	messageNameStructName = make(map[string]string, len(accordionMessageNames))
	for _, name := range accordionMessageNames {
		messageNameStructName[name] = accordionMessageStructName(packageName, name)
	}
	return
}

func AccordionSharedMessageFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.SharedMessage, packageName)
	return
}

func AccordionBackendMessengerFolderPath(packageName string, folderPaths *FolderPaths) (folderPath string) {
	folderPath = filepath.Join(folderPaths.BackendTXRX, packageName)
	return
}

func AccordionMessageNameFileName(packageName string) (messageFileNames map[string]string) {
	messageFileNames = make(map[string]string, len(accordionMessageNames))
	for _, accordionMessageName := range accordionMessageNames {
		messageFileNames[accordionMessageName] = accordionMessageFileName(packageName, accordionMessageName)
	}
	return
}

func accordionMessageFileName(packageName, accordionMessageName string) (fileName string) {
	fileName = DeCap(
		accordionMessageStructName(packageName, accordionMessageName) + GoFileExt,
	)
	return
}

func accordionMessageStructName(packageName, accordionMessageName string) (messageStructName string) {
	messageStructName = accordionMessageName + packageName
	return
}
