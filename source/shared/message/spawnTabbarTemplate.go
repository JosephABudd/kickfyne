package message

import (
	"fmt"
)

type SpawnTabbarTemplateData struct {
	DocTabPackageName string
	ImportPrefix      string
}

func SpawnTabbarFileName(docTabPackageName string) (filename string) {
	filename = fmt.Sprintf("spawn%sTabbar.go", docTabPackageName)
	return
}

const (
	SpawnTabbarTemplate = `package message

import (
	_doctabs_{{ .Funcs.LowerCase .DocTabPackageName }}_ "{{ .ImportPrefix }}/shared/doctabs/{{ .DocTabPackageName }}"
)

var SpawnTabbarID = NextID()

// Back to front request to spawn a new tab.
type SpawnTabbar struct {
	id   uint64
	name string

	// These are sent back to front.
	RemoteData        _doctabs_tabbar_.TabbarRemoteData

	// These are sent front to back.
	TabbarMessengerID string
	Error        bool
	Fatal        bool
	ErrorMessage string
}

// NewSpawnTabbar returns a *SpawnTabbar message.
func NewSpawnTabbar(
	spawnedMessengerID string,
	remoteData _doctabs_tabbar_.TabbarRemoteData,
) (msg *SpawnTabbar) {
	msg = &SpawnTabbar{
		id:   SpawnTabbarID,
		name: "SpawnTabbar",

		TabbarMessengerID: spawnedMessengerID,
		RemoteData: remoteData,
	}
	return
}

// SpawnTabbar implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *SpawnTabbar) ID() (id uint64) {
	id = msg.id
	return
}

// FrontendMessengerID returns the id of the front-end screen.
func (msg *SpawnTabbar) FrontendMessengerID() (id string) {
	id = msg.TabbarMessengerID
	return
}

// Name returns the message's Name.
func (msg *SpawnTabbar) Name() (name string) {
	name = msg.name
	return
}

// AsInterface returns the message as an any.
func (msg *SpawnTabbar) AsInterface() (m any) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *SpawnTabbar) FatalError() (fatal bool, message, messengerID string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	screenID = msg.TabbarMessengerID
	return
}
`
)
