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
	window       fyne.Window
	producer     ContentProducer
	isShowing    bool
	isInMainMenu bool
}

func NewWindowContentConsumer(window fyne.Window, isInMainMenu bool) (consumer *WindowContentConsumer) {
	consumer = &WindowContentConsumer{
		window:       window,
		isInMainMenu: isInMainMenu,
	}
	return
}

// ContentConsumer implementations.

func (consumer *WindowContentConsumer) IsMainMenu() (isInMainMenu bool) {
	isInMainMenu = consumer.isInMainMenu
	return
}

// Show sets consumer as the window's content.
// Show is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Show(isMainThread bool) {
	if currentWindowContentConsumer == consumer {
		// This consumer is currently showing.
		return
	}
	if currentWindowContentConsumer != nil {
		currentWindowContentConsumer.isShowing = false
	}
	consumer.isShowing = true
	currentWindowContentConsumer = consumer
	consumer.refresh(isMainThread, true)
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

func (consumer *WindowContentConsumer) CanUnBind() (canUnBind bool) {
	canUnBind = !consumer.isInMainMenu
	return
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

// Refresh refreshes the window if there is something to refresh.
// Refresh is thread safe.
// Refresh is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Refresh(isMainThread bool) {
	consumer.refresh(isMainThread, false)
}

// refresh:
// 1. Moves content from the producer to the window.
// 2. Has the window refresh.
// Param forceCanvasObject indicates if the canvas object must be refreshed no matter what.
// Refresh is thread safe.
// Refresh is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) refresh(isMainThread bool, forceCanvasObject bool) {
	var set bool
	if isMainThread {
		if forceCanvasObject {
			set = true
			canvasObject := consumer.producer.CanvasObjectForce(consumer)
			consumer.window.SetContent(canvasObject)
		} else {
			if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
				set = true
				consumer.window.SetContent(canvasObject)
			}
		}
		if title := consumer.producer.Title(consumer); title != nil && len(*title) > 0 {
			set = true
			consumer.window.SetTitle(*title)
		}
		if set {
			consumer.window.Content().Refresh()
		}

	} else {
		if forceCanvasObject {
			set = true
			canvasObject := consumer.producer.CanvasObjectForce(consumer)
			fyne.Do(
				func() {
					consumer.window.SetContent(canvasObject)
				},
			)
		} else {
			if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
				set = true
				fyne.Do(
					func() {
						consumer.window.SetContent(canvasObject)
					},
				)
			}
		}
		if title := consumer.producer.Title(consumer); title != nil && len(*title) > 0 {
			set = true
			fyne.Do(
				func() {
					consumer.window.SetTitle(*title)
				},
			)
		}
		if set {
			fyne.Do(
				func() {
					consumer.window.Content().Refresh()
				},
			)
		}
	}
}

// IsWindowContentConsumer returns true because this is a window consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) IsWindowContentConsumer() (is bool) {
	is = true
	return
}
`
)
