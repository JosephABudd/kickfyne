package types

const (
	interfacesFileName = "interfaces.go"
	interfacesTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type StateSetter func()

type Stater interface {
	Get() (state any)
	Set(state ...StateSetter)
}

// Panel
type Paneler interface {
	ID() (id int64)
	Show()
	State() (state Stater)
	CanvasObject() (content fyne.CanvasObject)
	Producer() (producer ContentProducer)
}

type ContentConsumer interface {
	Show()
	IsVisible() (is bool)
	Refresh()
	Bind(producer ContentProducer)
	UnBind() // Call producer.UnBind(self). Delete self. WindowContentProducer does nothing.
	IsWindowContentConsumer() (is bool)
}

// ContentProducer produces the content for a ContentConsumer.
type ContentProducer interface {
	// Window, TabItem, AccordionItem funcs.
	CanvasObject(consumer ContentConsumer) (canvasObject fyne.CanvasObject)
	Bind(consumer ContentConsumer)
	UnBind(consumer ContentConsumer) //Stop using this consumer. Delete the package if no other consumers.
	IsVisible() (is bool)

	// Window only func.
	Title(consumer ContentConsumer) (title *string)

	// TabItem only func.
	Icon(consumer ContentConsumer) (icon fyne.Resource)

	// TabItem and AccordionItem only funcs.
	Label(consumer ContentConsumer) (label *string)
}

// MessageSpawner sends a message to the backend
type MessageSpawner interface {
	StopReceiving()
}
`
)
