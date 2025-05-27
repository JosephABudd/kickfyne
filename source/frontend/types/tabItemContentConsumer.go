package types

const (
	tabItemContentConsumerFileName = "tabItemContentConsumer.go"
	tabItemContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2/container"
)

// TabItemContentConsumer consumes content from a producer and gives it to a tabItem.
// It is implemented by a tab item.
// TabItemContentConsumer implements ContentConsumer.
// TabItemContentConsumer implements UnSpawner.
type TabItemContentConsumer struct {
	tabItem  *container.TabItem
	producer ContentProducer // Panel's content producer
	spawned  bool
}

func NewTabItemContentConsumer(tabItem *container.TabItem, spawned bool) (consumer *TabItemContentConsumer) {
	consumer = &TabItemContentConsumer{
		tabItem: tabItem,
		spawned: spawned,
	}
	return
}

// ContentConsumer implementations.

// Show sets the TabItem's content.
// Show is the implementation of ScreenCanvasWatcher.
func (consumer *TabItemContentConsumer) Show() {
	consumer.tabItem.Content.Show()
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
// TODO: fix this.
func (consumer *TabItemContentConsumer) IsVisible() (is bool) {
	is = true
	return
}

// Refresh:
// 1. Moves content from the producer to the tabItem.
// 2. Has the tabItem refresh.
// Refresh is the implementation of ContentConsumer.
func (consumer *TabItemContentConsumer) Refresh() {
	if icon := consumer.producer.Icon(consumer); icon != nil {
		consumer.tabItem.Icon = icon
	}
	if label := consumer.producer.Label(consumer); label != nil {
		consumer.tabItem.Text = *label
	}
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.tabItem.Content = canvasObject
	}
	consumer.tabItem.Content.Refresh()
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *TabItemContentConsumer) Bind(producer ContentProducer) {
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
func (consumer *TabItemContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// UnBind from the producer.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
}

// IsWindowContentConsumer returns false because this is a tabItem consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *TabItemContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}
`
)
