package frontend

import (
	"fmt"
)

const (
	CmdScreen        = "screen"
	subCmdHelp       = "help"
	verbRemove       = "remove"
	verbList         = "list"
	verbAddSimple    = "add-simple"
	verbAddDocTabs   = "add-doctabs"
	verbAddAccordion = "add-accordion"
)

// Handler passes control to the correct handlers.
func Handler(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	if !isBuilt || len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Handler: %w", err)
		}
	}()

	switch args[0] {
	case CmdScreen:
		err = handleScreen(pathWD, args, isBuilt, importPrefix)
	case subCmdHelp:
		fmt.Println(Usage())
	default:
		fmt.Println(Usage())
	}
	return
}
