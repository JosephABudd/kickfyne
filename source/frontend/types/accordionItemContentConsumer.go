package types

const (
	accordionItemContentConsumerFileName = "accordionItemContentConsumer.go"
	accordionItemContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2/widget"
)

// AccordionItemContentConsumer consumes content from a producer and gives it to an accordionItem.
// AccordionItemContentConsumer implements ContentConsumer.
// AccordionItemContentConsumer implements UnSpawner.
type AccordionItemContentConsumer struct {
	accordionItem  *widget.AccordionItem
	producer       ContentProducer
	screenConsumer ContentConsumer
	spawned        bool
}

func NewAccordionItemContentConsumer(accordionItem *widget.AccordionItem, screenConsumer ContentConsumer, spawned bool) (consumer *AccordionItemContentConsumer) {
	consumer = &AccordionItemContentConsumer{
		accordionItem:  accordionItem,
		spawned:        spawned,
		screenConsumer: screenConsumer,
	}
	return
}

// Show sets the AccordionItem's screen's consumer.
func (consumer *AccordionItemContentConsumer) SetScreenConsumer(screenConsumer ContentConsumer) {
	consumer.screenConsumer = screenConsumer
}

// ContentConsumer implementations.

// Show sets the AccordionItem's content.
// Show is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Show() {
	consumer.accordionItem.Detail.Show()
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) IsVisible() (is bool) {
	is = consumer.screenConsumer.IsVisible()
	return
}

// Refresh:
// 1. Moves content from the producer to the accordionIItem.
// 2. Has the accordionIItem refresh.
// Refresh is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Refresh() {
	if label := consumer.producer.Label(consumer); label != nil {
		consumer.accordionItem.Title = *label
	}
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.accordionItem.Detail = canvasObject
	}
	consumer.accordionItem.Detail.Refresh()
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		// Already bound to a producer.
		return
	}
	// Bind to the producer.
	consumer.producer = producer
	producer.Bind(consumer)
}

// UnBind calls the producer's UnBind() and then unspawns.
// UnBind is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// This accordion was spawned so unspawn it.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
	// Remove this accordion from the accordionbar.
}

// IsWindowContentConsumer returns false because this is a accordionItem consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}
`
)
