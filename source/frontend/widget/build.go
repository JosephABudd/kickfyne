package widget

import (
	"fmt"

	"github.com/JosephABudd/kickfyne/source/frontend/widget/safebutton"
	"github.com/JosephABudd/kickfyne/source/frontend/widget/selection"
	"github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/widget/ files.
func CreateFramework(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("widget.CreateFramework: %w", err)
		}
	}()

	if err = safebutton.CreateFramework(folderPaths); err != nil {
		return
	}
	if err = selection.CreateFramework(folderPaths); err != nil {
		return
	}

	return
}
