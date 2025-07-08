package doctabs

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type IDsTemplateData struct {
	PackageName   string
	AllPanelNames []string
	Funcs         _utils_.Funcs
}

const (
	// Use the package name for a file name. Tabbar.go package Tabbar.
	TabsIDsTemplate = `package {{ call .Funcs.LowerCase .PackageName }}

import(
	"slices"
)
// Tab.
type Tab uint
const (
{{- if ne (len .AllPanelNames) 0 }}
	{{ index .AllPanelNames 0 }}Tab Tab =   iota
{{- end }}
{{- range $i, $panelName := .AllPanelNames }}
 {{- if ne $i 0 }}
	{{ $panelName }}Tab
 {{- end }}
{{- end }}
)

var TabbarMessengerIDTabMessengerIDs = make(map[string][]string)

func AddTabbarMessengerID(tabbarMessengerID string) {
	if _, found := TabbarMessengerIDTabMessengerIDs[tabbarMessengerID]; !found {
		TabbarMessengerIDTabMessengerIDs[tabbarMessengerID] = make([]string, 0, 5)
	}
}

func AddTabbarMessengerIDTabMessengerID(tabbarMessengerID, tabMessengerID string) {
	var tabIDs []string
	var found bool
	if tabIDs, found = TabbarMessengerIDTabMessengerIDs[tabbarMessengerID]; !found {
		tabIDs = make([]string, 0, 5)
	}
	TabbarMessengerIDTabMessengerIDs[tabbarMessengerID] = append(tabIDs, tabMessengerID)
}

func RemoveTabbarMessengerID(tabbarMessengerID string) {
	delete(TabbarMessengerIDTabMessengerIDs, tabbarMessengerID)
}

func RemoveTabMessengerID(tabbarMessengerID, tabMessengerID string) {
	tabIDs := TabbarMessengerIDTabMessengerIDs[tabbarMessengerID]
	if index := slices.Index(tabIDs, tabMessengerID); index > 0 {
		TabbarMessengerIDTabMessengerIDs[tabbarMessengerID] = slices.Delete(tabIDs, index, index+1)
	}
}
`
)
