package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type PaneslStateTemplateData struct {
	PackageName     string
	LocalPanelNames []string
	ImportPrefix    string
	Funcs           _utils_.Funcs
}

const (
	PaneslStateTemplate = `{{ $DOT := . -}}
package {{ call .Funcs.LowerCase .PackageName }}

type PanelsState struct {
{{- range $i, $panelName := .LocalPanelNames}}

	// State for the {{ $panelName }} local panel's content.
	// KICKFYNE TODO: Add your own state members here.
	{{ $panelName }}PanelHeading     string
	{{ $panelName }}PanelDescription string
{{- end }}
}
`
)
