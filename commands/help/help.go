package help

import (
	"fmt"

	_frontend_ "github.com/JosephABudd/kickfyne/commands/frontend"
	_message_ "github.com/JosephABudd/kickfyne/commands/message"
	_record_ "github.com/JosephABudd/kickfyne/commands/record"
)

const (
	Cmd = "help"
)

// Handler displays the requested help.
func Handler(args []string) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("help.Handler: %w", err)
		}
	}()

	switch args[0] {
	case _frontend_.CmdScreen:
		fmt.Println(_frontend_.UsageScreen())
	case _message_.Cmd:
		fmt.Println(_message_.Usage())
	case _record_.Cmd:
		fmt.Println(_record_.Usage())
	default:
		Usage()
	}
	return
}
