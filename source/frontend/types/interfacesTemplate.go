package types

const (
	interfacesFileName = "interfaces.go"
	interfacesTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type StateSetter func(isMainThread bool) (refreshCanvasObject bool)

type Stater interface {
	Get() (state any)
	Set(setters ...StateSetter)
}

// Panel
type Paneler interface {
	ID() (id string)
	Show(isMainThread bool)
	State() (state Stater)
	CanvasObject() (content fyne.CanvasObject)
	Producer() (producer ContentProducer)
	UnBindCleanUP()
}

type ContentConsumer interface {
	Show(isMainThread bool)
	IsVisible() (is bool)
	Refresh(isMainThread bool)
	Bind(producer ContentProducer)
	UnBind() // Call producer.UnBind(self). Delete self. WindowContentProducer does nothing.
	IsWindowContentConsumer() (is bool)
}

// ContentProducer produces the content for a ContentConsumer.
type ContentProducer interface {
	// Window, TabItem, AccordionItem funcs.
	CanvasObjectForce(consumer ContentConsumer) (canvasObject fyne.CanvasObject)
	CanvasObject(consumer ContentConsumer) (canvasObject fyne.CanvasObject)
	SetCanvasObject(canvasObject fyne.CanvasObject)
	Bind(consumer ContentConsumer)
	UnBind(consumer ContentConsumer) //Stop using this consumer. Delete the package if no other consumers.
	IsVisible() (is bool)
	Refresh(isMainThread bool)

	// Window only func.
	Title(consumer ContentConsumer) (title *string)

	// TabItem only func.
	Icon(consumer ContentConsumer) (icon fyne.Resource)

	// TabItem and AccordionItem only funcs.
	Label(consumer ContentConsumer) (label *string)
}

// StartupMessenger sends a message to the backend
type StartupMessenger interface {
	LoadStartupData(data any)
	StopReceiving()
}
`
)
