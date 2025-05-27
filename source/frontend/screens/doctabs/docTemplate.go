package simple

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type docTemplateData struct {
	PackageName string
	PackageDoc  string
	Funcs       _utils_.Funcs
}

const (
	docFileName = _utils_.DocFileName

	docTemplate = `{{ call .Funcs.Comment .PackageDoc }}
package {{ .PackageName }}
/*

Here is how this package works from top (the app window) to bottom (a panel's content).
At the top:
api.go:
  A "Consumer" of this screen's output from the Layout.
  The consumer is of one of 4 types below.
  1. A WindowContentConsumer consumers the content for the entire window.
  2. An AppTabItemContentConsumer consumes the content for a single TabItem in a separate AppTab screen.
  3. A DocTabItemContentConsumer consumes the content for a single TabItem in a separate DocTab screen.
  3. An AccordionConsumer consumes the content for a single AccordionItem in a separate Accordion screen.
layout/layout.go:
	"Producer" of tabbar content for any of the 4 consumers above.
	DocTabs Tabbar.
tabs/tag.go:
	"Consumer" of panel content for a tabItem.
panels/anypanel/content.go:
	"Producer" of panel content for a tabItem's consumer.
	Content.
	State.
*/
`
)
