package panels

import (
	"fmt"
	"path/filepath"
	"strings"

	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

const (
	LogFileName = "README.log"

	howPanelsWork = `%s DocTabs screen package.:
How panels work.
A panel only produces content for it's tab.
1. content.go is the panel's content.
   * The tab's icon, label and content are defined.
   * The content is layed out for display.
   * The content handles user input events using the messenger to send information to the back-end.
   KICKFYNE TODO: Modify content for your user interface.
2. state.go sets and gets the panel's content.go.
   KICKFYNE TODO: Modify state.go for content.go.
3. messenger.go sends and listens for and receives messages with the back-end.
   * Provides funcs that content uses to send messages to the back-end.
   * Uses the state to set content received from the back-end.
   KICKFYNE TODO: Modify messenger.go for content.go and state.go.
`

	ending = `
Content producers and consumers. (Not required learning.)
1. Each panel produces content for it's own tabItem in the DocTabs tabbar.
   * The panel's producer gives the panel's content to the tabItem's consumer.
   * The tabItem's consumer gives the content to the panel's tabItem.
   * The panel's tabItem is part of the package's DocTabs tabbar.
2. The package has a producer which gives the DocTabs tabbar's content to the package's consumer.
   The package's consumer will be 1 of these 4 consumer types.
   1. A WindowContentConsumer consumes the package's content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the package's content for a single TabItem in a separate AppTab screen.
   3. A DocTabsTabItemContentConsumer consumes the package's content for a single TabItem in a separate DocTabs screen.
   4. An AccordionConsumer consumes the package's content for a single AccordionItem in a separate Accordion screen.
`
)

func LogContent(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (logMessage string) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(howPanelsWork, screenPackageName))
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		messengerPath := _utils_.PanelMessengerFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString("  " + panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("    Content:   %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:     %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Messenger: %s.\n", messengerPath))
	}
	// Shared Messages
	messageNameStructName := _utils_.DocTabMessageNameStructName(screenPackageName)
	messageFolderPath := _utils_.DocTabSharedMessageFolderPath(screenPackageName, folderPaths)
	messageNameFileName := _utils_.DocTabMessageNameFileName(screenPackageName)
	messageNameDescription := _utils_.DocTabMessageNameDescription()
	builder.WriteString("\nShared messages:\n")
	for messageName, messageFileName := range messageNameFileName {
		messagePath := filepath.Join(messageFolderPath, messageFileName)
		structName := messageNameStructName[messageName]
		description := messageNameDescription[messageName]
		builder.WriteString(fmt.Sprintf("  %s: %s\n    %s\n", structName, messagePath, description))
		// builder.WriteString(fmt.Sprintf("  %s: %s\n", structName, _utils_.Clickable(messagePath)))
	}
	// Back-end Messengers
	messengerFolderPath := _utils_.DocTabBackendMessengerFolderPath(screenPackageName, folderPaths)
	builder.WriteString("\nBack-end messengers:\n")
	for messageName, messageFileName := range messageNameFileName {
		messengerPath := filepath.Join(messengerFolderPath, messageFileName)
		structName := messageNameStructName[messageName]
		builder.WriteString(fmt.Sprintf("  %s: %s\n", structName, messengerPath))
		// builder.WriteString(fmt.Sprintf("  %s: %s\n", structName, _utils_.Clickable(messengerPath)))
	}
	// The ending. The boring part.
	builder.WriteString(ending)
	logMessage = builder.String()
	return
}
