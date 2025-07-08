package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type IDsTemplateData struct {
	PackageName   string
	AllPanelNames []string
	Funcs         _utils_.Funcs
}

const (
	// Use the package name for a file name. Accordion.go package Accordion.
	IDsTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

import(
	"slices"
)
// AccordionItem.
type AccordionItem uint
const (
{{- if ne (len .AllPanelNames) 0 }}
	{{ index .AllPanelNames 0 }}AccordionItem AccordionItem =   iota
{{- end }}
{{- range $i, $panelName := .AllPanelNames }}
 {{- if ne $i 0 }}
	{{ $panelName }}AccordionItem
 {{- end }}
{{- end }}
)

var AccordionMessengerIDAccordionItemMessengerIDs = make(map[string][]string)

func AddAccordionMessengerID(accordionMessengerID string) {
	if _, found := AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID]; !found {
		AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID] = make([]string, 0, 5)
	}
}

func AddAccordionMessengerIDAccordionItemMessengerID(accordionMessengerID, accordionItemMessengerID string) {
	var accordionItemIDs []string
	var found bool
	if accordionItemIDs, found = AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID]; !found {
		accordionItemIDs = make([]string, 0, 5)
	}
	AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID] = append(accordionItemIDs, accordionItemMessengerID)
}

func RemoveAccordionMessengerID(accordionMessengerID string) {
	delete(AccordionMessengerIDAccordionItemMessengerIDs, accordionMessengerID)
}

func RemoveAccordionItemMessengerID(accordionMessengerID, accordionItemMessengerID string) {
	accordionItemIDs := AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID]
	if index := slices.Index(accordionItemIDs, accordionItemMessengerID); index > 0 {
		AccordionMessengerIDAccordionItemMessengerIDs[accordionMessengerID] = slices.Delete(accordionItemIDs, index, index+1)
	}
}
`
)
