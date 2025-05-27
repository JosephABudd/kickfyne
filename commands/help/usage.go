package help

import (
	_framework_ "github.com/JosephABudd/kickfyne/commands/framework"
	_frontend_ "github.com/JosephABudd/kickfyne/commands/frontend"
	_message_ "github.com/JosephABudd/kickfyne/commands/message"
	_record_ "github.com/JosephABudd/kickfyne/commands/record"
	_version_ "github.com/JosephABudd/kickfyne/commands/version"
)

const (
	newParagraph   = "\n\n"
	gettingStarted = `🍻 INTRODUCING kickfyne!
kickfyne is a tool to help build an application using the fyne toolkit which has among other things a very nice GUI. The fyne toolkit web site is located at https://fyne.io/. This kickfyne project is not in any way associated with this fyne toolkit project.
`
)

func Usage() (usage string) {
	usage =
		_version_.V() + newParagraph +
			gettingStarted + newParagraph +
			_framework_.Usage() + newParagraph +
			_frontend_.Usage() + newParagraph +
			_message_.Usage() + newParagraph +
			_record_.Usage() + newParagraph
	return
}
