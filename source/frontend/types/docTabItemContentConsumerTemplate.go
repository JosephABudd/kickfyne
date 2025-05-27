package types

const (
	docTabItemContentConsumerFileName = "docTabItemContentConsumer.go"
	docTabItemContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// DocTabItemContentConsumer consumes content from a producer and gives it to a tabItem.
// It is implemented by a tab item.
// DocTabItemContentConsumer implements ContentConsumer.
// DocTabItemContentConsumer implements UnSpawner.
type DocTabItemContentConsumer struct {
	tabbar   *container.DocTabs
	tabItem  *container.TabItem
	producer ContentProducer // Panel's content producer
	spawned  bool
}

func NewDocTabItemContentConsumer(tabbar *container.DocTabs, tabItem *container.TabItem, spawned bool) (consumer *DocTabItemContentConsumer) {
	consumer = &DocTabItemContentConsumer{
		tabbar:  tabbar,
		tabItem: tabItem,
		spawned: spawned,
	}
	return
}

// ContentConsumer implementations.

// Show sets the TabItem's content.
// Show is the implementation of ScreenCanvasWatcher.
func (consumer *DocTabItemContentConsumer) Show() {
	consumer.tabItem.Content.Show()
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
// TODO: fix this.
func (consumer *DocTabItemContentConsumer) IsVisible() (is bool) {
	is = true
	return
}

// Refresh:
// 1. Moves content from the producer to the tabItem.
// 2. Refreshes the tabItem.
// 3. Refreshes the tab-bar.
// Refresh is the implementation of ContentConsumer.
func (consumer *DocTabItemContentConsumer) Refresh() {
	if icon := consumer.producer.Icon(consumer); icon != nil {
		consumer.tabItem.Icon = icon
	}
	if label := consumer.producer.Label(consumer); len(label) > 0 {
		consumer.tabItem.Text = label
	}
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.tabItem.Content = canvasObject
	}
	fyne.Do(consumer.tabbar.Refresh)
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *DocTabItemContentConsumer) Bind(producer ContentProducer) {
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
func (consumer *DocTabItemContentConsumer) UnBind() {
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
func (consumer *DocTabItemContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}
`
)
