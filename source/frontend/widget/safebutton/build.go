package safebutton

import (
	"fmt"
	"path/filepath"

	"github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the safe-button for the framework.
func CreateFramework(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("safebutton.CreateFramework: %w", err)
		}
	}()

	// widget/safebutton/safebutton.go
	oPath := filepath.Join(folderPaths.FrontendWidgetSafeButton, safebuttonFileName)
	if err = utils.ProcessTemplate(safebuttonFileName, oPath, safebuttonTemplate, nil); err != nil {
		return
	}

	return
}
