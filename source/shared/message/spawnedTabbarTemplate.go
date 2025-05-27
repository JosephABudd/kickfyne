package message

import (
	"fmt"
)

type SpawnedTabbarTemplateData struct {
	DocTabPackageName string
	ImportPrefix      string
}

func SpawnedTabbarFileName(docTabPackageName string) (filename string) {
	filename = fmt.Sprintf("spawned%sTabbar.go", docTabPackageName)
	return
}

const (
	SpawnedTabbarTemplate = `package message

import (
	_doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_ "{{ .ImportPrefix }}/shared/doctabs/{{ .DocTabPackageName }}"
)

var Spawned{{ .DocTabPackageName }}TabbarID = NextID()

// This goes front to back to front.
// It declares that the front-end has spawned a  new tab in the {{ .DocTabPackageName }} Tabbar.
type Spawned{{ .DocTabPackageName }}Tabbar struct {
	id   uint64
	name string

	// These are sent from the front to the back.
	TabbarMessengerID string
	RemoteData        _doctabs_tabbar_.TabbarRemoteData

	// These are sent from the back to the front
	Error        bool   // Back to front.
	Fatal        bool   // Back to front.
	ErrorMessage string // Back to front.
}

// NewSpawned{{ .DocTabPackageName }}Tab returns a *Spawned{{ .DocTabPackageName }}Tabbar message.
func NewSpawned{{ .DocTabPackageName }}Tab(tabbarMessengerID string, remoteData _doctabs_tabbar_.TabbarRemoteData) (msg *Spawned{{ .DocTabPackageName }}Tabbar) {
	msg = &Spawned{{ .DocTabPackageName }}Tabbar{
		id:   Spawned{{ .DocTabPackageName }}TabID,
		name: "Spawned{{ .DocTabPackageName }}Tabbar",

		TabbarMessengerID: tabbarMessengerID,
		RemoteData:        remoteData,
	}
	return
}

// Spawned{{ .DocTabPackageName }}Tabbar implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *Spawned{{ .DocTabPackageName }}Tabbar) ID() (id uint64) {
	id = msg.id
	return
}

// FrontendMessengerID returns the id of the front-end screen.
func (msg *Spawned{{ .DocTabPackageName }}Tabbar) FrontendMessengerID() (id string) {
	id = msg.TabMessengerID
	return
}

// Name returns the message's Name.
func (msg *Spawned{{ .DocTabPackageName }}Tabbar) Name() (name string) {
	name = msg.name
	return
}

// AsInterface returns the message as an any.
func (msg *Spawned{{ .DocTabPackageName }}Tabbar) AsInterface() (m any) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *Spawned{{ .DocTabPackageName }}Tabbar) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	messengerID = msg.TabMessengerID
	return
}

`
)
