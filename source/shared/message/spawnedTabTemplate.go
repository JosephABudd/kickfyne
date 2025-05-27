package message

import (
	"fmt"
)

type SpawnedTabTemplateData struct {
	DocTabPackageName string
	ImportPrefix      string
}

func SpawnedTabFileName(docTabPackageName string) (filename string) {
	filename = fmt.Sprintf("spawned%sTab.go", docTabPackageName)
	return
}

const (
	SpawnedTabTemplate = `package message

import (
	_doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_ "{{ .ImportPrefix }}/shared/doctabs/{{ .DocTabPackageName }}"
)

var Spawned{{ .DocTabPackageName }}TabID = NextID()

// This goes front to back to front.
// It declares that the front-end has spawned a  new tab in the {{ .DocTabPackageName }} Tabbar.
type Spawned{{ .DocTabPackageName }}Tab struct {
	id   uint64
	name string

	// These are sent from the front to the back.
	TabbarMessengerID string
	TabMessengerID    string
	Tab               _doctabs_tabbar_.Tab
	RemoteData        _doctabs_tabbar_.TabRemoteData

	// These are sent from the back to the front
	Error        bool   // Back to front.
	Fatal        bool   // Back to front.
	ErrorMessage string // Back to front.
}

// NewSpawned{{ .DocTabPackageName }}Tab returns a *Spawned{{ .DocTabPackageName }}Tab message.
func NewSpawned{{ .DocTabPackageName }}Tab(tabbarMessengerID string, tabMessengerID string, tab _doctabs_tabbar_.Tab, remoteData _doctabs_tabbar_.TabRemoteData) (msg *Spawned{{ .DocTabPackageName }}Tab) {
	msg = &Spawned{{ .DocTabPackageName }}Tab{
		id:   Spawned{{ .DocTabPackageName }}TabID,
		name: "Spawned{{ .DocTabPackageName }}Tab",

		TabbarMessengerID: tabbarMessengerID,
		TabMessengerID:    tabMessengerID,
		Tab:               tab,
		RemoteData:        remoteData,
	}
	return
}

// Spawned{{ .DocTabPackageName }}Tab implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *Spawned{{ .DocTabPackageName }}Tab) ID() (id uint64) {
	id = msg.id
	return
}

// FrontendMessengerID returns the id of the front-end screen.
func (msg *Spawned{{ .DocTabPackageName }}Tab) FrontendMessengerID() (id string) {
	id = msg.TabMessengerID
	return
}

// Name returns the message's Name.
func (msg *Spawned{{ .DocTabPackageName }}Tab) Name() (name string) {
	name = msg.name
	return
}

// AsInterface returns the message as an any.
func (msg *Spawned{{ .DocTabPackageName }}Tab) AsInterface() (m any) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *Spawned{{ .DocTabPackageName }}Tab) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	messengerID = msg.TabMessengerID
	return
}

`
)
