package types

const (
	windowContentConsumerFileName = "windowContentConsumer.go"
	windowContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

// currentWindowContentConsumer is the current consumer for the app window content.
var currentWindowContentConsumer *WindowContentConsumer

// WindowContentConsumer consumes content from a producer and gives it to a window.
// WindowContentConsumer implements ContentConsumer
// It has no producer but instead has WindowContentConsumer.canvasObject which is a screen's canvas object.
type WindowContentConsumer struct {
	window    fyne.Window
	producer  ContentProducer
	isShowing bool
}

func NewWindowContentConsumer(window fyne.Window) (consumer *WindowContentConsumer) {
	consumer = &WindowContentConsumer{
		window: window,
	}
	return
}

// ContentConsumer implementations.

// Show sets consumer as the window's content.
// Show is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Show() {
	if currentWindowContentConsumer == consumer {
		// This consumer is currently showing.
		return
	}
	if currentWindowContentConsumer != nil {
		currentWindowContentConsumer.isShowing = false
	}
	consumer.isShowing = true
	currentWindowContentConsumer = consumer
	consumer.Refresh()
}

// IsVisible returns if this content is visible in the window.
// Show is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) IsVisible() (is bool) {
	is = consumer.isShowing
	return
}

// Bind binds to the producer and calls the screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		// Already bound to a producer.
		return
	}
	// Bind to the producer.
	consumer.producer = producer
	producer.Bind(consumer)
}

// UnBind unbinds the consumer from it's producer.
// Bind is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// UnBind from the producer.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
}

// Refresh:
// 1. Moves content from the producer to the window.
// 2. Has the window refresh.
// Refresh is thread safe.
// Refresh is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Refresh() {
	fyne.Do(
		func() {
			if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
				consumer.window.SetContent(canvasObject)
			}
			if title := consumer.producer.Title(consumer); len(title) > 0 {
				consumer.window.SetTitle(title)
			}
			// consumer.window.Content().Refresh()
		},
	)
}

// IsWindowContentConsumer returns true because this is a window consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) IsWindowContentConsumer() (is bool) {
	is = true
	return
}
`
)
