package record

import (
	"fmt"

	_source_ "github.com/JosephABudd/kickfyne/source"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

const (
	Cmd = "record"

	verbAdd    = "add"
	verbRemove = "remove"
	verbList   = "list"
	verbHelp   = "help"
)

// Handler handles all record commands.
func Handler(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.Handler: %w", err)
		}
	}()

	switch args[0] {
	case verbAdd:
		if !isBuilt {
			fmt.Println("The app must be initailized before a record can be added.")
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
		if isValid, errMessage, err = _utils_.ValidateNewRecordName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			fmt.Println(errMessage)
			return
		}
		// Add a record.
		if err = _source_.AddRecord(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		recordPath := _utils_.RecordFileRelativeFilePath(args[1])
		recordStorerPath := _utils_.RecordStorerFileRelativeFilePath(args[1])
		recordStoringPath := _utils_.RecordStoringFileRelativeFilePath(args[1])
		fmt.Printf("Success. Record named %q added.\n", args[1])
		fmt.Printf("KICKFYNE TODO: The record definition at %s may need some editing.\n", recordPath)
		fmt.Printf("KICKFYNE TODO: The storer interface definition at %s may need some editing.\n", recordStorerPath)
		fmt.Printf("KICKFYNE TODO: The storer interface implementation at %s may need some editing.\n", recordStoringPath)
	case verbRemove:
		if !isBuilt {
			fmt.Println("The app must be initailized before a record can be removed.")
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
		if isValid, errMessage, err = _utils_.ValidateCurrentRecordName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			fmt.Println(errMessage)
			return
		}
		// Remove a record.
		if err = _source_.RemoveRecord(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		fmt.Printf("Success. Record named %q removed.", args[1])
	case verbList:
		if !isBuilt {
			fmt.Println("The app must be initailized before a record names can be listed.")
			return
		}
		var folderPaths *_utils_.FolderPaths
		if folderPaths, err = _utils_.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// List all of the records.
		var recordNames []string
		if recordNames, err = _utils_.RecordNames(folderPaths); err != nil {
			return
		}
		fmt.Printf("There are %d record names:\n", len(recordNames))
		for i, recordName := range recordNames {
			j := i + 1
			switch {
			case j < 10:
				fmt.Printf("  %d  %s\n", j, recordName)
			default:
				fmt.Printf("  %d %s\n", j, recordName)
			}
		}
	case verbHelp:
		fmt.Println(Usage())
	default:
		fmt.Printf("\nUnknown command %q.\n", args[0])
	}
	return
}
