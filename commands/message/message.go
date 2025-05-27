package message

import (
	"fmt"

	_source_ "github.com/JosephABudd/kickfyne/source"
	_txrx_ "github.com/JosephABudd/kickfyne/source/backend/txrx"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

const (
	Cmd = "message"

	verbAdd    = "add"
	verbRemove = "remove"
	verbList   = "list"
	verbHelp   = "help"
)

// Handler handles all message commands.
func Handler(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.Handler: %w", err)
		}
	}()

	switch args[0] {
	case verbAdd:
		if !isBuilt {
			fmt.Println("The app must be initailized before a message can be added.")
			return
		}
		if len(args) < 2 {
			fmt.Println(Usage())
			return
		}
		var folderPaths *_utils_.FolderPaths
		if folderPaths, err = _utils_.BuildFolderPaths(pathWD); err != nil {
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = _utils_.ValidateNewMessageName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			fmt.Println(errMessage)
			return
		}
		// Add a shared message.
		if err = _source_.AddMessage(args[1], folderPaths); err != nil {
			return
		}
		// Add a back-end message handler.
		if err = _txrx_.AddMessageHandler(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		messageDefinitionPath := _utils_.MessageFileRelativeFilePath(args[1])
		messageHandlerPath := _utils_.MessageHandlerFileRelativeFilePath(args[1])
		fmt.Printf("Success. Message named %q added.\n", args[1])
		fmt.Printf("KICKFYNE TODO: The message definition at %s may need some editing.\n", messageDefinitionPath)
		fmt.Printf("KICKFYNE TODO: The back-end message handler at %s may need some editing.\n", messageHandlerPath)
	case verbRemove:
		if !isBuilt {
			fmt.Println("The app must be initailized before a message can be removed.")
			return
		}
		if len(args) < 2 {
			fmt.Println(Usage())
			return
		}
		var folderPaths *_utils_.FolderPaths
		if folderPaths, err = _utils_.BuildFolderPaths(pathWD); err != nil {
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = _utils_.ValidateRemoveCurrentMessageName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			fmt.Println(errMessage)
			return
		}
		// Remove a message.
		if err = _source_.RemoveMessage(args[1], folderPaths); err != nil {
			return
		}
		// Remove a back-end message handler.
		if err = _txrx_.RemoveMessageHandler(args[1], folderPaths); err != nil {
			return
		}
		fmt.Printf("Success. Message named %q removed.\n", args[1])
	case verbList:
		if !isBuilt {
			fmt.Println("The app must be initailized before a message names can be listed.")
			return
		}
		var folderPaths *_utils_.FolderPaths
		if folderPaths, err = _utils_.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// List all of the messages.
		var messageNames []string
		if messageNames, err = _utils_.UserMessageNames(folderPaths); err != nil {
			return
		}
		fmt.Printf("There are %d message names:\n", len(messageNames))
		for i, messageName := range messageNames {
			j := i + 1
			switch {
			case j < 10:
				fmt.Printf("  %d  %s\n", j, messageName)
			default:
				fmt.Printf("  %d %s\n", j, messageName)
			}
		}
	case verbHelp:
		fmt.Println(Usage())
	default:
		fmt.Println(Usage())
	}
	return
}
