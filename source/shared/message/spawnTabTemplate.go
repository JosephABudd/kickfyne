package message

import (
	"fmt"
)

type SpawnTabTemplateData struct {
	DocTabPackageName string
	ImportPrefix      string
	LocalPanelNames   []string
}

func SpawnTabFileName(docTabPackageName string) (filename string) {
	filename = fmt.Sprintf("spawn%sTab.go", docTabPackageName)
	return
}

const (
	SpawnTabTemplate = `package message

import (
	"fmt"

	"fyne.io/fyne/v2"

	_doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_ "{{ .ImportPrefix }}/shared/doctabs/{{ .DocTabPackageName }}"
)

var Spawn{{ .DocTabPackageName }}TabID = NextID()

// Back to front request to spawn a new {{ .DocTabPackageName }} tab-bar.
type Spawn{{ .DocTabPackageName }}Tab struct {
	id   uint64
	name string

	// These are sent back to front.
	RemoteData _doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_.TabRemoteData
	Tab        _doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_.Tab
	TabIcon    fyne.Resource
	TabLabel   string
	// Hello panel state.
	HelloPanelHeading     string // Back to front.
	HelloPanelDescription string // Back to front.
	// HelloAgain panel state.
	HelloAgainPanelHeading     string // Back to front.
	HelloAgainPanelDescription string // Back to front.

	// These are sent from front to back.
	// Messengers.
	TabbarMessengerID  string
	TabMessengerID string
	// Errors.
	Error        bool
	Fatal        bool
	ErrorMessage string
}

// NewSpawn{{ .DocTabPackageName }}Tab returns a *Spawn{{ .DocTabPackageName }}Tab message.
func NewSpawn{{ .DocTabPackageName }}Tab(
	tabbarMessengerID string,
	tab _doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_.Tab,
	initData _doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_.TabInitData,
) (msg *Spawn{{ .DocTabPackageName }}Tab) {
	msg = &Spawn{{ .DocTabPackageName }}Tab{
		id:   Spawn{{ .DocTabPackageName }}TabID,
		name: "Spawn{{ .DocTabPackageName }}Tab",

		TabbarMessengerID: tabbarMessengerID,
		Tab:               tab,
		TabIcon:           initData.TabIcon,
		TabLabel:          initData.TabLabel,
		RemoteData:        initData.RemoteData,
{{- range $panelName := .LocalPanelNames }}

		// {{ $panelName }} panel state.
		{{ $panelName }}PanelHeading:     initData.{{ $panelName }}PanelHeading,
		{{ $panelName }}PanelDescription: initData.{{ $panelName }}PanelDescription,
{{- end }}
	}
	return
}

// Spawn{{ .DocTabPackageName }}Tab implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *Spawn{{ .DocTabPackageName }}Tab) ID() (id uint64) {
	id = msg.id
	return
}

// FrontendMessengerID returns the id of the front-end screen.
func (msg *Spawn{{ .DocTabPackageName }}Tab) FrontendMessengerID() (id string) {
	id = msg.TabbarMessengerID
	return
}

// Name returns the message's Name.
func (msg *Spawn{{ .DocTabPackageName }}Tab) Name() (name string) {
	name = msg.name
	return
}

// FatalError returns false since there can be no fatal error at the back-end.
func (msg *Spawn{{ .DocTabPackageName }}Tab) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	screenID = msg.TabbarMessengerID
	return
}

// AsInterface returns the message as an any.
func (msg *Spawn{{ .DocTabPackageName }}Tab) AsInterface() (m any) {
	m = msg
	return
}
`
)
